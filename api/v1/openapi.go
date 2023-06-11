package v1

import "github.com/gogf/gf/v2/frame/g"

type ProjectListReq struct {
	g.Meta `path:"/api/project/list" tags:"projectList" method:"get" summary:"获得项目列表"`
}

type ProjectListRes struct {
	g.Meta `mime:"application/json" `
	Data   string `json:"data"`
}
