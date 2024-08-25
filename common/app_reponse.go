package common

//dùng generic cho kiểu dữ liệu nó an toàn
type BaseResponse[T any] struct {
	Status  string `json:"status"`
	Success bool   `json:"success"`
	Data    T      `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}
