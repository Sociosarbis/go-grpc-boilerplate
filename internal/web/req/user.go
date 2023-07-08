package req

type UserDetailDto struct {
	ID uint32 `params:"id" validate:"required,min=1"`
}

type UserMsLoginDto struct {
	Token string `params:"token" validate:"required,min=1"`
}
