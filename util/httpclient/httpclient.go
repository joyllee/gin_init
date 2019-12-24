package httpclient

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/LucasGao67/go-querystring/query"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type HttpClient struct{}

func (self *HttpClient) NewClient() *http.Client {
	t := &http.Transport{}
	t.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	t.IdleConnTimeout = 30000 * time.Millisecond
	return &http.Client{Transport: t, Timeout: 5000 * time.Millisecond}
}

func (self *HttpClient) DoRequest(traceId string, request *http.Request) (*http.Response, error) {
	//logrus.WithFields(logrus.Fields{"traceId": traceId, "httpclient request": request}).Trace()
	client := self.NewClient()
	resp, err := client.Do(request)
	if err != nil {
		//logrus.WithFields(logrus.Fields{"traceId": traceId, "httpclient request": request, "err": err}).Warn()
		logrus.WithFields(logrus.Fields{"traceId": traceId, "err": err}).Warn()
	}
	return resp, err
}

func (self *HttpClient) UnmarshalResponse(resp *http.Response, result proto.Message) error {
	m := jsonpb.Unmarshaler{AllowUnknownFields: true}
	return m.Unmarshal(resp.Body, result)
}

func (self *HttpClient) NewPostRequest(host string, path string, msg proto.Message) (*http.Request, error) {
	url := fmt.Sprintf("%s%s", host, path)
	m := jsonpb.Marshaler{OrigName: true}
	jsonstr, err := m.MarshalToString(msg)
	if err != nil {
		logrus.Error("fail to marshaltostring", msg)
		return nil, err
	}
	payload := bytes.NewBufferString(jsonstr)
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func (self *HttpClient) NewGetRequest(host string, path string, msg proto.Message) (*http.Request, error) {
	v, err := query.ValuesWithTag(msg, "json")
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s%s?%s", host, path, v.Encode())

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (self *HttpClient) NewPostRequestWithJson(host string, path string, jsonstr string) (*http.Request, error) {
	url := fmt.Sprintf("%s%s", host, path)
	payload := bytes.NewBufferString(jsonstr)
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func (self *HttpClient) NewGetRequestWithJson(host string, path string, jsonstr string) (*http.Request, error) {
	v, err := query.ValuesWithTag(jsonstr, "json")
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s%s?%s", host, path, v.Encode())

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
