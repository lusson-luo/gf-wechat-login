package v2_saas

import (
	"login-demo/internal/model"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type StationListReq struct {
	g.Meta  `path:"station/list" tags:"充电站管理" method:"get" summary:"充电站点列表"`
	Name    string `json:"name"`
	Address string `json:"address"`
	model.PageReq
}

type StationListRes struct {
	g.Meta     `mime:"application/json" example:"string"`
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Address    string    `json:"address"`
	TenantName string    `json:"tenantName"`
	Coordinate string    `json:"coordinate" `  //
	CreateAt   time.Time `json:"createAt"    ` //
	UpdateAt   time.Time `json:"updateAt"    ` //
}

type StationAddReq struct {
	g.Meta     `path:"station/add" tags:"充电站管理" method:"POST" summary:"添加充电站"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	Coordinate string `json:"coordinate" ` //
}

type StationAddRes struct {
	g.Meta `mime:"application/json" `
}

type StationDelReq struct {
	g.Meta `path:"station/delete" tags:"充电站管理" method:"DELETE" summary:"删除充电站"`
	Id     int `json:"id"`
}

type StationDelRes struct {
	g.Meta `mime:"application/json" `
}

type StationUpdateReq struct {
	g.Meta     `path:"station/update" tags:"充电站管理" method:"post" summary:"修改充电站"`
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	Coordinate string `json:"coordinate" ` //
}

type StationUpdateRes struct {
	g.Meta `mime:"application/json" `
}
