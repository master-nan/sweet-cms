/**
 * @Author: Nan
 * @Date: 2023/3/14 21:10
 */

package request

type SignInReq struct {
	UserName  string `json:"userName"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
}
