package auth

import "time"

type User struct {
	ID        uint       `json:"id" bson:"id"`
	CreatedAt time.Time  `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" bson:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
	Username  string     `json:"username" bson:"username"`
	Password  string     `json:"password" bson:"password"`
}

type LoginRequestBody struct {
	Username string `json:"username" bson:"username" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
}
type LoginUserSuccessResponse struct {
	Token string `json:"token"`
}

type AuthFailResponse struct {
	Message string `json:"message" bson:"message"`
}
