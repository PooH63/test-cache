package api

import (
	"net/http"
)

type Result string

// String возвращает строковое представление.
func (s Result) String() string {
	return string(s)
}

const (
	STORED    Result = "stored"
	IDENTICAL Result = "identical"
	DIFFER    Result = "differ"

	EMPTY_QUERY_PARAMS_ERROR string = "empty query params"
	EMPTY_PARAM_VALUE_ERROR string = "empty param value"
)

//Response структура ответа.
type Response struct {
	Success bool        `json:"status"`
	Error   string      `json:"error,omitempty"`
	Result  string      `json:"result,omitempty"`
	Stored  interface{} `json:"stored,omitempty"`
	Value   interface{} `json:"value,omitempty"`
}

// NewFailedResponse конструктор успешного ответа.
func newResponse(status bool, err string, result string, stored interface{}, value interface{}) *Response {
	return &Response{
		Success: status,
		Error:   err,
		Result:  result,
		Stored:  stored,
		Value:   value,
	}
}

func (c *Context) NewResponseError(err error) {
	c.JSON(
		http.StatusOK,
		newResponse(false, err.Error(), "", nil, nil),
	)
}

// ResponseOK возвращает успешный ответ.

func (c *Context) NewResponseSuccess(result string, stored interface{}, value interface{}) {
	resp := &Response{}

	if result == STORED.String() {
		resp = newResponse(true, "", result, nil, nil)
	} else {
		resp = newResponse(true, "", result, stored, value)
	}

	c.JSON(http.StatusOK, resp)
}
