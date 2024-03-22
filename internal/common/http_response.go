package common


type CustomError struct {
	Message string
}

type Response struct {
	Data   interface{} 
	Code    uint
	Error *CustomError
}