package req

type CmdCallDto struct {
	Script string `json:"script" validate:"required,min=1"`
	Wd     string `json:"wd" validate:"required,min=1"`
}

type CmdListFolderDto struct {
	Prefix string `json:"prefix" validate:"required,min=1"`
}

type CmdItem struct {
	Type  string `json:"type" validate:"required,oneof=exec path"`
	Value string `json:"value" validate:"required,min=1"`
}

type CmdAddDto struct {
	Items []CmdItem `json:"items" validate:"required,min=1"`
}

type CmdUpdateDto struct {
	ID    uint32    `json:"id" validate:"required,min=1"`
	Items []CmdItem `json:"items" validate:"required,min=1"`
}

type CmdDeleteDto struct {
	ID uint32 `json:"id" validate:"required,min=1"`
}

type CmdListDto struct {
	Pager
}
