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

import ()

type SmsServicer interface {
	GetUserInfo() (*SmsUser, error)
	SendSMS(string, []string) (*SmsResult, error)
	SendSMS_Tpl(int64, []string, []string) (*SmsResult, error)
}

type SmsUser struct {
	Nick             string `json:"nick"`
	Created          string `json:"gmt_created"`
	Mobile           string `json:"mobile"`
	Email            string `json:"email"`
	IpWhites         string `json:"ip_whitelist"`
	ApiVersion       string `json:"api_version"`
	Balance          uint64 `json:"balance"`
	AlarmBalance     uint64 `json:"alarm_balance"`
	EmergencyContact string `json:"emergency_contact"`
	EmergencyMobile  string `json:"emergency_mobile"`
}

type SmsResult struct {
	Count uint   `json:"count"` //成功发送的短信个数
	Fee   uint   `json:"fee"`   //扣费条数，70个字一条，超出70个字时按每67字一条计
	Sid   uint64 `json:"sid"`   ///短信id；群发时以该id+手机号尾号后8位作为短信id
}
