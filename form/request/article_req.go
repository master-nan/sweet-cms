package request

import (
	"sweet-cms/enum"
	"time"
)

type ArticleBasicCreateReq struct {
	Title        string           `json:"title" binding:"required"`
	Cover        string           `json:"cover"`
	Introduction string           `json:"introduction"`
	IsAd         bool             `json:"is_ad"`
	Status       enum.ArticleType `json:"status"`
	GmtReview    time.Time        `json:"gmt_review"`
	RejectReason string           `json:"reject_reason"`
	IsComment    bool             `json:"is_comment"`
	ChannelId    int              `json:"channel_id" `
}

type ArticleBasicQueryReq struct {
	Page         int               `form:"page"`
	Num          int               `form:"num"`
	Title        *string           `form:"title"`
	Introduction *string           `form:"introduction"`
	IsAd         *bool             `form:"is_ad"`
	Status       *enum.ArticleType `form:"status"`
	IsComment    *bool             `form:"is_comment"`
	ChannelId    *int              `form:"channel_id"`
}

type ArticleBasicUpdateReq struct {
	Title        *string          `json:"title"`
	Introduction *string          `json:"introduction"`
	IsAd         *bool            `json:"is_ad"`
	Status       enum.ArticleType `json:"status"`
	IsComment    *bool            `json:"is_comment"`
	ChannelId    *int             `json:"channel_id"`
}

type ArticleChannelCreateReq struct {
	PID       *int   `json:"pid"`
	Name      string `json:"name"  binding:"required"`
	Sequence  *uint8 `json:"sequence"`
	IsVisible *bool  `json:"is_visible"`
}

type ArticleChannelUpdateReq struct {
	PID       *int   `json:"pid"`
	Name      string `json:"name"`
	Sequence  *uint8 `json:"sequence"`
	IsVisible *bool  `json:"is_visible"`
}

type ArticleChannelQueryReq struct {
	Page      int     `json:"pageSize" form:"pageSize"`
	Num       int     `json:"pageNum" form:"pageNum"`
	Name      *string `json:"name" form:"name"`
	PID       *int    `json:"pid" form:"pid"`
	IsVisible *bool   `json:"isVisible" form:"isVisible"`
}

type ArticleContentCreateReq struct {
	ArticleUUID string `json:"article_uuid" binding:"required"`
	Content     string `json:"content" binding:"required"`
}

type ArticleContentUpdateReq struct {
	Content string `json:"content" binding:"required"`
}
