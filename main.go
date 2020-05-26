package main

import (
	"./DBConnections"
	"net/http"
)

func main() {
	http.HandleFunc("/insertintodb", DBConnections.DB_Insert_Student)
	http.ListenAndServe(":8080", nil)
}