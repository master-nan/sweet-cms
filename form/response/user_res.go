/**
 * @Author: Nan
 * @Date: 2024/7/3 下午12:03
 */

package response

import "sweet-cms/model"

type UserRes struct {
	Id           int              `json:"id"`
	State        bool             `json:"state"`
	UserName     string           `json:"userName"`
	Email        string           `json:"email"`
	PhoneNumber  string           `json:"phoneNumber"`
	IdCard       string           `json:"idCard"`
	EmployeeId   int              `json:"employeeId"`
	GmtLastLogin model.CustomTime `json:"gmtLastLogin"`
	Language     string           `json:"language"`
}
