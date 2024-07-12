/**
 * @Author: Nan
 * @Date: 2024/5/17 上午10:26
 */

package inter

type TokenGenerator interface {
	GenerateToken(id string) (token string, err error)
	ParseToken(token string) (id string, err error)
}
