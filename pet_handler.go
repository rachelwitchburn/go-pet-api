package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// CRUD

// GET
func getPet (w http.ResponseWriter, r *http.Request) {
	result := findAllPets()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)

	fmt.Println("Listando os pets")
}

// GET Pet by ID
func getPetByIdD (w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Id inválido", http.StatusBadRequest)
		return
	}

	pet, found := findPetByID(id)
	if !found {
		http.Error(w, "Pet não encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pet)
}

// POST
func createPet (w http.ResponseWriter, r *http.Request) {
	var newPet PetCreate

	// tratar o erro e desserializar
	err := json.NewDecoder(r.Body).Decode(&newPet)
	if err != nil {
    	fmt.Println("Erro ao decodificar JSON:", err)
    	http.Error(w, "JSON inválido", http.StatusBadRequest)
    	return
	}

	pet, err := createPetService(newPet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pet)

	fmt.Println("Criando o pet")
}

/*
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
	*/

// DELETE
func deletePet (w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id")) // id e err tem o mesmo valor da variavel pq o erro será a verificação do id invalido, e o id será o proprio id
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	deleted := removePetById(id)
	if !deleted {
		http.Error(w, "Pet não encontrado", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
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
	err = validatePetUpdate(petUpdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pet, err := updatePet(id, petUpdate)
	if err != nil {
		http.Error(w, "Pet não encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pet)

}