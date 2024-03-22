package common


type CustomError struct {
	Message string
}

type Response struct {
	Data   interface{} 
	Code   int
	Error *CustomError
}