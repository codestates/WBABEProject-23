package protocol

type ApiResponse[T any] struct {
	Code      int    `json:"code"`
	Data      T      `json:"data"`
	Message   string `json:"message"`
	ErrorName string `json:"error"`
}
