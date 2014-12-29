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
package client

import (
	"encoding/json"
	"github.com/globalways/utils_go/errors"
	"github.com/mreiferson/httpclient"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	transport = &httpclient.Transport{
		ConnectTimeout:        1 * time.Second,
		RequestTimeout:        10 * time.Second,
		ResponseHeaderTimeout: 5 * time.Second,
	}
	client = &http.Client{Transport: transport}
)

// 转发http请求
func forwardHttp(method, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	return client.Do(req)
}

// 获取http response body
func getForwardHttpBody(body io.ReadCloser) []byte {
	bodyBytes, _ := ioutil.ReadAll(body)

	return bodyBytes
}

// 请求API服务器
func ForwardAPI(method, url string, body io.Reader) (*errors.ClientRsp, error) {
	rsp, err := forwardHttp(method, url, body)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	clientRsp := new(errors.ClientRsp)
	if err := json.Unmarshal(getForwardHttpBody(rsp.Body), clientRsp); err != nil {
		return nil, err
	}

	return clientRsp, nil
}
