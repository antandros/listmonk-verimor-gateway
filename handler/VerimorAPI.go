// Copyright (C) 2023 Antandros Teknoloji A.Ş.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

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
