package log

import (
    "github.com/plexmediamanager/service/errors"
    "github.com/sirupsen/logrus"
)

var logger *logrus.Entry

type Fields logrus.Fields

// Entry point
func init() {
    logger = logrus.NewEntry(logrus.New())
}

// Get new instance of fields structure
func NewFields() *Fields {
    return &Fields{}
}

// Add new field with given value
func (fields *Fields) Add (key string, value interface{}) *Fields {
    local := *fields
    local[key] = value
    *fields = local
    return fields
}

// Merge two fields structures
func (fields *Fields) merge(newFields *Fields) {
    if newFields != nil {
        for key, value := range *newFields {
            fields.Add(key, value)
        }
    }
}

// Get instance of logrus with predefined fields
func (fields *Fields) logrus() logrus.Fields {
    return logrus.Fields(*fields)
}

// Set logger to our instance of logrus
func SetLogger(logrusInstance *logrus.Entry) {
    if logrusInstance != nil {
        logger = logrusInstance
    }
}

// Get global instance of logger
func GetLogger() *logrus.Entry {
    return logger
}

func Print(arguments ...interface{}) {
    arguments, fields := prepareArguments(arguments)
    logger.WithFields(fields.logrus()).Print(arguments...)
}

func Printf(format string, arguments ...interface{}) {
    arguments, fields := prepareArguments(arguments)
    logger.WithFields(fields.logrus()).Printf(format, arguments...)
}

func Info(arguments ...interface{}) {
    arguments, fields := prepareArguments(arguments)
    logger.WithFields(fields.logrus()).Info(arguments...)
}

func Infof(format string, arguments ...interface{}) {
    arguments, fields := prepareArguments(arguments)
    logger.WithFields(fields.logrus()).Infof(format, arguments...)
}

func Warn(arguments ...interface{}) {
    arguments, fields := prepareArguments(arguments)
    logger.WithFields(fields.logrus()).Warn(arguments...)
}

func Warnf(format string, arguments ...interface{}) {
    arguments, fields := prepareArguments(arguments)
    logger.WithFields(fields.logrus()).Warnf(format, arguments...)
}

func Error(arguments ...interface{}) {
    arguments, fields := prepareArguments(arguments)
    logger.WithFields(fields.logrus()).Error(arguments...)
}

func Errorf(format string, arguments ...interface{}) {
    arguments, fields := prepareArguments(arguments)
    logger.WithFields(fields.logrus()).Errorf(format, arguments...)
}

func Debug(arguments ...interface{}) {
    arguments, fields := prepareArguments(arguments)
    logger.WithFields(fields.logrus()).Debug(arguments...)
}

func Debugf(format string, arguments ...interface{}) {
    arguments, fields := prepareArguments(arguments)
    logger.WithFields(fields.logrus()).Debugf(format, arguments...)
}

func Panic(arguments ...interface{}) {
    arguments, fields := prepareArguments(arguments)
    logger.WithFields(fields.logrus()).Panic(arguments...)
}

func Panicf(format string, arguments ...interface{}) {
    arguments, fields := prepareArguments(arguments)
    logger.WithFields(fields.logrus()).Panicf(format, arguments...)
}

func TraceError(e error) {
    if err, ok := e.(*errors.Error); ok {
        if err.Err != nil {
            TraceError(err.Err)
        }
        Error(err.Message, NewFields().Add("error-code", err.Code))
    } else {
        Error(err)
    }
}

func TracePanicError(e error) {
    if err, ok := e.(*errors.Error); ok {
        if err.Err != nil {
            TraceError(err.Err)
        }
        Panic(err.Message, NewFields().Add("error-code", err.Code))
    } else {
        Panic(err)
    }
}

// Convert arguments array to the desired format
func prepareArguments(arguments []interface{}) ([]interface{}, *Fields) {
    var newArguments []interface{}
    fields := &Fields{}
    for _, value := range arguments {
        if field, ok := value.(*Fields); ok {
            fields.merge(field)
        } else {
            newArguments = append(newArguments, value)
        }
    }
    return newArguments, fields
}