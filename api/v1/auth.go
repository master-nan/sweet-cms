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
	"sweet-cms/model"
	"sweet-cms/server"
	"sweet-cms/utils"
)

type AuthApi struct {
}

func NewAuthApi() *AuthApi {
	return &AuthApi{}
}

func (c *AuthApi) Login(ctx *gin.Context) {
	var data request.SignInReq
	rsp := response.NewRespData(ctx)

	if err := ctx.ShouldBindBodyWith(&data, binding.JSON); err != nil {
		rsp.SetMsg(err.Error()).SetCode(http.StatusBadRequest).ReturnJson()
	} else {
		logServer := server.NewLogServer(ctx)
		var log = model.LoginLog{
			Ip:       ctx.ClientIP(),
			Locality: "",
			Username: data.Username,
		}
		_, err := logServer.CreateLoginLog(log)
		user, err := server.NewSystemServer().GetSysUser(data.Username)
		if err != nil || utils.Encryption(data.Password, global.ServerConf.Configure.Salt) != user.Password {
			rsp.SetMsg("用户名或密码错误").SetCode(http.StatusBadRequest).ReturnJson()
		} else {
			token, err := utils.NewJWTTokenGen().GenerateToken(strconv.Itoa(user.ID))
			if err != nil {
				rsp.SetMsg(err.Error()).SetCode(http.StatusBadRequest).ReturnJson()
			} else {
				signInRes := response.SignInRes{
					AccessToken: token,
					UserInfo:    user,
				}
				rsp.SetData(signInRes).ReturnJson()
			}
		}
	}

}
