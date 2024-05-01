package dtos

type SignInInput struct {
	Username string `json:"username" db:"username" binding:"required"`
	Password string `json:"password" db:"password" binding:"required"`
}
type UpdateStoryInput struct {
	Content string `json:"content" db:"content"`
}
type UpdateCommentInput struct {
	CommentText string `json:"new_comment"`
}
