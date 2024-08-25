package common

import (
	"fmt"
	"net/http"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"key_err"`
}
//function new apperror for custom error reponse
func NewAppError(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr: root,
		Message: msg,
		Log: log,
		Key: key,
	}
}
func NewFullReponseError(root error, status int,msg,log,key string) *AppError{
	return &AppError{
		StatusCode: status,
		RootErr: root,
		Message: msg,
		Log: log,
		Key: key,
	}
}
//recursive function to get the root error
func (e *AppError) RootError() error{
	if err,ok:=e.RootErr.(*AppError);ok{
		return err.RootError()
	}
	return e.RootErr;
}
func (e *AppError) Error() string {
	return e.RootErr.Error()
}
//custom error

//error from the user
func ErrorInvalidRequest(err error) *AppError{
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr: err,
		Message:"Something went wrong in the server",
		Log: err.Error(),
		Key: "ErrInternal",
	}
}
//error from the system
func ErrorInternal(err error) *AppError{
	return &AppError{
		StatusCode: http.StatusBadGateway,
		RootErr: err,
		Message:"Something went wrong in the server",
		Log: err.Error(),
		Key: "ErrInternal",
	}
}
func ErrorDB(err error) *AppError{
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr: err,
		Message: "Something went wrong in the server",
		Log: err.Error(),
		Key:"ErrInternal",
	}
}
func ErrorCreateEntity(err error, entityName string) *AppError{
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr: err,
		Message: fmt.Sprintf("%s can not created",entityName) ,
	}
}
func ErrorListEntity(err error, entityName string) *AppError{
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr: err,
		Message: fmt.Sprintf("%s can not find",entityName) ,
	}
}
func ErrorUpdateEntity(err error, entityName string) *AppError{
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr: err,
		Message: fmt.Sprintf("%s can not update",entityName) ,
	}
}
func ErrorDeleteEntity(err error, entityName string) *AppError{
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr: err,
		Message: fmt.Sprintf("%s can not delete",entityName) ,
	}
}