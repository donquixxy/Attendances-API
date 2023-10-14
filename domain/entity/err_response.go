package entity

type ErrResponse struct {
	Msg   string `json:"message"`
	Error string `json:"error"`
}
