package router

import (
	"context"
	"go-video/src/common/log4video"
	"go-video/src/common/rpcinterface/userrpcinterface"
	"go-video/src/common/universalresp"
	"go-video/src/services/api/rpc"
	"go-video/src/services/api/vo"
	"net/http"

	"github.com/gin-gonic/gin"
)

//RegisterUser is a handle of register user
func RegisterUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqVo vo.UserRegisterVO
		if err := c.ShouldBind(&reqVo); err != nil {
			log4video.E("request param error : %v", err)
			c.JSON(http.StatusBadRequest,
				universalresp.NewServiceResponse(universalresp.RespCodeBindReParamError))
			return
		}

		resp, err := rpc.UserServiceCli.UserRegister(context.TODO(), &userrpcinterface.RegisterReq{
			Username: reqVo.Username,
			Password: reqVo.Password,
		})

		if err != nil {
			log4video.E("rpc call(user service) error : %v", err)
			c.JSON(http.StatusInternalServerError,
				universalresp.NewServiceResponse(universalresp.RespCodeInteralError))
			return
		}

		if resp.Code != universalresp.RespCodeSuccess.Code {
			log4video.D("resp code is not suceess, maybe is bad request, please check. resp code is %d", resp.Code)
			c.JSON(http.StatusBadRequest,
				universalresp.NewServiceResponse(universalresp.ResponceCode{
					Code:    resp.Code,
					Message: resp.Message,
				}))
			return
		}

		c.JSON(http.StatusOK, universalresp.NewServiceResponse(universalresp.RespCodeSuccess))

	}
}
