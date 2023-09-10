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
	Comments    []*Comment    `json:"comments"`
	Like        int          `json:"like"`
}

type React struct {
	ID            int    `json:"id"`
	PostID        uint   `json:"post_id"`
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

type User struct {
	ID             int      `json:"id"`
	Email          string   `json:"email" `
	Password       string   `json:"password,omitempty"`
	Name           string   `json:"name"`
	UserName       string   `json:"user_name" `
	Phone          string   `json:"phone"`
	Websites       []string `json:"websites"`
	Bio            string   `json:"bio"`
	Gender         string   `json:"gender"`
	ProfilePicName string   `json:"profile_pic_name"`
	ProfilePicPath string   `json:"profile_pic_path"`
}

type RequestParams struct {
	Keyword  string `json:"keyword"`
}
