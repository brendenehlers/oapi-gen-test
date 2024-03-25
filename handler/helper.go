package handler

import "github.com/brendenehlers/oapi-gen-test/generated"

func petFromNewPet(newPet generated.NewPet) generated.Pet {
	return generated.Pet{
		Id: 1,
		Name: newPet.Name,
		Tag: newPet.Tag,	
	}
}