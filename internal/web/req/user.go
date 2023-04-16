package req

type UserDetailDto struct {
	ID uint32 `params:"id" validate:"required,min=1"`
}
