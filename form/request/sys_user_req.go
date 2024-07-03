/**
 * @Author: Nan
 * @Date: 2024/6/28 下午3:37
 */

package request

type UserCreateReq struct {
	UserName    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Email       string `json:"email" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	IDCard      string `json:"id_card" binding:"required"`
	EmployeeID  int    `json:"employee_id" binding:"required"`
}

type UserUpdateReq struct {
	ID           int     `json:"id" binding:"required"`
	UserName     *string `json:"username" binding:"required"`
	Password     *string `json:"password" binding:"required"`
	Email        *string `json:"email" binding:"required"`
	PhoneNumber  *string `json:"phone_number" binding:"required"`
	IDCard       *string `json:"id_card" binding:"required"`
	AccessTokens *string `json:"access_tokens"`
}
