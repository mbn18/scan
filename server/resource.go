package server

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func ResourceList(w http.ResponseWriter, r *http.Request, pl httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")

	kind, err := dbMapper.ListByKind(r.Context(), pl.ByName("type"))
	if err != nil {
		// @TODO , internal errors should not returned to consumer
		json.NewEncoder(w).Encode(Response{Error: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(Response{Data: kind})
}
