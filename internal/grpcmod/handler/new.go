package handler

type Handler struct {
	User User
}

func New(user User) Handler {
	return Handler{
		user,
	}
}
