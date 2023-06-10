package wsres

type Response[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data *T     `json:"data,omitempty"`
}

type empty struct{}

func Ok[T any](data *T) Response[T] {
	return Response[T]{
		0,
		"",
		data,
	}
}

func Err(code int, msg string) Response[empty] {
	return Response[empty]{
		code,
		msg,
		nil,
	}
}
