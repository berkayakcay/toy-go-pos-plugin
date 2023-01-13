package probegrp

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

type Handlers struct {
	Log *zap.SugaredLogger
}

func (h Handlers) Liveness(w http.ResponseWriter, r *http.Request) {
	h.Log.Infow("Liveness")

	status := struct {
		Status string
	}{
		Status: "OK",
	}
	json.NewEncoder(w).Encode(status)
}

func (h Handlers) Readiness(w http.ResponseWriter, r *http.Request) {
	h.Log.Infow("Readiness")

	status := struct {
		Status string
	}{
		Status: "OK",
	}
	json.NewEncoder(w).Encode(status)
}
