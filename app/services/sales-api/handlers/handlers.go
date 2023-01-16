// Package handlers manages the different versions of the API.
package handlers

import (
	"expvar"
	"github.com/berkayakcay/toy-pos-plugin/app/services/sales-api/handlers/probegrp"
	"github.com/berkayakcay/toy-pos-plugin/business/web/v1/mid"
	"github.com/berkayakcay/toy-pos-plugin/foundation/web"
	"net/http"
	"net/http/pprof"
	"os"

	"go.uber.org/zap"
)

// APIMuxConfig contains all the mandatory systems required by handlers.
type APIMuxConfig struct {
	Shutdown chan os.Signal
	Log      *zap.SugaredLogger
	Build    string
}

// APIMux constructs a http.Handler with all application routes defined.
func APIMux(cfg APIMuxConfig) *web.App {
	app := web.NewApp(cfg.Shutdown, mid.Logger(cfg.Log), mid.Errors(cfg.Log))

	probegrp := probegrp.Handlers{
		Log:   cfg.Log,
		Build: cfg.Build,
	}
	app.Handle(http.MethodGet, "/liveness", probegrp.Liveness)
	app.Handle(http.MethodGet, "/readiness", probegrp.Readiness)

	app.Handle(http.MethodGet, "/test400", probegrp.TestError400)
	app.Handle(http.MethodGet, "/test500", probegrp.TestError500)

	return app
}

// DebugStandardLibraryMux registers all the debug routes from the standard library
// into a new mux bypassing the use of the DefaultServerMux. Using the
// DefaultServerMux would be a security risk since a dependency could inject a
// handler into our service without us knowing it.
func DebugStandardLibraryMux() *http.ServeMux {
	mux := http.NewServeMux()

	// Register all the standard library debug endpoints.
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	mux.Handle("/debug/vars", expvar.Handler())

	return mux
}
