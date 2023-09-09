package controller_wx

import (
	"context"
	v2 "login-demo/api/wx"
	"login-demo/internal/logic"
	"login-demo/internal/model/do"
)

type WxUserController struct {
}

// 我的个人信息
func (c WxUserController) AboutMe(ctx context.Context, req *v2.WXMeInfoReq) (res v2.WXMeInfoRes, err error) {
	// 获得当前用户
	currentUser, err := logic.User.GetCurrentUser(ctx)
	if err != nil {
		return
	}
	res.Nickname, res.Balance, res.AvatarUrl, res.Username = currentUser.Nickname, currentUser.Balance, currentUser.AvatarUrl, currentUser.Passport
	return
}

// 个人头像上传
func (WxUserController) UploadAvatar(ctx context.Context, req *v2.WXUploadAvatarReq) (res *v2.WXUploadAvatarRes, err error) {
	err, _, fileUrl := logic.File.FileUpload(ctx, req.File)
	if err != nil {
		return
	}
	// 获得当前用户
	currentUser, err := logic.User.GetCurrentUser(ctx)
	if err != nil {
		return
	}
	err = logic.WxUser.Update(ctx, do.WxUser{
		UserId:    currentUser.Id,
		AvatarUrl: fileUrl,
	})
	return
}

// 修改个人昵称
func (WxUserController) UpdateNickname(ctx context.Context, req *v2.WXUpdateNicknameReq) (res *v2.WXUpdateNicknameRes, err error) {
	// 获得当前用户
	currentUser, err := logic.User.GetCurrentUser(ctx)
	if err != nil {
		return
	}
	err = logic.User.Update(ctx, do.User{
		Id:       currentUser.Id,
		Nickname: req.NewNickname,
	})
	return
}
