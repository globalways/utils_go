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
	"github.com/globalways/utils_go/algorith"
)

var (
	defaultChars = []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
)

const (
	// default sms auth code attribute
	defaultLen    = 6
	defaultExpire = 10 * 60
)

type SmsManager struct {
	len    int
	expire int64
	store  cache.Cache
}

func NewSmsManager(len int, expire int64, store cache.Cache) *SmsManager {
	return &SmsManager{
		len:    len,
		expire: expire,
		store:  store,
	}
}

func NewDefaultSmsManager() *SmsManager {
	return NewSmsManager(defaultLen, defaultExpire, cache.NewMemoryCache())
}

// generate sms auth code
func (s *SmsManager) getRandChars() []byte {
	return algorith.RandomCreateBytes(s.len, defaultChars...)
}

// generate cache key
func (s *SmsManager) key(tel string) string {
	return tel
}

func (s *SmsManager) GenSmsAuthCode(tel string) (string, error) {
	// get the auth code
	chars := s.getRandChars()

	// save to store
	k := s.key(tel)
	if s.store.IsExist(k) {
		s.store.Delete(k)
	}

	if err := s.store.Put(k, chars, s.expire); err != nil {
		return "", err
	}

	return string(chars), nil
}

func (s *SmsManager) Verify(tel, code string) (success bool) {
	if len(tel) == 0 || len(code) == 0 {
		return
	}

	var chars []byte

	k := s.key(tel)

	if v, ok := s.store.Get(k).([]byte); ok {
		chars = v
	} else {
		return
	}

	defer func() {
		// finally remove it
		s.store.Delete(k)
	}()

	if len(chars) != len(code) {
		return
	}

	// verify challenge
	for i, c := range chars {
		if c != code[i]-48 {
			return
		}
	}

	return true
}
