package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/brendenehlers/oapi-gen-test/generated"
)

// Contains the functions for implementing the `generated.ServerInterface`
// and some helper functions to abstract the `generated` package from
// `main.go`


func (s *PetstoreServerImpl) Handler() http.Handler {
	return generated.Handler(s)
}

func (s *PetstoreServerImpl) FindPets(w http.ResponseWriter, r *http.Request, params generated.FindPetsParams) {
	log.Println("Testing FindPets")

	pets, err := s.findPets(params)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	if pets != nil {
		json.NewEncoder(w).Encode(pets)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func (s *PetstoreServerImpl) AddPet(w http.ResponseWriter, r *http.Request) {
	log.Println("Testing AddPet")

	var newPet generated.NewPet
	_ = json.NewDecoder(r.Body).Decode(&newPet)

	err := s.addPet(newPet)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *PetstoreServerImpl) DeletePet(w http.ResponseWriter, r *http.Request, id int64) {
	log.Printf("Testing DeletePet, id: %v\n", id)

	err := s.deletePet(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (s *PetstoreServerImpl) FindPetById(w http.ResponseWriter, r *http.Request, id int64) {
	log.Printf("Testing FindPetById, id: %v", id)

	pet, err := s.findPetById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if pet != nil {
		json.NewEncoder(w).Encode(pet)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
