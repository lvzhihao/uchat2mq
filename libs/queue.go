package uchat2mq

import (
	"bytes"
	"fmt"
	"net/http"
)

func CreateExchange(api, user, passwd, vhost, exchange string) error {
	client := &http.Client{}
	b := bytes.NewBufferString(`{"type":"topic","auto_delete":false,"durable":true,"internal":false,"arguments":[]}`)
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/exchanges/%s/%s", api, vhost, exchange), b)
	if err != nil {
		return err
	}
	// enusre exchange
	req.SetBasicAuth(user, passwd)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if (resp.StatusCode == http.StatusNoContent) || (resp.StatusCode == http.StatusCreated) {
		return nil
	} else {
		return fmt.Errorf("CreateExchange StatusError: %d, %v", resp.StatusCode, resp)
	}
}

func CreateQueue(api, user, passwd, vhost, name string) error {
	client := &http.Client{}
	b := bytes.NewBufferString(`{"auto_delete":false, "durable":true, "arguments":[]}`)
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/queues/%s/%s", api, vhost, name), b)
	if err != nil {
		return err
	}
	// enusre queue
	req.SetBasicAuth(user, passwd)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if (resp.StatusCode == http.StatusNoContent) || (resp.StatusCode == http.StatusCreated) {
		return nil
	} else {
		return fmt.Errorf("CreateQueue StatusError: %d, %v", resp.StatusCode, resp)
	}
}

func BindRoutingKey(api, user, passwd, vhost, name, exchange, key string) error {
	client := &http.Client{}
	b := bytes.NewBufferString(`{"routing_key":"` + key + `", "arguments":[]}`)
	// ensure binding
	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/bindings/%s/e/%s/q/%s", api, vhost, exchange, name),
		b)
	req.SetBasicAuth(user, passwd)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if (resp.StatusCode == http.StatusNoContent) || (resp.StatusCode == http.StatusCreated) {
		return nil
	} else {
		return fmt.Errorf("BindRoutingKey StatusError: %d, %v", resp.StatusCode, resp)
	}
}

func RegisterQueue(api, user, passwd, vhost, name, exchange string, keys []string) error {
	err := CreateQueue(api, user, passwd, vhost, name)
	if err != nil {
		return err
	}
	for _, key := range keys {
		if exchange != "" && key != "" {
			err := BindRoutingKey(api, user, passwd, vhost, name, exchange, key)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
