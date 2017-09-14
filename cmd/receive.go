// Copyright © 2017 edwin <edwin.lzh@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo"
	"github.com/lvzhihao/goutils"
	uchat2mq "github.com/lvzhihao/uchat2mq/libs"
	"github.com/lvzhihao/uchatlib"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// receiveCmd represents the receive command
var receiveCmd = &cobra.Command{
	Use:   "receive",
	Short: "uchat2mq http port",
	Long:  `uchat2mq http port`,
	Run: func(cmd *cobra.Command, args []string) {
		var logger *zap.Logger
		if os.Getenv("DEBUG") == "true" {
			logger, _ = zap.NewDevelopment()
		} else {
			logger, _ = zap.NewProduction()
		}
		defer logger.Sync()

		// action config
		receiveActionConfig := viper.GetStringMapString("receive_action_config")
		logger.Debug("receive queue config", zap.Any("data", receiveActionConfig))

		// uchat client
		client := uchatlib.NewClient(viper.GetString("merchant_no"), viper.GetString("merchant_secret"))

		routeKeys := []string{}
		for _, route := range receiveActionConfig {
			routeKeys = append(routeKeys, route)
		}
		tool, err := uchat2mq.NewTool(
			fmt.Sprintf("amqp://%s:%s@%s/%s", viper.GetString("rabbitmq_user"), viper.GetString("rabbitmq_passwd"), viper.GetString("rabbitmq_host"), viper.GetString("rabbitmq_vhost")),
			viper.GetString("rabbitmq_receive_exchange_name"),
			routeKeys,
		)
		if err != nil {
			logger.Error("RabbitMQ Connect Error", zap.Error(err))
		}

		// uchat2mq logger
		uchat2mq.Logger = logger

		// port
		app := goutils.NewEcho()
		app.Any("/*", func(ctx echo.Context) error {
			act := ctx.QueryParam("act")
			if mqRoute, ok := receiveActionConfig[act]; ok {
				str := ctx.FormValue("strContext")
				if strings.Compare(client.Sign(str), ctx.FormValue("strSign")) == 0 {
					tool.Publish(mqRoute, str)
					logger.Debug("Receive Message", zap.String("route", mqRoute), zap.String("message", str))
				} else {
					logger.Error("Error sign", zap.String("strSign", ctx.FormValue("strSign")), zap.String("checkSign", client.Sign(str)))
				}
			} else {
				logger.Error("Unknow Action", zap.String("action", act))
			}
			return ctx.HTML(http.StatusOK, "SUCCESS")
		})
		goutils.EchoStartWithGracefulShutdown(app, viper.GetString("api_host"))
	},
}

func init() {
	RootCmd.AddCommand(receiveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// receiveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// receiveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}