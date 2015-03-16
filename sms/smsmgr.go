// Copyright 2014 mit.zhao.chiu@gmail.com
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
	"github.com/astaxie/beego/cache"
	"github.com/globalways/utils_go/random"
	"log"
	"math"
	"time"
)

const (
	// default sms auth code attribute
	defaultLen    = 6
	defaultExpire = 1 * 60 * 60
	defaultChan   = 100
)

type SmsManager struct {
	min        int
	max        int
	expire     int64
	store      cache.Cache
	service    SmsServicer
	smsChan    chan []string              // 0: text 1 - n-1: mobiles
	smsChanTpl chan map[int64][2][]string // key: template id, value[0]:mobiles, value[1]:args
}

// 内容直接发送协程
func (s *SmsManager) sendSMS(messages chan []string) {
	ticker := time.NewTicker(time.Hour)
	for {
		select {
		case message := <-messages:
			result, err := s.service.SendSMS(message[0], message[1:])
			if err != nil {
				log.Printf("---SMS Manager---: err: %v\n", err)
			} else {
				log.Printf("---SMS Manager---: result: %+v\n", result)
			}
		case <-ticker.C:
			log.Printf("---SMS Manager---: current sms chan len: %v\n", len(s.smsChan))
		}
	}
}

// 模版发送协程
func (s *SmsManager) sendSMS_tpl(messages chan map[int64][2][]string) {
	ticker := time.NewTicker(time.Hour)
	for {
		select {
		case message := <-messages:
			for k, v := range message {
				result, err := s.service.SendSMS_Tpl(k, v[0], v[1])
				if err != nil {
					log.Printf("---SMS Manager---: err: %v\n", err)
				} else {
					log.Printf("---SMS Manager---: result: %+v\n", result)
				}
			}

		case <-ticker.C:
			log.Printf("---SMS Manager---: current sms template chan len: %v\n", len(s.smsChan))
		}
	}
}

// 新建短信manager
func NewSmsManager(len int, expire int64, store cache.Cache, service SmsServicer, cLen int) *SmsManager {
	manager := &SmsManager{
		min:        int(math.Pow(10, float64(len-1))),
		max:        int(math.Pow(10, float64(len))),
		expire:     expire,
		store:      store,
		service:    service,
		smsChan:    make(chan []string, cLen),
		smsChanTpl: make(chan map[int64][2][]string, cLen),
	}

	// listen up for send sms by text
	go manager.sendSMS(manager.smsChan)
	// listen up for send sms by template
	go manager.sendSMS_tpl(manager.smsChanTpl)

	return manager
}

// 默认短信manager
func NewDefaultSmsManager(apikey string) *SmsManager {
	return NewSmsManager(defaultLen, defaultExpire, cache.NewMemoryCache(), NewYunPian(apikey), defaultChan)
}

// generate sms auth code
func (s *SmsManager) getRandCode() string {
	return random.RandIntStr(s.min, s.max)
}

// generate cache key
func (s *SmsManager) key(tel string) string {
	return tel
}

// sms code
func (s *SmsManager) Code(tel string) (string, error) {
	// if exist, directly return
	k := s.key(tel)
	if s.store.IsExist(k) {
		if v, ok := s.store.Get(k).(string); ok {
			return v, nil
		}
	}

	// get the auth code
	chars := s.getRandCode()

	// save to store
	if err := s.store.Put(k, chars, s.expire); err != nil {
		return "", err
	}

	return chars, nil
}

// varify sms code
func (s *SmsManager) Verify(tel, code string) (success bool) {
	if len(tel) == 0 || len(code) == 0 {
		return
	}

	var chars string

	k := s.key(tel)

	if v, ok := s.store.Get(k).(string); ok {
		chars = v
	} else {
		return
	}

	defer func() {
		// finally remove it
		s.store.Delete(k)
	}()

	if len(chars) != len(code) || chars != code {
		return
	}

	return true
}

// 文字短信
func (s *SmsManager) SendSMS(text string, mobiles []string) {
	args := make([]string, 0)
	args = append(args, text)
	args = append(args, mobiles...)

	s.smsChan <- args
}

// 模版短信
func (s *SmsManager) SendSMS_tpl(tplid int64, mobiles []string, arg ...string) {
	args := make(map[int64][2][]string)
	args[tplid] = [2][]string{
		mobiles,
		arg,
	}

	s.smsChanTpl <- args
}
