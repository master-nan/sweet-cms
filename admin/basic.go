package admin

import (
	"bytes"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"net/http"
	"sweet-cms/global"
	"sweet-cms/server"
	"sweet-cms/utils"
	"time"
)

type Basic struct {
}

func NewBasic() *Basic {
	return &Basic{}
}

func (b *Basic) Login(ctx *gin.Context) {
	if ctx.Request.Method == "GET" {
		ctx.HTML(http.StatusOK, "login.html", nil)
		return
	} else if ctx.Request.Method == "POST" {
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")
		captchaVal := ctx.PostForm("captcha")
		captchaId := utils.GetSessionString(ctx, "captcha")
		boolean := captcha.VerifyString(captchaId, captchaVal)
		data := gin.H{
			"username": username,
			"password": password,
			"captcha":  captchaVal,
		}
		if boolean == false {
			data["errorMsg"] = "验证码错误"
			ctx.HTML(http.StatusOK, "login.html", data)
			return
		}
		user, err := server.NewSystemServer().GetSysUser(username)
		if err != nil || utils.Encryption(password, global.ServerConf.Configure.Salt) != user.Password {
			data["errorMsg"] = "用户名或密码错误"
			ctx.HTML(http.StatusOK, "login.html", data)
		} else {
			ctx.HTML(http.StatusOK, "index.html", nil)
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
