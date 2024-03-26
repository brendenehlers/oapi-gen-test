package handler

import (
	"context"
	"slices"

	"github.com/brendenehlers/oapi-gen-test/generated"
)

// Contains the concrete implementations for interacting with the data source

type PetstoreServerImpl struct {
	ctx context.Context
	cancelFunc context.CancelFunc
	pets []generated.Pet
}

func New(ctx context.Context) *PetstoreServerImpl {
	context, cancelFunc := context.WithCancel(ctx)

	return &PetstoreServerImpl{
		ctx: context,
		cancelFunc: cancelFunc,
		pets: make([]generated.Pet, 0),
	}
}

func (s *PetstoreServerImpl) findPets(_ generated.FindPetsParams) ([]generated.Pet, error)  {
	return s.pets, nil
}

func (s *PetstoreServerImpl) addPet(newPet generated.NewPet) ( *generated.Pet, error ) {
	pet := petFromNewPet(newPet)

	s.pets = append(s.pets, pet)
	return &pet, nil
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