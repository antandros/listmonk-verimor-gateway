package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/antandros/listmonk-verimor-gateway/config"
	"github.com/antandros/listmonk-verimor-gateway/models"
)

func SendSMS(payload models.Listmonk) error {

	smsRequest := models.Verimor{
		Username:   config.Config("VERIMOR_USER"),
		Password:   config.Config("VERIMOR_PASSWD"),
		SourceAddr: config.Config("VERIMOR_SENDER"),
	}

	fmt.Println(payload)

	messages := []models.SMS{}

	for _, val := range payload.Recipients {
		fmt.Println(val)
		messages = append(messages, models.SMS{
			Msg:  payload.Body,
			Dest: val.Attribs.Phone,
		})
	}

	smsRequest.Messages = messages

	requestJson, _ := json.Marshal(smsRequest)
	fmt.Println(requestJson)
	requestBody := bytes.NewBuffer(requestJson)
	resp, err := http.Post("https://sms.verimor.com.tr/v2/send.json", "application/json", requestBody)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}
	fmt.Println(resp.Body)
	return nil
}
