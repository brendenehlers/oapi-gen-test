package data

import (
	"slices"

	"github.com/brendenehlers/oapi-gen-test/generated"
)

// Contains the concrete implementations for interacting with the data source

type ArrayDataStore struct {
	pets []generated.Pet
}

func NewArrayDataStore() *ArrayDataStore {
	return &ArrayDataStore{
		pets: make([]generated.Pet, 0),
	} 
}

func (ds *ArrayDataStore) FindPets(_ generated.FindPetsParams) ([]generated.Pet, error) {
	return ds.pets, nil
}

func (ds *ArrayDataStore) AddPet(newPet generated.NewPet) (*generated.Pet, error) {
	pet := petFromNewPet(newPet)

	ds.pets = append(ds.pets, pet)
	return &pet, nil
}

func (ds *ArrayDataStore) DeletePet(id int64) error {
	ds.pets = slices.DeleteFunc(ds.pets, func(pet generated.Pet) bool {
		return pet.Id == id
	})

	return nil
}

func (ds *ArrayDataStore) FindPetById(id int64) (*generated.Pet, error) {
	index := slices.IndexFunc(ds.pets, func(pet generated.Pet) bool {
		return pet.Id == id
	})

	return &ds.pets[index], nil
}

func petFromNewPet(newPet generated.NewPet) generated.Pet {
	return generated.Pet{
		Id: 1,
		Name: newPet.Name,
		Tag: newPet.Tag,	
	}
}
