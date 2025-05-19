package ihttp

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/valyala/fasthttp"
)

var (
	errStatus = errors.New("status error")
)

func Get(url string, data map[string]string, header map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// 添加查询字符串参数
	query := req.URL.Query()
	for k, v := range data {
		query.Add(k, v)
	}

	for k, v := range header {
		req.Header.Add(k, v)
	}

	req.URL.RawQuery = query.Encode()

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Get request failed!")
		return nil, err
	} else {
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Read response failed!")
			return nil, err
		} else {
			return body, nil
		}
	}
}

func Post(url string, data map[string]string, header map[string]string) ([]byte, error) {
	return Request(url, "POST", data, header)
}

func PostJson(url string, data []byte, extraHeader map[string]string) ([]byte, error) {
	args := &fasthttp.Args{}
	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()

	defer func() {
		fasthttp.ReleaseArgs(args)
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()

	req.SetRequestURI(url)
	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/json")
	req.Header.Add("Accept", "application/json")
	for k, v := range extraHeader {
		req.Header.Add(k, v)
	}

	req.SetBody(data)

	client := &fasthttp.Client{}
	if err := client.Do(req, res); err != nil {
		return nil, err
	}

	if res.StatusCode() != fasthttp.StatusOK {
		return nil, errors.New("status error")
	}

	return res.Body(), nil
}

func Request(url, method string, param map[string]string, header map[string]string) ([]byte, error) {
	args := &fasthttp.Args{}
	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()

	defer func() {
		fasthttp.ReleaseArgs(args)
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()

	for k, v := range param {
		args.Add(k, v)
	}

	req.SetRequestURI(url)
	req.Header.SetMethod(method)
	if len(header) == 0 {
		req.Header.SetContentType("application/x-www-form-urlencoded")
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}
	req.SetBody(args.QueryString())

	client := &fasthttp.Client{}
	if err := client.Do(req, res); err != nil {
		return nil, err
	}

	if res.StatusCode() != fasthttp.StatusOK {
		return nil, errStatus
	}

	return res.Body(), nil
}
