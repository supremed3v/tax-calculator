package iomanager

type IOManger interface {
	ReadLines() ([]string, error)
	WriteResult(data interface{}) error
}
