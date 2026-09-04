package main

import "testing"

func TestValidadePetCreate_Valid(t *testing.T) {
	pet := PetCreate {
		Tipo: "cachorro",
		Raca: "Labrador",
		Idade: 3,
		Vacinado: true,
	}

	err := validatePetCreate(pet)

	if err != nil {
		t.Errorf("esperava nil, mas recebeu erro: %v",err)
	}
}

func TestValidadePetCreate_NegativeAge(t *testing.T) {
	pet := PetCreate{
		Tipo: "cachorro",
		Raca: "labrador",
		Idade: -2,
		Vacinado: true,
	}
	err := validatePetCreate(pet)
	if err == nil {
		t.Errorf("esperava erro para idade negativa, mas recebeu nil")
	}
}

func TestValidadePetUpdate_Valid(t *testing.T) {
	idade := 10
	vacinado := true
	pet := PetUpdate {
		Idade: &idade,
		Vacinado: &vacinado,
	}
	err := validatePetUpdate(pet)
	
	if err != nil {
		t.Errorf("esperava nil, mas recebeu erro: %v",err)
	}
}