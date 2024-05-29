/**
 * @Author: Nan
 * @Date: 2024/5/21 下午2:46
 */

package request

type ConfigureUpdateReq struct {
	ID            int  `json:"id" binding:"required"`
	EnableCaptcha bool `json:"enable_captcha" binding:"required"`
}
