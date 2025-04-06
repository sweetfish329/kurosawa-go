package errors

// Handler はエラー処理のための共通インターフェースを定義します。
type Handler interface {
	Handle(err error) error
	Recover(panicked interface{}) error
}

// ErrorCode はエラーの種類を表します。
type ErrorCode int

const (
	ErrCodeInternal ErrorCode = iota
	ErrCodeInvalidInput
	ErrCodeFFmpeg
	ErrCodeIO
)

// KurosawaError は詳細なエラー情報を提供します。
type KurosawaError struct {
	Code    ErrorCode
	Message string
	Err     error
}
