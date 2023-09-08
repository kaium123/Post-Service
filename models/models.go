package models

type Attachment struct {
	Path string `json:"path"`
	Name string `json:"name"`
}

type Post struct {
	ID          int          `json:"id"`
	Content     string       `json:"content" `
	Attachments []Attachment `json:"attachments,omitempty"`
	UserID      int          `json:"user_id"`
}

type React struct {
	ID            int    `json:"id"`
	PostID        uint   `json:"post_id"`
	ReactType     string `json:"react_type"`
	ReactedUserID uint   `json:"reacted_user_id"`
	PostType      string `json:"post_type"`
}

type Comment struct {
	ID                 int    `json:"id"`
	PostID             uint   `json:"post_id"`
	Content            string `json:"content"`
	CommentAttachments []int  `json:"attachments" gorm:"type:integer[]"`
	ParentCommentID    uint   `json:"parent_comment_id"`
	Rank               uint   `json:"rank"`
}

type Share struct {
	ID       int  `json:"id"`
	PostID   uint `json:"post_id"`
	SharedID int  `json:"shared_id"`
}
