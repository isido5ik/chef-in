package dtos

type SignUpInput struct {
	Username string `json:"username" db:"username" binding:"required"`
	Email    string `json:"email" db:"email" binding:"required"`
	Password string `json:"password" db:"password_hash" binding:"required"`
}

type SignInInput struct {
	Username string `json:"username" db:"username" binding:"required"`
	Password string `json:"password" db:"password" binding:"required"`
}

type CreateRecipeInput struct {
	Content string `json:"content"`
}

type NewComment struct {
	CommentText string `json:"comment_text"`
}
type UpdateRecipeInput struct {
	Content string `json:"content" db:"content"`
}
type UpdateCommentInput struct {
	CommentText string `json:"comment_text"`
}
