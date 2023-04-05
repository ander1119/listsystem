package models

// Articles here should be created as new struct, which contains {id, title, content, author etc...}
// I think it's quite trivial so I just skip this work
type Page struct {
	PageKey		 string `json:"page_key" gorm:"primaryKey"`
	Articles	 string `json:"articles"`
	NextPageKey *string `json:"next_page_key,omitempty"`
	CreateAt  	  int64 `json:"create_at" gorm:"index"`
	UpdateAt  	  int64 `json:"update_at"`
}

type Head struct {
	ListKey		string `json:"list_key" gorm:"primaryKey"`
	NextPageKey string `json:"next_page_key"`
	CreateAt 	 int64 `json:"create_at" gorm:"index"`
	UpdateAt 	 int64 `json:"update_at"`
}

type HeadInput struct {
	ListKey		string `json:"list_key" validate:"required"`
	NextPageKey string `json:"next_page_key" validate:"required"`
}

type PageInput struct {
	Articles	string `json:"articles" validate:"required"`
	NextPageKey string `json:"next_page_key" validate:"omitempty"`
}
