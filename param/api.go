package param

type UserRegister struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	MinerId  string `json:"minerId" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Email    string `json:"email"`
}
type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type UserAuthorize struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type UserUpdate struct {
	Username string `json:"username" binding:"required"`
	MinerId  string `json:"minerId" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Email    string `json:"email"`
}
