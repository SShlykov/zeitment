package models

type WebResponse[T any] struct {
	Status string        `json:"status"`
	Data   T             `json:"data"`
	Paging *PageMetadata `json:"paging,omitempty"`
	Errors string        `json:"errors,omitempty"`
}

type PageResponse[T any] struct {
	Data         []T          `json:"data,omitempty"`
	PageMetadata PageMetadata `json:"paging,omitempty"`
}

type PageMetadata struct {
	Page      int   `json:"page"`
	PageSize  int   `json:"page_size"`
	TotalItem int64 `json:"total_item"`
	TotalPage int64 `json:"total_page"`
}

type PageOptions struct {
	Page     uint64 `json:"page,omitempty"`
	PageSize uint64 `json:"page_size,omitempty"`
}

func (po *PageOptions) GetPageAndPageSize() (uint64, uint64) {
	return po.Page, po.PageSize
}
