package req

type Pager struct {
	Page uint32 `query:"page" validate:"required,min=1"`
	Size uint32 `query:"size" validate:"required,min=1"`
}
