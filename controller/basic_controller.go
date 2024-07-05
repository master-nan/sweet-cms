package controller

import (
	"bytes"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"strconv"
	"strings"
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
	translators         map[string]ut.Translator
}

func NewBasicController(tokenGenerator inter.TokenGenerator, serverConfig *config.Server, sysConfigureService *service.SysConfigureService, logService *service.LogService, sysUserService *service.SysUserService, translators map[string]ut.Translator) *BasicController {
	return &BasicController{
		tokenGenerator,
		serverConfig,
		sysConfigureService,
		logService,
		sysUserService,
		translators,
	}
}

func (b *BasicController) Login(ctx *gin.Context) {
	var data request.SignInReq
	resp := response.NewResponse()
	ctx.Set("response", resp)
	translator, _ := b.translators["zh"]
	if err := ctx.ShouldBindBodyWith(&data, binding.JSON); err != nil {
		if err == io.EOF {
			// 客户端请求体为空
			e := &response.AdminError{
				Code:    http.StatusBadRequest,
				Message: "请求参数数据",
			}
			ctx.Error(e)
			return
		}
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			// 如果是验证错误，则翻译错误信息
			var errorMessages []string
			for _, e := range ve {
				errMsg := e.Translate(translator)
				errorMessages = append(errorMessages, errMsg)
			}
			e := &response.AdminError{
				Code:    http.StatusBadRequest,
				Message: strings.Join(errorMessages, ","),
			}
			ctx.Error(e)
			return
		}
		ctx.Error(err)
		return
	} else {
		configUre, err := b.sysConfigureService.Query()
		if err != nil {
			ctx.Error(err)
			return
		}
		if configUre.EnableCaptcha {
			//captchaId := utils.GetSessionString(ctx, "captcha")
			boolean := captcha.VerifyString(data.CaptchaId, data.Captcha)
			if boolean == false {
				e := &response.AdminError{
					Code:    http.StatusBadRequest,
					Message: "验证码错误",
				}
				ctx.Error(e)
				return
			}
		}
		var log = model.LoginLog{
			Ip:       ctx.ClientIP(),
			Locality: "",
			UserName: data.UserName,
		}
		err = b.logService.CreateLoginLog(log)
		user, err := b.sysUserService.GetByUserName(data.UserName)
		if err != nil || utils.Encryption(data.Password, b.serverConfig.Config.Salt) != user.Password || !user.State {
			e := &response.AdminError{
				Code:    http.StatusBadRequest,
				Message: "用户名或密码错误",
			}
			ctx.Error(e)
			return
		} else {
			token, err := b.tokenGenerator.GenerateToken(strconv.Itoa(user.Id))
			if err != nil {
				ctx.Error(err)
				return
			} else {
				var up request.UserUpdateReq
				up.Id = user.Id
				up.AccessTokens = utils.UpdateAccessTokens(user.AccessTokens, token)
				up.GmtLastLogin = model.CustomTime(time.Now())
				err = b.sysUserService.Update(up)
				if err != nil {
					ctx.Error(err)
					return
				}
				var userRes response.UserRes
				utils.Assignment(&user, &userRes)
				signInRes := response.SignInRes{
					AccessToken: token,
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
	//utils.SaveSession(ctx, "captcha", captchaId)
	var content bytes.Buffer
	//ctx.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	//ctx.Writer.Header().Set("Pragma", "no-cache")
	//ctx.Writer.Header().Set("Expires", "0")
	//ctx.Writer.Header().Set("Content-Type", "image/png")
	_ = captcha.WriteImage(&content, captchaId, w, h)
	//http.ServeContent(ctx.Writer, ctx.Request, captchaId+".png", time.Time{}, bytes.NewReader(content.Bytes()))
	imageData := content.Bytes()
	// 返回JSON数据，包含captchaId和图片的base64编码
	//ctx.JSON(http.StatusOK, gin.H{
	//	"captchaId": captchaId,
	//	"image":     imageData,
	//})
	resp := response.NewResponse()
	ctx.Set("response", resp)
	resp.SetData(gin.H{
		"captchaId": captchaId,
		"image":     imageData,
	})
}

func (b *BasicController) Configure(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	configUre, err := b.sysConfigureService.Query()
	if err != nil {
		ctx.Error(err)
		return
	}
	resp.SetData(configUre)
	return
}

func (b *BasicController) Logout(ctx *gin.Context) {
	utils.DeleteSession(ctx, "captcha")
}
