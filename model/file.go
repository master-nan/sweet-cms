/**
 * @Author: Nan
 * @Date: 2024/8/5 下午11:47
 */

package model

type File struct {
	Basic
	FileName string `gorm:"size:128;comment:文件名" json:"fileName"`
	FilePath string `gorm:"size:512;comment:文件路径" json:"filePath"`
	FileType string `gorm:"size:128;comment:文件类型" json:"fileType"`
	FileUrl  string `gorm:"size:256;comment:文件地址" json:"fileUrl"`
	FileSize int    `gorm:"comment:文件大小" json:"fileSize"`
	FileMd5  string `gorm:"size:128;comment:文件md5" json:"fileMd5"`
	FileExt  string `gorm:"size:128;comment:文件扩展名" json:"fileExt"`
	FileHash string `gorm:"size:128;comment:文件hash" json:"fileHash"`
	FileUuid string `gorm:"size:128;comment:文件uuid" json:"fileUuid"`
}
