package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "myapp/routes"
)

func main() {
    r := mux.NewRouter()
    routes.RegisterItemRoutes(r)
    http.Handle("/", r)
    http.ListenAndServe(":8080", nil)
}
