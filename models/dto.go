package models

// LAYOUT DTO
type LayoutDTO struct {
	Code     string       `json:"code"`
	Pages    []PageDTO    `json:"pages"`
	Sections []SectionDTO `json:"sections"`
	Contents []ContentDTO `json:"contents"`
}

// PAGE DTO
type PageDTO struct {
	ID          string    `json:"page_id"`
	SectionsIDs *[]string `json:"sections"`
}

// SECTION DTO
type SectionDTO struct {
	ID          string    `json:"section_id"`
	ContentsIDs *[]string `json:"contents"`
}

// CONTENT DTO
type ContentDTO struct {
	ID    string `json:"content_id"`
	Value string `json:"value"`
}
