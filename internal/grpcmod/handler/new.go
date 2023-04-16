package handler

type Handler struct {
	User User
	Cmd  Cmd
}

func New(user User, cmd Cmd) Handler {
	return Handler{
		user,
		cmd,
	}
}
