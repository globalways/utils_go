// Copyright 2015 mit.zhao.chiu@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.
package sms

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/globalways/utils_go/convert"
	"github.com/globalways/utils_go/http/client"
	"log"
	"net/url"
	"reflect"
	"strings"
)

const (
	_BaseUrl_YP string = "http://yunpian.com" // 服务http地址
	_Version_YP string = "v1"                 // 服务版本号
)

const (
	Tpl_YP_1 int64 = iota + 1
	Tpl_YP_2
	Tpl_YP_3
	Tpl_YP_4
	Tpl_YP_5
	Tpl_YP_6
	Tpl_YP_7
	Tpl_YP_8
	Tpl_YP_9
)

var (
	_Uri_Get_User_Info_YP string = fmt.Sprintf("%s/%s/user/get.json", _BaseUrl_YP, _Version_YP)     // 查询账户信息的http地址
	_Uri_Send_SMS_YP      string = fmt.Sprintf("%s/%s/sms/send.json", _BaseUrl_YP, _Version_YP)     // 通用发送接口的http地址
	_Uri_Tpl_Send_SMS_YP  string = fmt.Sprintf("%s/%s/sms/tpl_send.json", _BaseUrl_YP, _Version_YP) // 模版发送接口的http地址

	_Tpl_Messages map[int64]string = map[int64]string{
		Tpl_YP_1: "#company#=%s&#code#=%s",
		Tpl_YP_2: "#company#=%s&#code#=%s",
		Tpl_YP_3: "#company#=%s&#name#=%s&#code#=%s",
		Tpl_YP_4: "#company#=%s&#name#=%s&#code#=%s&#hour#=%s",
		Tpl_YP_5: "#company#=%s&#app#=%s&#code#=%s",
		Tpl_YP_6: "#company#=%s&#app#=%s&#code#=%s",
		Tpl_YP_7: "#company#=%s&#code#=%s",
		Tpl_YP_8: "#company#=%s&#code#=%s&#tel#=%s",
		Tpl_YP_9: "#company#=%s&#code#=%s&#app#=%s",
	}
)

type YunPianService struct {
	ApiKey string
}

/**
* 新建云片服务
 */
func NewYunPian(apikey string) *YunPianService {
	return &YunPianService{
		ApiKey: apikey,
	}
}

/**
* 取账户信息
 */
func (y *YunPianService) GetUserInfo() (*SmsUser, error) {
	params := url.Values{
		"apikey": []string{y.ApiKey},
	}

	url := fmt.Sprintf("%s?%s", _Uri_Get_User_Info_YP, params.Encode())
	rsp, err := client.ForwardHttp("GET", url, nil)
	if err != nil {
		log.Printf("err: %v\n", err)
		return nil, errors.New("[YunPian] " + err.Error())
	}

	data := client.GetForwardHttpBody(rsp.Body)
	body := make(map[string]interface{})
	if err := json.Unmarshal(data, &body); err != nil {
		log.Printf("err: %v\n", err)
		return nil, errors.New("[YunPian] " + err.Error())
	}

	if code, ok := body["code"]; !ok {
		return nil, errors.New("[YunPian] invalid request.")
	} else if code.(float64) != 0 {
		return nil, errors.New("[YunPian] " + body["msg"].(string))
	}

	user, ok := body["user"].(map[string]interface{})
	if !ok {
		log.Printf("body[user] interface: %v", reflect.TypeOf(body["user"]))
		return nil, errors.New("[YunPian] can not parse body[user] to map[string]interface{}.")
	}

	return &SmsUser{
		Nick:             user["nick"].(string),
		Created:          user["gmt_created"].(string),
		Mobile:           user["mobile"].(string),
		Email:            user["email"].(string),
		IpWhites:         user["ip_whitelist"].(string),
		ApiVersion:       user["api_version"].(string),
		Balance:          uint64(user["balance"].(float64)),
		AlarmBalance:     uint64(user["alarm_balance"].(float64)),
		EmergencyContact: user["emergency_contact"].(string),
		EmergencyMobile:  user["emergency_mobile"].(string),
	}, nil
}

/**
* 发短信
 */
func (y *YunPianService) SendSMS(text string, mobiles []string) (*SmsResult, error) {
	params := &url.Values{
		"apikey": []string{y.ApiKey},
		"text":   []string{text},
		"mobile": []string{strings.Join(mobiles, ",")},
	}

	rsp, err := client.ForwardHttp("POST", _Uri_Send_SMS_YP, bytes.NewBufferString(params.Encode()))
	if err != nil {
		log.Printf("err: %v\n", err)
		return nil, errors.New("[YunPian] " + err.Error())
	}

	data := client.GetForwardHttpBody(rsp.Body)
	body := make(map[string]interface{})
	if err := json.Unmarshal(data, &body); err != nil {
		log.Printf("err: %v\n", err)
		return nil, errors.New("[YunPian] " + err.Error())
	}

	if code, ok := body["code"]; !ok {
		return nil, errors.New("[YunPian] invalid request.")
	} else if code.(float64) != 0 {
		return nil, errors.New("[YunPian] " + body["msg"].(string))
	}

	result, ok := body["result"].(map[string]interface{})
	if !ok {
		log.Printf("body[result] interface: %v", reflect.TypeOf(body["user"]))
		return nil, errors.New("[YunPian] can not parse body[result] to map[string]interface{}.")
	}

	return &SmsResult{
		Count: uint(result["count"].(float64)),
		Fee:   uint(result["fee"].(float64)),
		Sid:   uint64(result["sid"].(float64)),
	}, nil
}

/**
* 通过模版发送短信
 */
func (y *YunPianService) SendSMS_Tpl(tplid int64, mobiles []string, args []string) (*SmsResult, error) {

	tpl, ok := _Tpl_Messages[tplid]
	if !ok {
		return nil, errors.New("[YunPian] invalid sms template.")
	}

	params := &url.Values{
		"apikey": []string{y.ApiKey},
		"mobile": []string{strings.Join(mobiles, ",")},
		"tpl_id": []string{convert.Int642str(tplid)},
		"tpl_value": []string{fmt.Sprintf(tpl, func() []interface{} {
			faces := make([]interface{}, 0)
			for _, arg := range args {
				faces = append(faces, arg)
			}

			return faces
		}()...)},
	}

	rsp, err := client.ForwardHttp("POST", _Uri_Tpl_Send_SMS_YP, bytes.NewBufferString(params.Encode()))
	if err != nil {
		log.Printf("err: %v\n", err)
		return nil, errors.New("[YunPian] " + err.Error())
	}

	data := client.GetForwardHttpBody(rsp.Body)
	body := make(map[string]interface{})
	if err := json.Unmarshal(data, &body); err != nil {
		log.Printf("err: %v\n", err)
		return nil, errors.New("[YunPian] " + err.Error())
	}

	if code, ok := body["code"]; !ok {
		return nil, errors.New("[YunPian] invalid request.")
	} else if code.(float64) != 0 {
		return nil, errors.New("[YunPian] " + body["msg"].(string))
	}

	result, ok := body["result"].(map[string]interface{})
	if !ok {
		log.Printf("body[result] interface: %v", reflect.TypeOf(body["user"]))
		return nil, errors.New("[YunPian] can not parse body[result] to map[string]interface{}.")
	}

	return &SmsResult{
		Count: uint(result["count"].(float64)),
		Fee:   uint(result["fee"].(float64)),
		Sid:   uint64(result["sid"].(float64)),
	}, nil
}
