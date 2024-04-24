package dtos

import "time"

type SignInInput struct {
	Username string `json:"username" db:"username" binding:"required"`
	Password string `json:"password" db:"password" binding:"required"`
}

type Roles struct {
	RoleId   int
	RoleName string
}

type User struct {
	UserID    int       `json:"user_id" db:"user_id"`
	Username  string    `json:"username" db:"username" binding:"required"`
	Email     string    `json:"email" db:"email" binding:"required"`
	Password  string    `json:"password" db:"password_hash" binding:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type Post struct {
	PostID    int       `json:"post_id" db:"post_id"`
	UserID    int       `json:"user_id" db:"user_id"`
	Content   string    `json:"content" db:"content" binding:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type Comment struct {
	CommentID   int       `json:"comment_id" db:"comment_id"`
	UserID      int       `json:"user_id" db:"user_id"`
	PostID      int       `json:"post_id" db:"post_id"`
	ParentID    *int      `json:"parent_id,omitempty" db:"parent_id"`
	CommentText string    `json:"comment_text" db:"comment_text" binding:"required"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type Like struct {
	LikeID    int       `json:"like_id" db:"like_id"`
	UserID    int       `json:"user_id" db:"user_id"`
	PostID    int       `json:"post_id" db:"post_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
