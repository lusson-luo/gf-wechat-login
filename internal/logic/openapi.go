package logic

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type OpenApiLogic struct{}

var (
	Token      string
	OpenApiUrl string
	OpenApi    OpenApiLogic = OpenApiLogic{}
)

const (
	TokenConfigKey   = "coding-openapi.token"
	OpenApiConfigKey = "coding-openapi.url"
)

func getConfig(ctx context.Context, key string) (string, error) {
	val, err := g.Cfg().Get(context.Background(), key)
	if err != nil {
		return "", gerror.New(fmt.Sprintf("无法启动，未找到配置: %s", TokenConfigKey))
	}
	configVal := val.String()
	if strings.TrimSpace(configVal) == "" {
		return "", gerror.New(fmt.Sprintf("无法启动，配置的值为空: %s", TokenConfigKey))
	}
	return configVal, nil
}

func init() {
	token, err := getConfig(context.Background(), TokenConfigKey)
	if err != nil {
		panic(err)
	}
	Token = token
	openApiUrl, err := getConfig(context.Background(), OpenApiConfigKey)
	if err != nil {
		panic(err)
	}
	OpenApiUrl = openApiUrl
}

func (OpenApiLogic) SendOpenAPI() (string, error) {
	client := g.Client()
	client.SetHeaderMap(map[string]string{
		"Authorization": "token " + Token,
		"Content-Type":  "application/json",
	})
	resp, err := client.Post(context.Background(), fmt.Sprintf("%s/open-api", OpenApiUrl), map[string]interface{}{
		"Action":      "DescribeCodingProjects",
		"ProjectName": "",
		"PageNumber":  1,
		"PageSize":    10,
	})
	if err != nil {
		return "", err
	}
	responseBody := resp.ReadAllString()
	return responseBody, err
}
