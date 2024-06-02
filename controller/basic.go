package controller

import (
	"bytes"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"strconv"
	"sweet-cms/cache"
	"sweet-cms/config"
	"sweet-cms/form/request"
	"sweet-cms/form/response"
	"sweet-cms/global"
	"sweet-cms/inter"
	"sweet-cms/middlewares"
	"sweet-cms/model"
	"sweet-cms/service"
	"sweet-cms/utils"
	"time"
)

type BasicController struct {
	TokenGenerator inter.TokenGenerator
	serverConfig   *config.Server
}

type TokenGenerator interface {
}

func NewBasicController(serverConfig *config.Server) *BasicController {
	return &BasicController{
		TokenGenerator: utils.NewJWTTokenGen(),
		serverConfig:   serverConfig,
	}
}

func (b *BasicController) Login(ctx *gin.Context) {
	var data request.SignInReq
	resp := middlewares.NewResponse()
	if err := ctx.ShouldBindBodyWith(&data, binding.JSON); err != nil {
		resp.SetMsg(err.Error()).SetCode(http.StatusBadRequest)
		return
	} else {
		captchaId := utils.GetSessionString(ctx, "captcha")
		boolean := captcha.VerifyString(captchaId, data.Captcha)
		if boolean == false {
			resp.SetMsg("验证码错误").SetCode(http.StatusUnauthorized)
			return
		}
		logServer := service.NewLogServer(ctx)
		var log = model.LoginLog{
			Ip:       ctx.ClientIP(),
			Locality: "",
			Username: data.Username,
		}
		_, err := logServer.CreateLoginLog(log)
		user, err := service.NewSysUserService().Get(data.Username)
		if err != nil || utils.Encryption(data.Password, b.serverConfig.Config.Salt) != user.Password {
			resp.SetMsg("用户名或密码错误").SetCode(http.StatusBadRequest)
			return
		} else {
			token, err := b.TokenGenerator.GenerateToken(strconv.Itoa(user.ID))
			if err != nil {
				resp.SetMsg(err.Error()).SetCode(http.StatusBadRequest)
			} else {
				signInRes := response.SignInRes{
					AccessToken: token,
					//UserInfo:    user,
				}
				resp.SetData(signInRes)
				return
			}
		}
	}
}

func (b *BasicController) Captcha(ctx *gin.Context) {
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

func (b *BasicController) Configure(ctx *gin.Context) {
	configureCache := cache.NewSysConfigureCache(service.NewConfigureService(), utils.NewRedisUtil(global.RedisClient))
	configUre, err := configureCache.Get("")
	resp := middlewares.NewResponse()
	ctx.Set("response", resp)
	if err != nil {
		resp.SetMsg(err.Error()).SetCode(http.StatusUnauthorized)
		return
	}
	resp.SetData(configUre)
	return
}

func (b *BasicController) Logout(ctx *gin.Context) {
	utils.DeleteSession(ctx, "captcha")
}
