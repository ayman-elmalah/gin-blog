package auth

type RegisterRequest struct {
	Name     string `form:"name" binding:"required,min=3,max=191"`
	Email    string `form:"email" binding:"required,email,min=3,max=191"`
	Password string `form:"password" binding:"required,min=8"`
}
