/**
 * @Author: Nan
 * @Date: 2024/7/3 下午12:03
 */

package response

import "sweet-cms/model"

type UserRes struct {
	ID           int              `json:"id"`
	State        bool             `json:"state"`
	UserName     string           `json:"username"`
	Email        string           `json:"email"`
	PhoneNumber  string           `json:"phone_number"`
	IDCard       string           `json:"id_card"`
	EmployeeID   int              `json:"employee_id"`
	GmtLastLogin model.CustomTime `json:"gmt_last_login"`
	Language     string           `json:"language"`
}
