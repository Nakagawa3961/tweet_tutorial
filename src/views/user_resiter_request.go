package views

type UserSignupRequest struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binging:"required"`
}
