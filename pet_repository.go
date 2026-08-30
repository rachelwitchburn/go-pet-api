package main

var pets []Pet

var petCreate []PetCreate

var nextID = 1

func savePet(pet Pet) {
	pets = append(pets, pet)
}

func findAllPets() []Pet {
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

func removePetById(id int) bool {
	for i, pet := range pets {
		if pet.ID == id {
			pets = append(pets[:i], pets[i+1:]...)
			return true
		}
	}
	return false
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