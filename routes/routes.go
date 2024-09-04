package routes

import (
    "github.com/gorilla/mux"
    "myapp/controllers"
)

func RegisterItemRoutes(router *mux.Router) {
    router.HandleFunc("/item", controllers.CreateItem).Methods("GET")
}
