package errors

import (
    format "fmt"
    "github.com/plexmediamanager/service/proto"
    "strconv"
)

type Service uint32
type Type uint32
type Number uint32

type Code struct {
    Service     Service
    ErrorType   Type
    ErrorNumber Number
}

var serviceIdentifier = Service(0)

func SetServiceIdentifier(service Service) {
    serviceIdentifier = service
}

func GenerateErrorCode(errorType Type, errorNumber uint32) Code {
    return Code {
        Service:     serviceIdentifier,
        ErrorType:   errorType,
        ErrorNumber: Number(errorNumber),
    }
}

func (code *Code) String() string {
    return format.Sprintf("%s%s%s", code.Service, code.ErrorType, code.ErrorNumber)
}

func (code *Code) UnmarshalText(bytesArray []byte) error {
    if number, err := strconv.ParseUint(string(bytesArray[:2]), 10, 64); err != nil {
        return err
    } else {
        code.Service = Service(number)
    }
    if number, err := strconv.ParseUint(string(bytesArray[2]), 10, 64); err != nil {
        return err
    } else {
        code.ErrorType = Type(number)
    }
    if number, err := strconv.ParseUint(string(bytesArray[3:]), 10, 64); err != nil {
        return err
    } else {
        code.ErrorNumber = Number(number)
    }
    return nil
}

func (code *Code) ToProto() *proto.Error_Code {
    return &proto.Error_Code{
        Service:              uint32(code.Service),
        Type:                 code.ErrorType.ToProto(),
        Number:               uint32(code.ErrorNumber),
    }
}

func (code *Code) FromProto(protoCode *proto.Error_Code) {
    code.Service = Service(protoCode.Service)
    code.ErrorType.FromProto(protoCode.Type)
    code.ErrorNumber = Number(protoCode.Number)
}

func (service Service) String() string {
    return format.Sprintf("%02d", service)
}

func (errorType Type) String() string {
    return format.Sprintf("%d", errorType)
}

func (errorType Type) ToProto() proto.Error_Code_Type {
    switch errorType {
    case TypeService:
        return proto.Error_Code_Service
    case TypeNetwork:
        return proto.Error_Code_Network
    case TypeDateTime:
        return proto.Error_Code_DateTime
    case TypeMicro:
        return proto.Error_Code_Micro
    case TypeLibrary:
        return proto.Error_Code_Library
    case TypeWrapper:
        return proto.Error_Code_Wrapper
    default:
        return proto.Error_Code_Undefined
    }
}

func (errorType *Type) FromProto(errorProto proto.Error_Code_Type) {
    switch errorProto {
    case proto.Error_Code_Service:
        *errorType = TypeService
    case proto.Error_Code_Network:
        *errorType = TypeNetwork
    case proto.Error_Code_DateTime:
        *errorType = TypeDateTime
    case proto.Error_Code_Micro:
        *errorType = TypeMicro
    case proto.Error_Code_Library:
        *errorType = TypeLibrary
    case proto.Error_Code_Wrapper:
        *errorType = TypeWrapper
    default:
        *errorType = TypeUndefined
    }
}

func (number Number) String() string {
    return format.Sprintf("%03d", number)
}

const (
    TypeUndefined   Type = 0
    TypeService     Type = 1
    TypeNetwork     Type = 2
    TypeDateTime    Type = 3
    TypeMicro       Type = 4
    TypeLibrary     Type = 5
    TypeWrapper     Type = 6
)