/**
 * @Author: Nan
 * @Date: 2023/3/14 21:10
 */

package request

type SignInReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Captcha  string `json:"captcha"`
}
