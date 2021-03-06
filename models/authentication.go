package models

type Authentication struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Role     string `json:"role" bson:"role"`
}

type Authorization struct {
	Code     int    `json:"code" bson:"code"`
	Token    string `json:"token" bson:"token"`
	Expire   string `json:"expire" bson:"expire"`
	Username string `json:"username" bson:"username"`
	Role     string `json:"role" bson:"role"`
}
