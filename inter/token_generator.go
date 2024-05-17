/**
 * @Author: Nan
 * @Date: 2024/5/17 上午10:26
 */

package inter

type TokenGenerator interface {
	GenerateToken(id string) (string, error)
}
