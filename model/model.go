package model

type UserCredential struct {
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
}

type VideoInfo struct {
	Id           string
	AuthorId     int
	Name         string
	DisplayCtime string
}
