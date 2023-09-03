package wsres

type Response[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data *T     `json:"data,omitempty"`
}

type SeqFrame[T any] struct {
	SeqId    string      `json:"SeqId"`
	Index    int         `json:"index"`
	Response Response[T] `json:"response"`
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

func Seq[T any](id string, index int, res Response[T]) SeqFrame[T] {
	return SeqFrame[T]{
		id,
		index,
		res,
	}
}
