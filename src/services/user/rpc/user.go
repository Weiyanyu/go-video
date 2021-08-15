package rpc

import (
	"context"
	"go-video/src/common/rpcinterface/userrpcinterface"
)

//UserHandler for rpc interface
type UserHandler struct {
}

//UserRegister for regitser user
func (*UserHandler) UserRegister(ctx context.Context,
	req *userrpcinterface.RegisterReq,
	resp *userrpcinterface.ReigisterResp) error {

	//TODO:imp
	return nil
}

//QueryUserInfo for query user info, and than send to user
func (*UserHandler) QueryUserInfo(ctx context.Context,
	req *userrpcinterface.QueryUserInfoReq,
	resp *userrpcinterface.QueryUserInfoResp) error {

	//TODO:imp
	return nil
}

//UpdateUserProfile for update user profile
func (*UserHandler) UpdateUserProfile(ctx context.Context,
	req *userrpcinterface.UpdateUserProfileReq,
	resp *userrpcinterface.UpdateUserProfileResp) error {

	//TODO:imp
	return nil
}

//ChangeUserPassword for change user password request
func (*UserHandler) ChangeUserPassword(ctx context.Context,
	req *userrpcinterface.ChangeUserPasswordReq,
	resp *userrpcinterface.ChangeUserPasswordResp) error {

	//TODO:imp
	return nil
}
