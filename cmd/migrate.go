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
	"log"
	"os"
	"strings"

	uchat2mq "github.com/lvzhihao/uchat2mq/libs"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate rabbitmq queue and binding",
	Long:  `migrate rabbitmq queue and binding`,
	Run: func(cmd *cobra.Command, args []string) {
		var logger *zap.Logger
		if os.Getenv("DEBUG") == "true" {
			logger, _ = zap.NewDevelopment()
		} else {
			logger, _ = zap.NewProduction()
		}
		defer logger.Sync()

		// queue config
		receiveQueueConfig := viper.GetStringMapStringSlice("receive_queue_config")
		logger.Debug("receive queue config", zap.Any("data", receiveQueueConfig))

		// receive queue migrate
		for k, v := range receiveQueueConfig {
			migrateQueue(k, viper.GetString("rabbitmq_receive_exchange_name"), v)
		}
	},
}

func migrateQueue(name, exchange string, key []string) {
	rmqApi := viper.GetString("rabbitmq_api")
	rmqUser := viper.GetString("rabbitmq_user")
	rmqPasswd := viper.GetString("rabbitmq_passwd")
	rmqVhost := viper.GetString("rabbitmq_vhost")
	err := uchat2mq.RegisterQueue(
		rmqApi, rmqUser, rmqPasswd, rmqVhost,
		name, exchange, key,
	)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("queue create success: %s bind %s %s\n", name, exchange, strings.Join(key, ","))
	}
}

func init() {
	RootCmd.AddCommand(migrateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
