/**
 * @Author: Nan
 * @Date: 2023/3/29 13:42
 */

package model

import (
	"github.com/google/uuid"
	"time"
)

type QuestionnaireContent struct {
	BasicModel
	UUID     uuid.UUID `gorm:"type:uuid;column:uuid" json:"uuid" binding:"required"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Status   string    `json:"status"`
	Creator  int       `json:"creator"`
	GmtStart time.Time `json:"gmtStart"`
	GmtEnd   time.Time `json:"gmtEnd"`
}

type QuestionnaireAnswer struct {
	BasicModel
	QuestionUUID uuid.UUID `gorm:"type:uuid;column:question_uuid" json:"questionUUID" binding:"required"`
	Content      string    `json:"content"`
	Creator      int       `json:"creator"`
}
