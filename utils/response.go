package utils

//type MetaData struct {
//	TotalCount  int
//	TotalPage   int
//	CurrentPage int
//	PerPage     int
//}

type JsonResponse struct {
	Error int 			`json:"error"`
	Msg   string 		`json:"msg"`
	Data  interface{} 	`json:"data"`
	//	Meta MetaData
}

func wrapResponse(error int, msg string, data interface{}) *JsonResponse {
	return &JsonResponse{error, msg, data}
}

func Success(data interface{}) *JsonResponse {
	return wrapResponse(0, "ok", data)
}

func Error(code int, msg string) *JsonResponse {
	return wrapResponse(code, msg, nil)
}

