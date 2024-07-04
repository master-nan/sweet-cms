/**
 * @Author: Nan
 * @Date: 2024/6/28 下午3:37
 */

package request

import "sweet-cms/model"

type UserCreateReq struct {
	UserName    string `json:"userName" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Email       string `json:"email" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	IdCard      string `json:"idCard" binding:"required"`
	EmployeeId  int    `json:"employeeId" binding:"required"`
}

type UserUpdateReq struct {
	Id           int              `json:"id" binding:"required"`
	UserName     string           `json:"userName" binding:"required"`
	Password     string           `json:"password"`
	Email        string           `json:"email" binding:"required"`
	PhoneNumber  string           `json:"phoneNumber" binding:"required"`
	IdCard       string           `json:"idCard" binding:"required"`
	AccessTokens string           `json:"accessTokens"`
	GmtLastLogin model.CustomTime `json:"gmtLastLogin"`
}
