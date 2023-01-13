package mid

import (
	"context"
	"github.com/berkayakcay/toy-pos-plugin/foundation/web"
	"net/http"

	"go.uber.org/zap"
)

func Logger(log *zap.SugaredLogger) web.Middleware {

	m := func(h web.Handler) web.Handler {

		f := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

			log.Infow("request started", "method", r.Method, "path", r.URL.Path,
				"remoteaddr", r.RemoteAddr)

			err := h(ctx, w, r)

			log.Infow("request completed", "method", r.Method, "path", r.URL.Path,
				"remoteaddr", r.RemoteAddr)

			return err
		}

		return f
	}

	return m
}
