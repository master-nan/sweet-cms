/**
 * @Author: Nan
 * @Date: 2023/3/14 16:55
 */

package response

type SignInRes struct {
	AccessToken string `json:"access_token"`
	//UserInfo    model.SysUser `json:"user_info"`
}
