package routes

import "net/http"

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func ResisterApplicationRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", RootHandler)
}
