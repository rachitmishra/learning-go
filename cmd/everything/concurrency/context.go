package concurrency

/**
- Only top-level functions like main() should create base contexts. HTTP libraries provide request-specific contexts. Mid-level functions can create child contexts for sharing data or control.
- Contexts are passed downward in function calls. If you don't have a context, use context.TODO(). Don't return contexts from functions.
- When a function receives a context as a parameter, it should use that context only while the function is executing. Once the function returns, the context should no longer be used.
- When passed to a function or method, context.Context is always the first parameter.

type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key any) any
}
*/
