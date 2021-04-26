package singleprocess

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/hashicorp/vagrant/internal/server/proto/vagrant_server"
)

// TODO: test
func (s *service) UpsertBasis(
	ctx context.Context,
	req *vagrant_server.UpsertBasisRequest,
) (*vagrant_server.UpsertBasisResponse, error) {
	result := req.Basis
	if err := s.state.BasisPut(result); err != nil {
		return nil, err
	}

	return &vagrant_server.UpsertBasisResponse{Basis: result}, nil
}

// TODO: test
func (s *service) GetBasis(
	ctx context.Context,
	req *vagrant_server.GetBasisRequest,
) (*vagrant_server.GetBasisResponse, error) {
	result, err := s.state.BasisGet(req.Basis)
	if err != nil {
		return nil, err
	}

	return &vagrant_server.GetBasisResponse{Basis: result}, nil
}

func (s *service) FindBasis(
	ctx context.Context,
	req *vagrant_server.FindBasisRequest,
) (*vagrant_server.FindBasisResponse, error) {
	result, err := s.state.BasisFind(req.Basis)
	if err != nil {
		return nil, err
	}

	return &vagrant_server.FindBasisResponse{Basis: result, Found: true}, nil
}

// TODO: test
func (s *service) ListBasis(
	ctx context.Context,
	req *empty.Empty,
) (*vagrant_server.ListBasisResponse, error) {
	result, err := s.state.BasisList()
	if err != nil {
		return nil, err
	}

	return &vagrant_server.ListBasisResponse{Basis: result}, nil
}