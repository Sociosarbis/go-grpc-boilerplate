package wsreq

type CmdExecuteReq struct {
	Id     string `json:"id" validate:"required,min=1"`
	Script string `json:"script" validate:"required,min=1"`
	Wd     string `json:"wd" validate:"required,min=1"`
}
