package models

import (
	"encoding/json"
)

type Input struct {
	Data   Data   `json:"data"`
	Layout Layout `json:"layout"`
}

type Data struct {
	ID      string          `json:"id_data"`
	JSONRaw json.RawMessage `json:"data"`
	parsed  map[string]interface{}
}

type Layout struct {
	Code  string `json:"layout_code"`
	Pages []Page `json:"layout_pages"`
}

type Page struct {
	ID       string    `json:"page_id"`
	Order    int       `json:"page_order"`
	Sections []Section `json:"page_sections"`
}

type Section struct {
	ID       string    `json:"section_id"`
	Order    int       `json:"section_order"`
	Contents []Content `json:"section_contents"`
}

type Content struct {
	ID    string `json:"content_id"`
	Order int    `json:"content_order"`
	Value string `json:"value"`
}

func (data *Data) ParseRaw() {
	parsed := make(map[string]interface{})
	json.Unmarshal(data.JSONRaw, &parsed)
	data.parsed = parsed
}

func (data *Data) GetParsed() map[string]interface{} {
	return data.parsed
}
