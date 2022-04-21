package articles

type StoreRequest struct {
	Title   string `form:"title" binding:"required,min=3,max=191"`
	Content string `form:"content" binding:"required,min=3,max=60000"`
}
