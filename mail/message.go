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
	"github.com/globalways/utils_go/container"
	"github.com/go-gomail/gomail"
)

type Message struct {
	sender      [2]string
	receivers   [][2]string
	ccs         [][2]string
	bccs        [][2]string
	subject     string
	contentType string
	body        string
}

func NewMessage() *Message {
	return &Message{
		sender:    [2]string{},
		receivers: make([][2]string, 0),
		ccs:       make([][2]string, 0),
		bccs:      make([][2]string, 0),
	}
}

func (msg *Message) SetSender(sender, senderName string) {
	msg.sender[0] = sender
	msg.sender[1] = senderName
}

func (msg *Message) SetReceivers(receivers ...[2]string) {
	for _, receiver := range receivers {
		msg.SetReceiver(receiver[0], receiver[1])
	}
}

func (msg *Message) SetReceiver(receiverAddr, receiverName string) {
	receiver := [2]string{receiverAddr, receiverName}
	if container.Contains(receiver, msg.receivers) {
		return
	}

	msg.receivers = append(msg.receivers, receiver)
}

func (msg *Message) SetCcs(ccs ...[2]string) {
	for _, cc := range ccs {
		msg.SetCc(cc[0], cc[1])
	}
}

func (msg *Message) SetCc(ccAddr, ccName string) {
	cc := [2]string{ccAddr, ccName}
	if container.Contains(cc, msg.ccs) {
		return
	}

	msg.ccs = append(msg.ccs, cc)
}

func (msg *Message) SetBccs(bccs ...[2]string) {
	for _, bcc := range bccs {
		msg.SetBcc(bcc[0], bcc[1])
	}
}

func (msg *Message) SetBcc(bccAddr, bccName string) {
	bcc := [2]string{bccAddr, bccName}
	if container.Contains(bcc, msg.bccs) {
		return
	}

	msg.bccs = append(msg.bccs, bcc)
}

func (msg *Message) SetSubject(subject string) {
	msg.subject = subject
}

func (msg *Message) SetPlainBody(body string) {
	msg.contentType = "text/plain"
	msg.body = body
}

func (msg *Message) SetHttpBody(body string) {
	msg.contentType = "text/html"
	msg.body = body
}

func (msg *Message) mail() *gomail.Message {
	message := gomail.NewMessage()

	// From
	senderAddr := msg.sender[0]
	senderName := msg.sender[1]
	if senderName != "" {
		message.SetAddressHeader("From", senderAddr, senderName)
	} else {
		message.SetHeader("From", senderAddr)
	}

	// To
	receivers := make([]string, 0)
	for _, receiver := range msg.receivers {
		receiverAddr := receiver[0]
		receiverName := receiver[1]
		if receiverName != "" {
			receivers = append(receivers, message.FormatAddress(receiverAddr, receiverName))
		} else {
			receivers = append(receivers, receiverAddr)
		}
	}
	message.SetHeader("To", receivers...)

	// CC
	ccs := make([]string, 0)
	for _, cc := range msg.ccs {
		ccAddr := cc[0]
		ccName := cc[1]
		if ccName != "" {
			ccs = append(ccs, message.FormatAddress(ccAddr, ccName))
		} else {
			ccs = append(ccs, ccAddr)
		}
	}
	message.SetHeader("Cc", ccs...)

	// Bcc
	bccs := make([]string, 0)
	for _, bcc := range msg.bccs {
		bccAddr := bcc[0]
		bccName := bcc[1]

		if bccName != "" {
			bccs = append(bccs, message.FormatAddress(bccAddr, bccName))
		} else {
			bccs = append(bccs, bccAddr)
		}
	}
	message.SetHeader("Bcc", bccs...)

	// Subject
	message.SetHeader("Subject", msg.subject)

	// Body
	message.SetBody(msg.contentType, msg.body)

	return message
}
