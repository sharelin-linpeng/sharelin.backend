package httppojo

import (
	"encoding/json"
	"net/http"
)

const (
	SUCCESS = 0
	ERROR   = 1
)

type ServerResponse struct {
	Status  int8        `json:"status"`
	Message string      `json:"message"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func NewServerResponse() *ServerResponse {
	return &ServerResponse{}
}

func (receiver *ServerResponse) CreateSuccess(message string) {
	receiver.Message = message
	receiver.Status = SUCCESS
}

func (receiver *ServerResponse) CreateSuccessData(message string, data interface{}) {
	receiver.Message = message
	receiver.Status = SUCCESS
	receiver.Data = data
}

func (receiver *ServerResponse) CreateError(message string) {
	receiver.Message = message
	receiver.Status = ERROR
}

type Builder struct {
}

func NewBuilder() *Builder {
	return &Builder{}
}

var ServerResponseBuilder = NewBuilder()

func (*Builder) CreateSuccess(message string) *ServerResponse {
	return &ServerResponse{Message: message, Status: SUCCESS}
}

func (*Builder) CreateSuccessData(message string, data interface{}) *ServerResponse {
	return &ServerResponse{Message: message, Status: SUCCESS, Data: data}
}

func (*Builder) CreateError(message string) *ServerResponse {
	return &ServerResponse{Message: message, Status: ERROR}
}

func WriteResponse(w *http.ResponseWriter, response *ServerResponse) {
	res, _ := json.Marshal(response)
	(*w).Write(res)
}
