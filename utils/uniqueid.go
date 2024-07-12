/**
 * @Author: Nan
 * @Date: 2023/9/12 15:53
 */

package utils

import (
	"fmt"
	"sync"
	"time"
)

const (
	workerIDBits = 5
	sequenceBits = 12
	maxWorkerID  = -1 ^ (-1 << workerIDBits)
	maxSequence  = -1 ^ (-1 << sequenceBits)
	epoch        = 1693497600000 // 2023-09-01 00:00:00 的毫秒时间戳
)

// Snowflake 结构体用于保存 Snowflake ID 的状态
type Snowflake struct {
	mu            sync.Mutex
	workerID      int64
	sequence      int64
	lastTimestamp int64
}

// NewSnowflake 创建一个新的 Snowflake 实例
func NewSnowflake(workerID int64) (*Snowflake, error) {
	if workerID < 0 || workerID > maxWorkerID {
		return nil, fmt.Errorf("worker ID 超出范围")
	}

	return &Snowflake{
		workerID: workerID,
	}, nil
}

// GenerateUniqueID 生成唯一 ID
func (s *Snowflake) GenerateUniqueID() (int64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 获取当前时间戳（毫秒级）
	timestamp := time.Now().UnixNano() / 1e6

	// 如果当前时间小于上一次生成的时间戳，等待时钟追赶上来
	if timestamp < s.lastTimestamp {
		return 0, fmt.Errorf("时钟回拨，无法生成唯一 ID")
	}

	// 如果是同一毫秒内生成的 ID，递增序列号
	if timestamp == s.lastTimestamp {
		s.sequence = (s.sequence + 1) & maxSequence
		if s.sequence == 0 {
			// 序列号溢出，等待下一毫秒
			timestamp = s.waitNextMillisecond(s.lastTimestamp)
		}
	} else {
		// 不同毫秒内生成的 ID，重置序列号
		s.sequence = 0
	}

	// 更新上一次生成 ID 的时间戳
	s.lastTimestamp = timestamp

	// 构建唯一 ID
	uniqueID := ((timestamp - epoch) << (workerIDBits + sequenceBits)) |
		(s.workerID << sequenceBits) |
		s.sequence

	return uniqueID, nil
}

// 等待下一毫秒的函数
func (s *Snowflake) waitNextMillisecond(lastTimestamp int64) int64 {
	timestamp := time.Now().UnixNano() / 1e6
	for timestamp <= lastTimestamp {
		timestamp = time.Now().UnixNano() / 1e6
	}
	return timestamp
}
