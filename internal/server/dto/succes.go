package dto

type Success struct {
	Ok bool `json:"ok"`
}

func NewSuccess(ok bool) Success {
	return Success{Ok: ok}
}
