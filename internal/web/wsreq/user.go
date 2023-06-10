package wsreq

type UserLoginReq struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
