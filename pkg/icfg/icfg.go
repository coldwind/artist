package icfg

import (
	"errors"
	"io"
	"net/http"
	"os"
	"sync"

	"github.com/coldwind/artist/pkg/ilog"
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

type CfgType int

type service struct {
	sync.RWMutex
	cache map[string]interface{}
}

var (
	CfgTypeYaml CfgType = 1
	CfgTypeJson CfgType = 2
)

var handle *service

func init() {
	handle = &service{
		cache: map[string]interface{}{},
	}
}

func Load(t CfgType, key, path string, data interface{}) error {
	cfg, err := os.ReadFile(path)
	if err != nil {
		ilog.Error("get cfg file error", zap.String("path", path), zap.Error(err))
		return err
	}

	err = errors.New("register failure")
	if t == CfgTypeYaml {
		err = yaml.Unmarshal(cfg, data)
		if err != nil {
			ilog.Error("yaml Unmarshal file error", zap.String("path", path), zap.Error(err))
		}
	} else if t == CfgTypeJson {
		err = jsoniter.Unmarshal(cfg, data)
		if err != nil {
			ilog.Error("json Unmarshal file error", zap.String("path", path), zap.Error(err))
		}
	}

	if err != nil {
		return err
	}

	handle.Lock()
	defer handle.Unlock()

	handle.cache[key] = data

	return nil
}

func LoadRemote(t CfgType, key, url string, data interface{}) error {
	// 发送HTTP GET请求
	resp, err := http.Get(url)
	if err != nil {
		// 处理错误
		panic(err)
	}
	defer resp.Body.Close() // 确保关闭响应体

	// 读取响应体的内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// 处理错误
		return err

	}

	if t == CfgTypeYaml {
		err = yaml.Unmarshal(body, data)
		if err != nil {
			ilog.Error("yaml Unmarshal file error", zap.String("url", url), zap.Error(err))
		}
	} else if t == CfgTypeJson {
		err = jsoniter.Unmarshal(body, data)
		if err != nil {
			ilog.Error("json Unmarshal file error", zap.String("url", url), zap.Error(err))
		}
	}

	if err != nil {
		return err
	}

	handle.Lock()
	defer handle.Unlock()

	handle.cache[key] = data

	return nil
}

func Get(key string) interface{} {
	handle.RLock()
	defer handle.RUnlock()
	if _, ok := handle.cache[key]; ok {
		return handle.cache[key]
	}

	return nil
}
