package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/evanlimanto/quickapi/src/database"
	"github.com/google/uuid"
)

type GetBalanceBody struct {
	Login database.Login `json:"login"`
}

var banks = []string{"bca", "bni", "bri", "cimb", "mandiri"}

func contains(list []string, s string) bool {
	for _, item := range list {
		if item == s {
			return true
		}
	}
	return false
}

func (app *App) GetBalance(w http.ResponseWriter, r *http.Request) {
	var body GetBalanceBody

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body.Login); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if !contains(banks, body.Login.Bank) {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid bank %s", body.Login.Bank))
		return
	}

	// Check if Login is present in the database.
	if err := body.Login.GetLoginByBankAndUsername(app.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			// Create Login if it's not present yet.
			body.Login.ID = uuid.New().String()
			if err = body.Login.CreateLogin(app.DB); err != nil {
				respondWithError(w, http.StatusInternalServerError, err.Error())
				return
			}
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondOK(w, http.StatusOK)
}
