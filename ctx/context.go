package ctx

import "context"

var globalContext context.Context

func Context() context.Context {
    if globalContext == nil {
        globalContext = context.Background()
    }
    return globalContext
}

func WithValue(key, value interface{}) {
    globalContext = context.WithValue(Context(), key, value)
}

func Value(key interface{}) interface{} {
    return globalContext.Value(key)
}