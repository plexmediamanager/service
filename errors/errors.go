package errors

import (
    format "fmt"
    "github.com/plexmediamanager/service/proto"
    micro "github.com/micro/go-micro/errors"
    "net"
    "regexp"
)

var (
    regex = regexp.MustCompile(`^\[(\d{6})\] Error: (.*)`)
)

type Error struct {
    Code        Code
    Message     string
    Err         *Error
}

func New(code Code, message string, nativeError error) *Error {
    newError := &Error {
        Code:       code,
        Message:    message,
    }
    newError.Err = ParseError(nativeError)
    return newError
}

func (err *Error) Error() string {
    return format.Sprintf("[%s] Error: %s", err.Code.String(), err.Message)
}

func (err *Error) Unwrap() *Error {
    return err.Err
}

func (err Error) ToProto() *proto.Error {
    protoError := &proto.Error {
        Code:       err.Code.ToProto(),
        Message:    err.Message,
    }
    if err.Err != nil {
        protoError.Error = err.Err.ToProto()
    }
    return protoError
}

func (err *Error) FromProto(protoError *proto.Error) {
    err.Code.FromProto(protoError.Code)
    err.Message = protoError.Message
    if protoError.Error != nil {
        err.Err = &Error{}
        err.Err.FromProto(protoError.Error)
    }
}

func FromNativeError(err error) *Error {
    return &Error{
        Code:       GenerateErrorCode(TypeLibrary, 999),
        Message:    err.Error(),
        Err:        nil,
    }
}

func (err *Error) ParseError(nativeError error) {
    result := regex.FindStringSubmatch(nativeError.Error())
    if result != nil {
        err.Code.UnmarshalText([]byte(result[1]))
        err.Message = result[2]
    } else {
        *err = *FromNativeError(nativeError)
    }
}

func FromMicroError(err *micro.Error) *Error {
    return &Error {
        Code:       GenerateErrorCode(TypeLibrary, 999),
        Message:    err.Error(),
        Err:        nil,
    }
}

func FromNetworkError(err net.Error) *Error {
    networkError := &Error {}
    if err.Timeout() {
        networkError.Code = GenerateErrorCode(TypeNetwork, 1)
        networkError.Message = "Network connection timeout"
    } else if err.Temporary() {
        networkError.Code = GenerateErrorCode(TypeNetwork, 2)
        networkError.Message = "Network connection aborted"
    } else {
        networkError.Code = GenerateErrorCode(TypeNetwork, 999)
        networkError.Message = "Unknown network error"
    }
    return networkError
}

func ParseError(err error) *Error {
    var newError *Error
    if err != nil {
        switch errorType := err.(type) {
        case *micro.Error:
            newError = FromMicroError(errorType)
        case *Error:
            newError = errorType
        case net.Error:
            newError = FromNetworkError(errorType)
        default:
            if format.Sprintf("%T", errorType) == "client.serverError" {
                newError = &Error{}
                newError.ParseError(errorType)
            } else {
                newError = FromNativeError(errorType)
            }
        }
    }
    return newError
}

func (err Error) ToError(nativeError error) *Error {
    return New(
        err.Code,
        err.Message,
        nativeError,
    )
}

func (err Error) ToErrorWithArguments(nativeError error, args ...interface{}) *Error {
    return New(
        err.Code,
        format.Sprintf(err.Message, args...),
        nativeError,
    )
}