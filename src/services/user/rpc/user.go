package rpc

import (
	"context"
	"go-video/src/common/rpcinterface/userrpcinterface"
	"go-video/src/common/universalresp"
	"go-video/src/common/utils"
	"go-video/src/services/user/db/repository"
)

//UserHandler for rpc interface
type UserHandler struct {
}

//UserRegister for regitser user
func (*UserHandler) UserRegister(ctx context.Context,
	req *userrpcinterface.RegisterReq,
	resp *userrpcinterface.ReigisterResp) error {

	if repository.ExistUserByUsername(req.Username) {
		resp.Code = universalresp.RespCodeUserAlreadyRegistered.Code
		resp.Message = universalresp.RespCodeUserAlreadyRegistered.Message
		return nil
	}

	password := utils.GenHashWithMD5(req.Password)

	err := repository.InsertUser(req.Username, password)
	if err != nil {
		resp.Code = universalresp.RespCodeUserRegisterError.Code
		resp.Message = universalresp.RespCodeUserRegisterError.Message
		return err
	}

	resp.Code = universalresp.RespCodeSuccess.Code
	resp.Message = universalresp.RespCodeSuccess.Message

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
