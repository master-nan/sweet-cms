/**
 * @Author: Nan
 * @Date: 2024/8/1 下午10:32
 */

package request

type SysRoleMenuCreateReq struct {
	RoleId int `json:"roleId" binding:"required"`
	MenuId int `json:"menuId" binding:"required"`
}
