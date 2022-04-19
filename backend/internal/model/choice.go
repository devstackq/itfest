package model

import "time"

type Choice struct {
	ID        int
	Key       string
	Title     string
	Value     string
	Fields    map[string]interface{} `json:"fields"`
	Template  `json:"template"`
	Tag       string
	TagValue  string
	CreatedAt time.Time
}

type Choice2 struct {
	ID          int
	Title       string
	Description string
	TH          []TableHeader `json:"theader"` // filled
	TV          []TableValue  `json:"tvalue"`  // empty val
	Template    `json:"template"`
	Tag         string
	CreatedAt   time.Time
	Fields      map[string]interface{} `json:"fields"`
}

type Header struct {
	Key   string
	Value string
}

type TableHeader struct {
	Key   string
	Value string
}

type TableValue struct {
	Key   string
	Value string
}

// {
// 	"key": "Description",
// 	"value": "",
// 	"title": "Описание проекта 2022",
// 	"template": {
// 		"id": 1
// 	},
// 	"tag": "<p> </p>"
// }
