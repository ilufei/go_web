package action 

type output struct {
	Code int
	Message string
	Data interface{} 
}

var result = &output{
	Code : 200,
	Message : "sucess",
}