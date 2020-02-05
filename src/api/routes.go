package api

import (
	"encoding/json"
	"net/http"
)

type getBalanceBody struct {
	bank     string `json:"bank"`
	username string `json:"username"`
	password string `json:"password"`
}

func (a *App) getBalance(w http.ResponseWriter, r *http.Request) {
	var p getBalanceBody

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	respondOK(w, http.StatusOK)
}
