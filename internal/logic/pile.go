package logic

import (
	"context"
	"fmt"
	"login-demo/internal/dao"
	"login-demo/internal/model"
	"login-demo/internal/model/do"
	"login-demo/internal/model/entity"

	"github.com/gogf/gf/v2/os/gtime"
)

var Pile PileLogic

type PileLogic struct {
}

func (*PileLogic) PileList(ctx context.Context, query entity.Pile, page model.PageReq) (piles []model.Pile, count int, err error) {
	model := dao.Pile.Ctx(ctx).LeftJoin("station", "station.id = pile.station_id")
	if query.StationId != 0 {
		model = model.Where("pile.station_id = ?", query.StationId)
	}
	if query.Code != "" {
		model = model.Where("pile.code like ?", fmt.Sprintf("%%%s%%", query.Code))
	}
	count, err = model.Count()
	if count == 0 {
		return
	}
	if err != nil {
		return
	}
	err = model.Fields("pile.*, station.name").Page(page.PageNo, page.PageSize).Scan(&piles)
	return
}

// 获取充电桩数量
func (*PileLogic) PileTotal(ctx context.Context, query do.Pile) (count int, err error) {
	model := dao.Pile.Ctx(ctx).LeftJoin("station", "station.id = pile.station_id")
	if query.StationId != nil {
		model = model.Where("pile.station_id = ?", query.StationId)
	}
	if query.Code != nil {
		model = model.Where("pile.code like ?", fmt.Sprintf("%%%s%%", query.Code))
	}
	if query.State != nil {
		model = model.Where("pile.state = ?", query.State)
	}
	count, err = model.Count()
	return
}

// 查询订单并关联充电站信息
// func QueryOrders(ctx context.Context) (gdb.Result, error) {
//     return dao.Pile.Ctx(ctx).Join("user u", "u.id = order.user_id").Fields("order.*, u.name").Select()
// }

func (*PileLogic) Add(ctx context.Context, pile do.Pile) (err error) {
	pile.CreateAt, pile.UpdateAt = gtime.Now(), gtime.Now()
	_, err = dao.Pile.Ctx(ctx).Insert(pile)
	return
}

func (*PileLogic) Del(ctx context.Context, id int) (err error) {
	_, err = dao.Pile.Ctx(ctx).Delete("id = ?", id)
	return
}

func (*PileLogic) Update(ctx context.Context, pile do.Pile) (err error) {
	pile.UpdateAt = gtime.Now()
	_, err = dao.Pile.Ctx(ctx).Update(pile, "id = ?", pile.Id)
	return
}
