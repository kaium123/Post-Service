package models

type Attachment struct {
	Path string `json:"path"`
	Name string `json:"name"`
}

type User struct {
	Email      string  `json:"email" gorm:"unique"`
	Password   *string `json:"password,omitempty"`
	Name       string  `json:"name"`
	UserName   string  `json:"user_name" gorm:"index"`
	Phone      *string `json:"phone"`
	Website    *string `json:"website"`
	Bio        *string `json:"bio"`
	Gender     *string `json:"gender"`
	ProfilePic *string `json:"profile_pic"`
}

type SignInData struct {
	Email    string  `json:"email" gorm:"unique"`
	Password *string `json:"password,omitempty"`
}
