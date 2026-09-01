package main

import (
	"errors"
)

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

func validatePetUpdate(update PetUpdate) error {
	if update.Idade == nil && update.Vacinado == nil {
		return errors.New("Nenhum campo para atualizar")
	}
	if update.Idade != nil && *update.Idade < 0 {
		return errors.New("idade não pode ser negativa")
	}
	return nil
}

func updatePet(id int, update PetUpdate) (Pet, error) {
	err := validatePetUpdate(update)
	if err != nil {
		return Pet{}, err
	}
	pet, found := updatePetByID(id, update)
	if !found {
		return Pet{}, errors.New("pet não encontrado")
	}
	return pet, nil
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

func createPetService (input PetCreate) (Pet, error) {
	err := validatePetCreate(input)
	if err != nil {
		return Pet{}, err
	}
	pet := buildPet(input)

	savePet(pet)
	return pet, nil
}