package model

import uuid "github.com/satori/go.uuid"

type User struct {
	*Model
	UUID        uuid.UUID `json:"uuid"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	NickName    string    `json:"nick_name"`
	HeaderImage string    `json:"header_image"`
}
