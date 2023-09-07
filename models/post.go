package models

type Attachment struct {
	Path string `json:"path"`
	Name string `json:"name"`
}

type Post struct {
	ID            int    `json:"id"`
	Content       string `json:"content" `
	AttachmentIDs []int  `json:"attachments_ids,omitempty"`
	UserID        int    `json:"user_id"`
}
