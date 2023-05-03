package controller_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/gogf/gf/v2/net/gclient"
	"github.com/stretchr/testify/assert"

	_ "login-demo/internal/packed"
)

func TestLogin(t *testing.T) {
	client := gclient.New()
	params := map[string]string{
		"username": "test_user",
		"password": "test_password",
	}
	// 发送 HTTP POST 请求
	resp, err := client.Post(context.Background(), "http://127.0.0.1:8000/api/user/login", params)
	assert.Empty(t, err)
	defer resp.Close()
	assert.Equal(t, resp.StatusCode, 200)
	var data map[string]interface{}
	err = json.Unmarshal([]byte(resp.ReadAllString()), &data)
	// data := resp.ReadAllString()
	// .GetMap("data")
	assert.NotEmpty(t, data["token"].(string))
	// assert.Equal(t, data["role"].(string)), "test_role")
	// assert.Equal(t, data["role"].(string)), "test_role")
}

// func TestRefresh(t *testing.T) {
// 	r := ghttp.NewTestRequest()
// 	r.SetMethod("POST")
// 	r.SetUrl("/refresh")
// 	r.SetHeader("Authorization", "Bearer "+testToken)
// 	resp := r.GetResponse()
// 	defer resp.Close()
// 	assert.Equal(t, resp.StatusCode, 200)
// 	data := resp.GetJson().GetMap("data")
// 	assert.NotEmpty(t, data.GetString("token"))
// }
