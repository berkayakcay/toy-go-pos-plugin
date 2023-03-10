package mid

import (
	"context"
	"fmt"
	"github.com/berkayakcay/toy-pos-plugin/business/web/metrics"
	"github.com/berkayakcay/toy-pos-plugin/foundation/web"
	"net/http"
	"runtime/debug"
)

func Panics() web.Middleware {
	m := func(handler web.Handler) web.Handler {
		f := func(ctx context.Context, w http.ResponseWriter, r *http.Request) (err error) {
			// Defer a function to recover from a panic and set the err return
			// variable after the fact.
			defer func() {
				if rec := recover(); rec != nil {

					// Stack trace will be provided.
					trace := debug.Stack()
					err = fmt.Errorf("PANIC [%v] TRACE[%s]", rec, string(trace))

					// Updates the metrics stored in the context.
					metrics.AddPanics(ctx)
				}
			}()

			// Call the next handler and set its return value in the err variable.
			return handler(ctx, w, r)
		}
		return f
	}

	return m
}
