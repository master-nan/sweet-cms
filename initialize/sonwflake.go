/**
 * @Author: Nan
 * @Date: 2023/9/12 15:55
 */

package initialize

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"sweet-cms/config"
	"sweet-cms/utils"
)

func InitSnowflake(serverConfig *config.Server) (*utils.Snowflake, error) {
	workerID := serverConfig.WorkerID // 从配置文件读取默认值
	podName := os.Getenv("WORKER_ID")
	if podName != "" {
		re := regexp.MustCompile(`-(\d+)$`)
		matches := re.FindStringSubmatch(podName)
		if len(matches) != 2 {
			return nil, fmt.Errorf("invalid pod name format: %s", podName)
		}
		id, err := strconv.ParseInt(matches[1], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid WORKER_ID: %v", err)
		}
		workerID = id
	}
	return utils.NewSnowflake(workerID)
}
