// -*- coding: utf-8 -*-

package requests

import (
	"fmt"
	"go.uber.org/zap"
	"io"
	"net/http"
	"strings"
	"time"
)

// Requests 封装HTTP请求
func Requests(method, url string, data string, token string) ([]byte, error) {
	// 创建HTTP实例
	client := &http.Client{Timeout: 20 * time.Second}

	// 添加请求数据
	var ReqData = strings.NewReader(data)
	req, err := http.NewRequest(method, url, ReqData)
	// 添加请求Token
	if token != "" {
		Token := fmt.Sprintf("Bearer %s", token)
		req.Header.Set("Authorization", Token)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "LDDP client")
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		zap.L().Error(err.Error())
		return []byte(""), err
	}
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		zap.L().Error(err.Error())
		return []byte(""), err
	}

	zap.L().Debug(fmt.Sprintf("%s\n", bodyText))
	return bodyText, err

}

// Version 封装HTTP请求
func Version(method, url, data, token string) ([]byte, error) {
	// 创建HTTP实例
	client := &http.Client{Timeout: 20 * time.Second}

	// 添加请求数据
	var ReqData = strings.NewReader(data)
	req, err := http.NewRequest(method, url, ReqData)
	// 添加请求Token
	if token != "" {
		Token := fmt.Sprintf("Bearer %s", token)
		req.Header.Set("Authorization", Token)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "LDDP client")
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		zap.L().Error(err.Error())
		return []byte(""), err
	}
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		zap.L().Error(err.Error())
		return []byte(""), err
	}

	return bodyText, err

}
