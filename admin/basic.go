package admin

import (
	"bytes"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"strconv"
	"sweet-cms/form/request"
	"sweet-cms/form/response"
	"sweet-cms/global"
	"sweet-cms/inter"
	"sweet-cms/model"
	"sweet-cms/server"
	"sweet-cms/utils"
	"time"
)

type Basic struct {
	TokenGenerator inter.TokenGenerator
}

func NewBasic() *Basic {
	return &Basic{
		TokenGenerator: utils.NewJWTTokenGen(),
	}
}

func (b *Basic) Login(ctx *gin.Context) {
	var data request.SignInReq
	rsp := response.NewRespData(ctx)
	if err := ctx.ShouldBindBodyWith(&data, binding.JSON); err != nil {
		rsp.SetMsg(err.Error()).SetCode(http.StatusBadRequest).ReturnJson()
	} else {
		captchaId := utils.GetSessionString(ctx, "captcha")
		boolean := captcha.VerifyString(captchaId, data.Captcha)
		if boolean == false {
			rsp.SetMsg("验证码错误").SetCode(http.StatusUnauthorized).ReturnJson()
		}
		logServer := server.NewLogServer(ctx)
		var log = model.LoginLog{
			Ip:       ctx.ClientIP(),
			Locality: "",
			Username: data.Username,
		}
		_, err := logServer.CreateLoginLog(log)
		user, err := server.NewSysServer().GetSysUser(data.Username)
		if err != nil || utils.Encryption(data.Password, global.ServerConf.Configure.Salt) != user.Password {
			rsp.SetMsg("用户名或密码错误").SetCode(http.StatusBadRequest).ReturnJson()
		} else {
			token, err := b.TokenGenerator.GenerateToken(strconv.Itoa(user.ID))
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

func (b *Basic) Captcha(ctx *gin.Context) {
	l := captcha.DefaultLen
	w, h := 110, 50
	captchaId := captcha.NewLen(l)
	utils.SaveSession(ctx, "captcha", captchaId)
	var content bytes.Buffer
	ctx.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	ctx.Writer.Header().Set("Pragma", "no-cache")
	ctx.Writer.Header().Set("Expires", "0")
	ctx.Writer.Header().Set("Content-Type", "image/png")
	_ = captcha.WriteImage(&content, captchaId, w, h)
	http.ServeContent(ctx.Writer, ctx.Request, captchaId+".png", time.Time{}, bytes.NewReader(content.Bytes()))
}
