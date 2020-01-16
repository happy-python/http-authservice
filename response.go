package main

import (
	"log"
	"net/http"
)

// 认证通过 200
func authorizedResponse(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s 200", r.RequestURI)
	w.WriteHeader(http.StatusOK)
}

// 未认证通过 401
func unauthorizedResponse(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s 401", r.RequestURI)
	w.WriteHeader(http.StatusUnauthorized)
}

// 没有权限 403
func forbiddenResponse(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s 403", r.RequestURI)
	w.WriteHeader(http.StatusForbidden)
}
