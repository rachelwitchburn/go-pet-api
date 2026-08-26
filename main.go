package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"errors"
)

type Pet struct {
	ID int `json:"id"`
	Tipo string `json:"tipo"` // cachorro, gato, coelho, hamster
	Raca string `json:"raca"`
	Idade int `json:"idade"`
	Vacinado bool `json:"vacinado"`
}

type PetUpdate struct {
	Idade *int `json:"idade"`
	Vacinado *bool `json:"vacinado"`
}

type PetCreate struct {
	Tipo string `json:"tipo"`
	Raca string `json:"raca"`
	Idade int `json:"idade"`
	Vacinado bool `json:"vacinado"`
}

var pets []Pet

var petCreate []PetCreate

var nextID = 1

func validatePetCreate (pet PetCreate) error {
	if pet.Tipo == "" {
		return errors.New("Tipo do pet é obrigatório")
	}

	if pet.Raca == "" {
		return errors.New("Raça do pet é obrigatória")
	}

	if pet.Idade < 0 {
		return errors.New("Idade não pode ser negativa")
	}
	return nil
}

func buildPet(input PetCreate) Pet {
	pet := Pet {
		ID:       nextID,
        Tipo:     input.Tipo,
        Raca:     input.Raca,
        Idade:    input.Idade,
        Vacinado: input.Vacinado, 
	}

	nextID++
	return pet
}

func savePet (pet Pet) {
	pets = append(pets, pet)
}

func findAllPets () []Pet {
	return pets
}

func findPetByID(id int) (Pet, bool) {
	for _, pet := range pets {
		if pet.ID == id {
			return pet, true
		}
	}
	return Pet{}, false
}

func removePetById (id int) bool {
	for i, pet := range pets {
		if pet.ID == id {
			pets = append(pets[:i], pets[i + 1:]...)
			return true	
		}
	}
	return false
}

func validatePetUpdate(update PetUpdate) error {
	if update.Idade == nil && update.Vacinado == nil {
		return errors.New("Nenhum campo para atualizar")
	}
	if update.Idade != nil && *update.Idade < 0 {
		return errors.New("idade não pode ser negativa")
	}
	return nil
}

func updatePetByID(id int, update PetUpdate) (Pet, bool) {
	for i, pet := range pets {
		if pet.ID == id {

			if update.Idade != nil {
				pets[i].Idade = *update.Idade
			}

			if update.Vacinado != nil {
				pets[i].Vacinado = *update.Vacinado
			}
			return pets[i], true
		}
	}
	return Pet{}, false
}

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

	err = validatePetCreate(newPet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pet := buildPet(newPet)
	savePet(pet)

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

	pet, found := updatePetByID(id, petUpdate)
	if !found {
		http.Error(w, "Pet não encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pet)

}


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