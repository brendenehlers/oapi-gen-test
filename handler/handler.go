package handler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/brendenehlers/oapi-gen-test/generated"
)

// Contains the functions for implementing the `generated.ServerInterface`
// and some helper functions to abstract the `generated` package from
// `main.go`

type DataStore interface {
	FindPets(params generated.FindPetsParams) ([]generated.Pet, error)
	FindPetById(id int64) (*generated.Pet, error)
	AddPet(newPet generated.NewPet) (*generated.Pet, error)
	DeletePet(id int64) error
}

type PetstoreServerImpl struct {
	ctx context.Context
	cancelFunc context.CancelFunc
	ds DataStore
}

func New(ctx context.Context) *PetstoreServerImpl {
	context, cancelFunc := context.WithCancel(ctx)

	return &PetstoreServerImpl{
		ctx: context,
		cancelFunc: cancelFunc,
		ds: &DataStoreImpl{
			pets: make([]generated.Pet, 0),
		},
	}
}

func (s *PetstoreServerImpl) Handler() http.Handler {
	return generated.Handler(s)
}

func (s *PetstoreServerImpl) FindPets(w http.ResponseWriter, r *http.Request, params generated.FindPetsParams) {
	pets, err := s.ds.FindPets(params)
	panicIfError(err)

	panicIfError(sendResponse(w, pets))
}

func (s *PetstoreServerImpl) FindPetById(w http.ResponseWriter, r *http.Request, id int64) {
	pet, err := s.ds.FindPetById(id)
	panicIfError(err)

	panicIfError(sendResponse(w, pet))
}

func (s *PetstoreServerImpl) AddPet(w http.ResponseWriter, r *http.Request) {
	newPet, err := readBody[generated.NewPet](r.Body)
	panicIfError(err)

	pet, err := s.ds.AddPet(newPet)
	panicIfError(err)

	panicIfError(sendResponseAsJSON(w, pet))
}

func (s *PetstoreServerImpl) DeletePet(w http.ResponseWriter, r *http.Request, id int64) {
	err := s.ds.DeletePet(id)
	panicIfError(err)
}

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func sendResponse(w http.ResponseWriter, data any) error {
	if data != nil {
		return sendResponseWithData(w, data)
	} else {
		return sendMissingDataResponse(w)
	}
}

func sendResponseWithData(w http.ResponseWriter, data any) error {
	return sendResponseAsJSON(w, data)
}

func sendMissingDataResponse(w http.ResponseWriter) error {
	w.WriteHeader(http.StatusNotFound)
	return nil
}

func sendResponseAsJSON(w http.ResponseWriter, body any) error {
	return json.NewEncoder(w).Encode(body)
}

func readBody[T any](r io.Reader) (T, error) {
	var val T
	err := json.NewDecoder(r).Decode(&val)
	return val, err
}
