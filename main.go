package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "car_project/cmd/handlers"
    "car_project/pkg/db"
)

// Middleware to set CORS headers
func setCORSHeaders(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        if r.Method == "OPTIONS" {
            return
        }
        next.ServeHTTP(w, r)
    })
}

func main() {
    // Initialize the database
    db.InitDB()

    // Create a new router
    r := mux.NewRouter()
    api := r.PathPrefix("/api").Subrouter()
    api.Use(handlers.Authenticate)
    
    // Apply CORS middleware
    r.Use(setCORSHeaders)

    // Define routes
    api.HandleFunc("/cars", handlers.CreateCar).Methods("POST")
    api.HandleFunc("/cars", handlers.GetAllCars).Methods("GET")
    api.HandleFunc("/cars/{id}", handlers.GetCar).Methods("GET")
    api.HandleFunc("/cars/{id}", handlers.UpdateCar).Methods("PUT")
    api.HandleFunc("/cars/{id}", handlers.DeleteCar).Methods("DELETE")

    api.HandleFunc("/carhistory", handlers.CreateCarHistory).Methods("POST") 
    api.HandleFunc("/carhistory", handlers.GetAllCarHistory).Methods("GET") 
    api.HandleFunc("/carhistory/{id}", handlers.GetCarHistoryByID).Methods("GET") 
    api.HandleFunc("/carhistory/{id}", handlers.UpdateCarHistory).Methods("PUT") 
    api.HandleFunc("/carhistory/{id}", handlers.DeleteCarHistory).Methods("DELETE") 

    api.HandleFunc("/ratings", handlers.CreateRating).Methods("POST")
    api.HandleFunc("/cars/{id}/ratings", handlers.GetRating).Methods("GET")
    api.HandleFunc("/ratings", handlers.UpdateRating).Methods("PUT")
    api.HandleFunc("/ratings", handlers.DeleteRating).Methods("DELETE")

    r.HandleFunc("/user/register", handlers.RegisterUser).Methods("POST")
    r.HandleFunc("/user/login", handlers.LoginUser).Methods("POST")
    r.HandleFunc("/user/activate", handlers.Activate).Methods("POST")

    // Start the server
    log.Fatal(http.ListenAndServe("0.0.0.0:8080", r))
}
