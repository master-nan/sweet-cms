/**
 * @Author: Nan
 * @Date: 2023/3/18 22:48
 */

package model

type AccessLog struct {
	Basic
	Method   string `gorm:"size:64;comment:操作" json:"method"`
	Ip       string `gorm:"size:128;comment:ip"json:"ip"`
	Locality string `gorm:"size:128;comment:用户名" json:"locality"`
	Url      string `gorm:"size:512;comment:路径" json:"url"`
	Body     string `gorm:"type:text;comment:请求数据" json:"body"`
	Query    string `gorm:"type:text;comment:查询" json:"query"`
	Response string `gorm:"type:text;comment:相应数据" json:"response"`
}

type LoginLog struct {
	Basic
	Ip       string `gorm:"size:128;comment:用户名" json:"ip"`
	Locality string `gorm:"size:128;comment:用户名" json:"locality"`
	UserName string `gorm:"size:128;comment:用户名" json:"userName"`
}
