package model

type UserLogin struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Role     string `json:"role" bson:"role"`
}

type Payload struct {
	User string `json:"user"`
	Role string `json:"role"`
}