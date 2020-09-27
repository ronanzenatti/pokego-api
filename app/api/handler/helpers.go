package handler

type genContext interface {
	JSON(code int, i interface{}) error
	Param(name string) string
}
