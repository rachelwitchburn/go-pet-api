package main

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