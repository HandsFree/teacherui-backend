package entity

import "time"

type Resource struct {
	ID            uint64    `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	FileName      string    `json:"fileName"`
	ContentType   string    `json:"contentType"`
	ContentLength uint64    `json:"contentLength"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	Owner         string    `json:"owner"`
	OwnedByMe     bool      `json:"ownedByMe"`
	ReadOnly      bool      `json:"readOnly"`
}

type ResourceLink struct {
	GlpID      string `json:"gamifiedLessonPathId"`
	ResourceID string `json:"resourceId"`
}

// The API takes a string when posting
// but returns an integer on GET.
type ResourceLinkResponse struct {
	GlpID      uint64 `json:"gamifiedLessonPathId"`
	ResourceID uint64 `json:"resourceId"`
}

type ResourceData struct {
	ID      uint64 `json:"id"`
	Content string `json:"data"`
}
