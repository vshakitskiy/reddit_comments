package model

type ConnectionMeta struct {
	TotalRow   int32 `json:"totalRow"`
	TotalPages int32 `json:"totalPages"`
	Limit      int32 `json:"limit"`
	Page       int32 `json:"page"`
}

type CommentsConnection struct {
	Meta *ConnectionMeta `json:"meta,omitempty"`
	Rows []*Comment      `json:"rows,omitempty"`
}

type PostsConnection struct {
	Meta *ConnectionMeta `json:"meta,omitempty"`
	Rows []*Post         `json:"rows,omitempty"`
}
