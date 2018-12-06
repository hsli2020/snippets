func Perform() {
    for {
        SomeFunction()
        time.Sleep(time.Second)
    }
}

go Perform()

/**
 * type Context interface {
 *     Done() <-chan struct{}
 *     Err() error
 *     Deadline() (deadline time.Time, ok bool)
 *     Value(key interface{}) interface{}
 * }
 * 
 *   - ctx.Done() return cancelation channel, which is used to check if context is canceled.
 *   - ctx.Err() return cancelation reason (DeadlineExceeded or Canceled).
 *   - ctx.Deadline() return deadline, if set.
 *   - ctx.Value(key) return value for key.
 * 
 * // ctx, cancel := context.WithDeadline(parentContext, time) // or
 * // ctx, cancel := context.WithTimeout(parentContext, duration)
 */

ctx := context.Background()
ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
go Perform(ctx)

func Perform(ctx context.Context) {
    for {
        SomeFunction()

        select {
        case <-ctx.Done():
            // ctx is canceled
            return
        default:
            // ctx is not canceled, continue immediately
        }
    }
}

func Perform(ctx context.Context) error {
    for {
        SomeFunction()

        select {
        case <-ctx.Done():
            return ctx.Err()
        case <-time.After(time.Second):
            // wait for 1 second
        }
    }
    return nil
}

ctx := context.WithValue(parentContext, key, value)
value := ctx.Value(key)
