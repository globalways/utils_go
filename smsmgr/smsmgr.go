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
package smsmgr

import (
	"github.com/astaxie/beego/cache"
	"math"
	"github.com/globalways/utils_go/random"
	"github.com/globalways/utils_go/convert"
)

const (
	// default sms auth code attribute
	defaultLen    = 6
	defaultExpire = 10 * 60
)

type SmsManager struct {
	min    int
	max    int
	expire int64
	store  cache.Cache
}

func NewSmsManager(len int, expire int64, store cache.Cache) *SmsManager {
	return &SmsManager{
		min: convert.Float642Int(math.Pow(10, float64(len - 1))),
		max: convert.Float642Int(math.Pow(10, float64(len))),
		expire: expire,
		store:  store,
	}
}

func NewDefaultSmsManager() *SmsManager {
	return NewSmsManager(defaultLen, defaultExpire, cache.NewMemoryCache())
}

// generate sms auth code
func (s *SmsManager) getRandCode() string {
	return random.RandIntStr(s.min, s.max)
}

// generate cache key
func (s *SmsManager) key(tel string) string {
	return tel
}

func (s *SmsManager) GenSmsAuthCode(tel string) (string, error) {
	// get the auth code
	chars := s.getRandCode()

	// save to store
	k := s.key(tel)
	if s.store.IsExist(k) {
		s.store.Delete(k)
	}

	if err := s.store.Put(k, chars, s.expire); err != nil {
		return "", err
	}

	return chars, nil
}

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

	if len(chars) != len(code) || chars != code{
		return
	}

	return true
}
