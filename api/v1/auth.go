/**
 * @Author: Nan
 * @Date: 2023/3/12 23:32
 */

package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"strconv"
	"sweet-cms/form/request"
	"sweet-cms/form/response"
	"sweet-cms/global"
	"sweet-cms/middlewares"
	"sweet-cms/model"
	"sweet-cms/service"
	"sweet-cms/utils"
)

type AuthApi struct {
	TokenGenerator TokenGenerator
}
type TokenGenerator interface {
	GenerateToken(id string) (string, error)
}

func NewAuthApi() *AuthApi {
	return &AuthApi{
		TokenGenerator: utils.NewJWTTokenGen(),
	}
}

func (c *AuthApi) Login(ctx *gin.Context) {
	var data request.SignInReq
	resp := middlewares.NewResponse()
	if err := ctx.ShouldBindBodyWith(&data, binding.JSON); err != nil {
		resp.SetCode(http.StatusBadRequest).SetMsg(err.Error())
	} else {
		logServer := service.NewLogServer(ctx)
		var log = model.LoginLog{
			Ip:       ctx.ClientIP(),
			Locality: "",
			Username: data.Username,
		}
		_, err := logServer.CreateLoginLog(log)
		user, err := service.NewSysUserService().Get(data.Username)
		if err != nil || utils.Encryption(data.Password, global.ServerConf.Config.Salt) != user.Password {
			resp.SetMsg("用户名或密码错误").SetCode(http.StatusBadRequest)
		} else {
			token, err := c.TokenGenerator.GenerateToken(strconv.Itoa(user.ID))
			if err != nil {
				resp.SetMsg(err.Error()).SetCode(http.StatusBadRequest)
			} else {
				signInRes := response.SignInRes{
					AccessToken: token,
					//UserInfo:    user,
				}
				resp.SetData(signInRes)
			}
		}
	}

}
