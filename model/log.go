/**
 * @Author: Nan
 * @Date: 2023/3/18 22:48
 */

package model

type AccessLog struct {
	Basic
	Method   string `json:"method"`
	Ip       string `json:"ip"`
	Locality string `json:"locality"`
	Url      string `json:"url"`
	Data     string `json:"data"`
}

type LoginLog struct {
	Basic
	Ip       string `json:"ip"`
	Locality string `json:"locality"`
	Username string `json:"username"`
}
