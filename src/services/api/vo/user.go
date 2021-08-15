package vo

//UserRegisterVO mean view object for user register request
type UserRegisterVO struct {
	Username string `form:"username" bind:"required"`
	Password string `form:"password" bind:"required"`
}
