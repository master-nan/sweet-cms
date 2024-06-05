package controller

import (
	"bytes"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"strconv"
	"sweet-cms/config"
	"sweet-cms/form/request"
	"sweet-cms/form/response"
	"sweet-cms/inter"
	"sweet-cms/model"
	"sweet-cms/service"
	"sweet-cms/utils"
	"time"
)

type BasicController struct {
	tokenGenerator      inter.TokenGenerator
	serverConfig        *config.Server
	sysConfigureService *service.SysConfigureService
	logService          *service.LogService
	sysUserService      *service.SysUserService
}

func NewBasicController(tokenGenerator inter.TokenGenerator, serverConfig *config.Server, sysConfigureService *service.SysConfigureService, logService *service.LogService, sysUserService *service.SysUserService) *BasicController {
	return &BasicController{
		tokenGenerator,
		serverConfig,
		sysConfigureService,
		logService,
		sysUserService,
	}
}

func (b *BasicController) Login(ctx *gin.Context) {
	var data request.SignInReq
	resp := response.NewResponse()
	ctx.Set("response", resp)
	if err := ctx.ShouldBindBodyWith(&data, binding.JSON); err != nil {
		resp.SetErrorMessage(err.Error()).SetErrorCode(http.StatusBadRequest)
		return
	} else {
		configUre, err := b.sysConfigureService.Query()
		if err != nil {
			resp.SetErrorMessage(err.Error()).SetErrorCode(http.StatusInternalServerError)
			return
		}
		if configUre.EnableCaptcha {
			captchaId := utils.GetSessionString(ctx, "captcha")
			boolean := captcha.VerifyString(captchaId, data.Captcha)
			if boolean == false {
				resp.SetErrorMessage("验证码错误").SetErrorCode(http.StatusUnauthorized)
				return
			}
		}
		var log = model.LoginLog{
			Ip:       ctx.ClientIP(),
			Locality: "",
			Username: data.Username,
		}
		err = b.logService.CreateLoginLog(log)
		user, err := b.sysUserService.GetByUserName(data.Username)
		if err != nil || utils.Encryption(data.Password, b.serverConfig.Config.Salt) != user.Password {
			resp.SetErrorMessage("用户名或密码错误").SetErrorCode(http.StatusBadRequest)
			return
		} else {
			token, err := b.tokenGenerator.GenerateToken(strconv.Itoa(user.ID))
			if err != nil {
				resp.SetErrorMessage(err.Error()).SetErrorCode(http.StatusBadRequest)
			} else {
				signInRes := response.SignInRes{
					AccessToken: token,
					UserInfo:    user,
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
	configUre, err := b.sysConfigureService.Query()
	resp := response.NewResponse()
	ctx.Set("response", resp)
	if err != nil {
		resp.SetErrorMessage(err.Error()).SetErrorCode(http.StatusUnauthorized)
		return
	}
	resp.SetData(configUre)
	return
}

func (b *BasicController) Logout(ctx *gin.Context) {
	utils.DeleteSession(ctx, "captcha")
}
