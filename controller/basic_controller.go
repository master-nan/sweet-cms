package controller

import (
	"bytes"
	"fmt"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"go.uber.org/zap"
	"log"
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
	"vitess.io/vitess/go/vt/sqlparser"
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
	err := utils.ValidatorBody[request.SignInReq](ctx, &data, translator)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	configUre, err := b.sysConfigureService.Query()
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	if configUre.EnableCaptcha {
		boolean := captcha.VerifyString(data.CaptchaId, data.Captcha)
		if boolean == false {
			e := &response.AdminError{
				ErrorCode:    http.StatusBadRequest,
				ErrorMessage: "验证码错误",
				Success:      false,
			}
			_ = ctx.Error(e)
			return
		}
	}
	var loginLog = model.LoginLog{
		Ip:       ctx.ClientIP(),
		Locality: "",
		UserName: data.UserName,
	}
	// 异步保存登录日志
	go func(loginLog model.LoginLog) {
		e := b.logService.CreateLoginLog(loginLog)
		if e != nil {
			zap.L().Error("login loginLog err", zap.Error(err))
		}
	}(loginLog)
	user, err := b.sysUserService.GetByUserName(data.UserName)
	if err != nil || utils.Encryption(data.Password, b.serverConfig.Config.Salt) != user.Password || !user.State {
		e := &response.AdminError{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: "用户名或密码错误",
		}
		_ = ctx.Error(e)
		return
	} else {
		token, err := b.tokenGenerator.GenerateToken(strconv.Itoa(user.Id))
		if err != nil {
			_ = ctx.Error(err)
			return
		} else {
			go func() {
				var up request.UserUpdateReq
				up.Id = user.Id
				up.AccessTokens = utils.UpdateAccessTokens(user.AccessTokens, token)
				up.GmtLastLogin = model.CustomTime(time.Now())
				err := b.sysUserService.Update(ctx, up)
				zap.L().Error("login update err", zap.Error(err))
			}()
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

func (b *BasicController) Captcha(ctx *gin.Context) {
	l := captcha.DefaultLen
	w, h := 110, 50
	captchaId := captcha.NewLen(l)
	var content bytes.Buffer
	_ = captcha.WriteImage(&content, captchaId, w, h)
	imageData := content.Bytes()
	// 返回JSON数据，包含captchaId和图片的base64编码
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
		_ = ctx.Error(err)
		return
	}
	resp.SetData(configUre)
	return
}

func (b *BasicController) Logout(ctx *gin.Context) {
	utils.DeleteSession(ctx, "captcha")
}

func (b *BasicController) Test(ctx *gin.Context) {
	sql := "SELECT u.id as user_id, u.name as username, o.order_id FROM users u JOIN orders o ON u.id = o.user_id"
	parser := sqlparser.NewTestParser()
	stmt, err := parser.Parse(sql)
	if err != nil {
		// 处理解析错误
		log.Fatalf("Error parsing SQL: %v", err)
	}

	// 断言 *sqlparser.Select 类型来访问选择语句的细节
	selectStmt, ok := stmt.(*sqlparser.Select)
	if !ok {
		log.Fatalf("Not a SELECT statement: %v", stmt)
	}

	// 遍历解析的树来提取字段信息
	// sqlparser.String() 可以将节点转回 SQL 字符串
	for _, selectExpr := range selectStmt.SelectExprs {
		if aliasedExpr, ok := selectExpr.(*sqlparser.AliasedExpr); ok {
			if colName, ok := aliasedExpr.Expr.(*sqlparser.ColName); ok {
				fmt.Printf("Column: %v, Table: %v\n", colName.Name, colName.Qualifier)
			}
		}
	}
}
