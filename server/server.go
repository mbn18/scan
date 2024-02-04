package server

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/mbn18/scan/mapper"
	"log"
	"net/http"
)

var dbMapper mapper.Mapper

type Response struct {
	Data  any    `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

func Start(address string, mapper mapper.Mapper) {
	dbMapper = mapper

	router := httprouter.New()
	router.GET("/", rootRequest)
	// @todo, depend on the logic, we can move the `type` out of the path to a url param.
	router.GET("/resource/type/:type", ResourceList)

	http.ListenAndServe(address, LoggerMiddleware(router))
}

func rootRequest(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		log.Printf("%s %s", r.Method, r.URL.Path)
	})
}
