package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Pet struct {
	ID int `json:"id"`
	Tipo Tipo `json:"tipo"`
	Idade int `json:"idade"`
	Vacinado bool `json:"vacinado"`
}

type PetUpdate struct {
	Idade *int `json:"idade"`
	Vacinado *bool `json:"vacinado"`
}

type Tipo struct {
	Tipo string `json:"tipo"`
}


var pets []Pet

// CRUD

// GET
func getPet (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pets)
	fmt.Println("Listando os pets")
	return
}

// GET Pet by ID
func getPetByIdD (w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Id inválido", http.StatusBadRequest)
		return
	}

	for _, pet := range pets {
		if pet.ID == id {
			json.NewEncoder(w).Encode(pet)
			return
		}
	}
	http.Error(w, "Pet não encontrado", http.StatusNotFound)
}

// POST
func createPet (w http.ResponseWriter, r *http.Request) {
	var newPet Pet

	// tratar o erro e desserializar
	err := json.NewDecoder(r.Body).Decode(&newPet)
	if err != nil {
		http.Error(w, "Json inválido", http.StatusBadRequest)
		return
	}
    
	// salvar em pets
	pets = append(pets, newPet)
	fmt.Println("Criando o pet")
	return
}

// PUT
func putPet (w http.ResponseWriter, r *http.Request) {
	var newPet Pet
	
	id, errTipo := strconv.Atoi(r.PathValue("id"))
	if errTipo != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	
	errJson := json.NewDecoder(r.Body).Decode(&newPet)
	if errJson != nil {
		http.Error(w, "Json inválido", http.StatusBadRequest)
		return
	}

	for i, pet := range pets {
		if pet.ID == id {
			pets[i] = newPet
			fmt.Fprintln(w, "Pet atualizado")
			return
		}
	}
}

// DELETE
func deletePet (w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id")) // id e err tem o mesmo valor da variavel pq o erro será a verificação do id invalido, e o id será o proprio id
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	for i, pet := range pets {
		if pet.ID == id {
			pets = append(pets[:i], pets[i + 1:]...)
			fmt.Fprintln(w, "Pet deletado")
			return
		}
	}
    http.Error(w, "Pet não encontrado", http.StatusNotFound)
}

// PATCH
func patchPet(w http.ResponseWriter, r *http.Request) {
	var petUpdate PetUpdate
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	
	err = json.NewDecoder(r.Body).Decode(&petUpdate)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if petUpdate.Idade == nil && petUpdate.Vacinado == nil {
		http.Error(w, "Nenhum campo para atualizar", http.StatusBadRequest)
		return
	}

	for i, pet := range pets {
		if pet.ID == id {
			if petUpdate.Idade != nil {
				pets[i].Idade = *petUpdate.Idade
			}
			if petUpdate.Vacinado != nil {
				pets[i].Vacinado = *petUpdate.Vacinado
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(pets[i])

			return
		}
	}
	http.Error(w, "Pet não encontrado", http.StatusNotFound)
}


func main () {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /pets", getPet)
	mux.HandleFunc("GET /pets/{id}", getPetByIdD)
	mux.HandleFunc("POST /pets", createPet)
	mux.HandleFunc("PUT /pets/{id}", putPet)
	mux.HandleFunc("DELETE /pets/{id}", deletePet)
	mux.HandleFunc("PATCH /pets/{id}", patchPet)

	http.ListenAndServe(":8080", nil)
}