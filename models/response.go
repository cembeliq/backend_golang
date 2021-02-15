package models

// Response for templating response api
type Response struct {
	Status   int         `json:"status"`
	Messsage string      `json:"message"`
	Data     interface{} `json:"data"`
}
