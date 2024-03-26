package handler

import (
	"slices"

	"github.com/brendenehlers/oapi-gen-test/generated"
)

// Contains the concrete implementations for interacting with the data source

type DataStoreImpl struct {
	pets []generated.Pet
}

func (d *DataStoreImpl) FindPets(_ generated.FindPetsParams) ([]generated.Pet, error) {
	return d.pets, nil
}

func (d *DataStoreImpl) AddPet(newPet generated.NewPet) (*generated.Pet, error) {
	pet := petFromNewPet(newPet)

	d.pets = append(d.pets, pet)
	return &pet, nil
}

func (d *DataStoreImpl) DeletePet(id int64) error {
	d.pets = slices.DeleteFunc(d.pets, func(pet generated.Pet) bool {
		return pet.Id == id
	})

	return nil
}

func (d *DataStoreImpl) FindPetById(id int64) (*generated.Pet, error) {
	index := slices.IndexFunc(d.pets, func(pet generated.Pet) bool {
		return pet.Id == id
	})

	return &d.pets[index], nil
}
