package handler

import (
	"slices"

	"github.com/brendenehlers/oapi-gen-test/generated"
)

// Contains the concrete implementations for interacting with the data source


func (s *PetstoreServerImpl) findPets(_ generated.FindPetsParams) ([]generated.Pet, error)  {
	return s.pets, nil
}

func (s *PetstoreServerImpl) addPet(newPet generated.NewPet) error {
	pet := petFromNewPet(newPet)

	s.pets = append(s.pets, pet)
	return nil
}

func (s *PetstoreServerImpl) deletePet(id int64) error {
	s.pets = slices.DeleteFunc(s.pets, func(pet generated.Pet) bool {
		return pet.Id == id
	})

	return nil
}

func (s *PetstoreServerImpl) findPetById(id int64) (*generated.Pet, error) {
	index := slices.IndexFunc(s.pets, func(pet generated.Pet) bool {
		return pet.Id == id
	})

	return &s.pets[index], nil
}