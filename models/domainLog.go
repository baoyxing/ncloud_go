package models

type DomainLog struct {
	/* 下载链接 (Optional) */
	Url string `json:"url"`
	/* 文件名 (Optional) */
	FileName string `json:"fileName"`
	/* 文件大小 (Optional) */
	Size int64 `json:"size"`
}
