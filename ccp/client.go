/*Copyright (c) 2019 Cisco and/or its affiliates.

This software is licensed to you under the terms of the Cisco Sample
Code License, Version 1.0 (the "License"). You may obtain a copy of the
License at

               https://developer.cisco.com/docs/licenses

All use of the material herein must be in accordance with the terms of
the License. All rights not expressly granted by the License are
reserved. Unless required by applicable law or agreed to separately in
writing, software distributed under the License is distributed on an "AS
IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
or implied.*/

package ccp

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"reflect"
)

//import "encoding/json"

type Client struct {
	Username string
	Password string
	BaseURL  string
}

var jar, err = cookiejar.New(nil)

func NewClient(username, password, baseURL string) *Client {

	return &Client{
		Username: username,
		Password: password,
		BaseURL:  baseURL,
	}
}

func (s *Client) doRequest(req *http.Request) ([]byte, error) {

	var client *http.Client

	req.Header.Add("Content-Type", "application/json")
	//req.SetBasicAuth(s.Username, s.Password)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	if len(jar.Cookies(req.URL)) > 0  && jar.Cookies(req.URL)[0].Name == "CXAccessToken" {
		log.Print("adding token to request")
		req.Header.Add("x-auth-token", jar.Cookies(req.URL)[0].Value)
	}
	client = &http.Client{Transport: tr, Jar: jar}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if 200 != resp.StatusCode && 201 != resp.StatusCode && 202 != resp.StatusCode && 204 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}

	if err != nil {
		return nil, err
	}

	return body, nil
}

// Helper routine used to return pointer - will used to simplify the use of the clientlibrary
func Bool(value bool) *bool {
	return &value
}

// Helper routine used to return pointer - will used to simplify the use of the clientlibrary
func Int(value int) *int {
	return &value
}

// Helper routine used to return pointer - will used to simplify the use of the clientlibrary
func Int64(value int64) *int64 {
	return &value
}

// Helper routine used to return pointer - will used to simplify the use of the clientlibrary
func String(value string) *string {
	return &value
}

// Helper routine used to return pointer - will used to simplify the use of the clientlibrary
func Float32(value float32) *float32 {
	return &value
}

// Helper routine used to return pointer - will used to simplify the use of the clientlibrary
func Float64(value float64) *float64 {
	return &value
}

//modified from unexported nonzero function in the validtor package
//https://github.com/go-validator/validator/blob/v2/builtins.go
func nonzero(v interface{}) bool {
	st := reflect.ValueOf(v)
	nonZeroValue := false
	switch st.Kind() {
	case reflect.Ptr, reflect.Interface:
		nonZeroValue = st.IsNil()
	case reflect.Invalid:
		nonZeroValue = true // always invalid
	case reflect.Struct:
		nonZeroValue = false // always valid since only nil pointers are empty
	default:
		return true
	}

	if nonZeroValue {
		return true
	}
	return false
}
