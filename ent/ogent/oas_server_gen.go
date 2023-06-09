// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// CreatePerson implements createPerson operation.
	//
	// Creates a new Person and persists it to storage.
	//
	// POST /people
	CreatePerson(ctx context.Context, req *CreatePersonReq) (CreatePersonRes, error)
	// DeletePerson implements deletePerson operation.
	//
	// Deletes the Person with the requested ID.
	//
	// DELETE /people/{id}
	DeletePerson(ctx context.Context, params DeletePersonParams) (DeletePersonRes, error)
	// ListPerson implements listPerson operation.
	//
	// List Persons.
	//
	// GET /people
	ListPerson(ctx context.Context, params ListPersonParams) (ListPersonRes, error)
	// ReadPerson implements readPerson operation.
	//
	// Finds the Person with the requested ID and returns it.
	//
	// GET /people/{id}
	ReadPerson(ctx context.Context, params ReadPersonParams) (ReadPersonRes, error)
	// UpdatePerson implements updatePerson operation.
	//
	// Updates a Person and persists changes to storage.
	//
	// PATCH /people/{id}
	UpdatePerson(ctx context.Context, req *UpdatePersonReq, params UpdatePersonParams) (UpdatePersonRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h Handler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		baseServer: s,
	}, nil
}
