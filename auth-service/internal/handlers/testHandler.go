package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

func Test(ctx http.ResponseWriter, r *http.Request) {
	resp := map[string]any{
		"Status":     "success",
		"Code":       http.StatusOK,
		"message":    "test handler working",
		"method":     r.Method,
		"path":       r.URL.Path,
		"request_id": middleware.GetReqID(r.Context()),
		"remote_ip":  r.RemoteAddr,
		"headers":    r.Header,
	}

	ctx.Header().Set("Content-Type", "application/json")
	json.NewEncoder(ctx).Encode(resp)
}
