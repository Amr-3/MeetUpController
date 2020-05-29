package main

import (
	"./DBConnections"
	"net/http"
)

func main() {
	http.HandleFunc("/insertintodb", DBConnections.DB_Insert_Student)
	http.HandleFunc("/removefromdb", DBConnections.DB_Delete_Student)
	http.ListenAndServe(":8080", nil)
}