/**
 * @Author: Nan
 * @Date: 2023/3/18 22:48
 */

package model

type AccessLog struct {
	BasicModel
	Method   string `json:"method"`
	Ip       string `json:"ip"`
	Locality string `json:"locality"`
	Url      string `json:"url"`
	Data     string `json:"data"`
}

type LoginLog struct {
	BasicModel
	Ip       string `json:"ip"`
	Locality string `json:"locality"`
	Username string `json:"username"`
}
