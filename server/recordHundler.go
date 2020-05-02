package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleRecord() func(w http.ResponseWriter, r *http.Request, _ gin.HandlerFunc) {
	return func(w http.ResponseWriter, r *http.Request, _ gin.HandlerFunc) {

	}
}
