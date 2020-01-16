package main

import (
	"log"
	"net/http"
	"os"
)

var ADDRESS = getEnv("ADDRESS", ":20020")

// 获取环境变量
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// 请求检查
func handle(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	log.Printf("Authorization: %s", auth)

	if auth == "" {
		unauthorizedResponse(w, r)
		return
	}

	if auth != "123" {
		forbiddenResponse(w, r)
		return
	}

	authorizedResponse(w, r)
}

func main() {
	http.HandleFunc("/", handle)

	log.Printf("start server at: %s", ADDRESS)
	log.Fatal(http.ListenAndServe(ADDRESS, nil))
}
