package main

import (
	"fmt"
	"net/http"
)

func main () {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /pets", getPet)
	mux.HandleFunc("GET /pets/{id}", getPetByIdD)
	mux.HandleFunc("POST /pets", createPet)
	//mux.HandleFunc("PUT /pets/{id}", putPet)
	mux.HandleFunc("DELETE /pets/{id}", deletePet)
	mux.HandleFunc("PATCH /pets/{id}", patchPet)

	fmt.Println("Servidor rodando em http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}