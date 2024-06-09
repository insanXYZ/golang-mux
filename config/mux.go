package config

import "net/http"

func Mux() *http.ServeMux {
	return http.NewServeMux()
}
