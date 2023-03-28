package models

import "time"

type Head struct {
	ListKey		string `json:"list_key"`
	NextPageKey string `json:"next_page_key"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

type Page struct {
	Articles	string `json:"article"`
	NextPageKey string `json:"next_page_key"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}