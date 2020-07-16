package controllers

import (
	"net/http"

	"github.com/j127/golang_rest_api_jwt/utils"
)

// ProtectedEndpoint needs a JWT to access it
func (c Controller) ProtectedEndpoint() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.ResponseJSON(w, "yes")
	}

}
