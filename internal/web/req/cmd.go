package req

type CmdCallDto struct {
	Script string `json:"script" validate:"required,min=1"`
}
