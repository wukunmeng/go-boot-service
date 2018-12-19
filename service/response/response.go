/**
 * Create with IntelliJ IDEA
 * Project name : go-boot-service
 * Package name : 
 * Author : Wukunmeng
 * User : wukm
 * Date : 18-12-19
 * Time : 下午3:29
 * ---------------------------------
 * 
 */
package response

import (
    "net/http"
    "github.com/labstack/echo"
)

var (
    OK  = &Response{Message: "SUCCESS", Code: "SUCCESS", Success:true}
    Err = &Response{Message: "Failed", Code: "Failed", Success:false}
)

type Response struct {
    Success     bool    `json:"success"`
    Code        string  `json:"code"`
    Message     string  `json:"message"`
    Data interface{}    `json:"data,omitempty"`
}

func (r *Response) Clone() *Response {
    c := *r
    return &c
}

func (r *Response) SetCode(code string) *Response {
    r.Code = code
    return r
}

func (r *Response) SetMessage(message string) *Response {
    r.Message = message
    return r
}

func (r *Response) SetData(data ...interface{}) *Response {
    if data == nil {
        return r
    }

    var d interface{} = data
    if len(data) == 1 {
        d = data[0]
    }
    r.Data = d
    return r
}

func Success(c echo.Context, data ...interface{}) error {
    return c.JSON(http.StatusOK, OK.Clone().SetData(data...))
}

func Error(c echo.Context, code string, message string, data ...interface{}) error {
    return c.JSON(http.StatusOK,
        Err.Clone().
            SetCode(code).
            SetMessage(message).
            SetData(data),
    )
}
