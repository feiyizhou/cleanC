package error

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var (
	ErrNotFound       = &Error{Code: 404, Message: "Resource not found"}
	ErrInternalServer = &Error{Code: 500, Message: "Internal server error"}
)
