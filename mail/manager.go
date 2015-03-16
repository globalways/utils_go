// Copyright 2015 mint.zhao.chiu@gmail.com
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
package mail

import (
	"github.com/go-gomail/gomail"
	"log"
)

type MailManager struct {
	host        string
	port        int
	username    string
	password    string
	messageChan chan *Message
	mailer      *gomail.Mailer
}

func (mgr *MailManager) send(messages <-chan *Message) {
	for {
		select {
		case msg := <-messages:
			if err := mgr.mailer.Send(msg.mail()); err != nil {
				log.Printf("mail send err:%v\n", err)
			}
		}
	}
}

func (mgr *MailManager) Send(message *Message) {
	mgr.messageChan <- message
}

func NewMailManager(host, userName, password string, port, chanLen int) *MailManager {
	manager := &MailManager{
		host:        host,
		port:        port,
		username:    userName,
		password:    password,
		messageChan: make(chan *Message, chanLen),
		mailer:      gomail.NewMailer(host, userName, password, port),
	}

	// message channel start
	go manager.send(manager.messageChan)

	return manager
}
