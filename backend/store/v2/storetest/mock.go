package storetest

import (
	"github.com/sensu/sensu-go/backend/store"
	storev2 "github.com/sensu/sensu-go/backend/store/v2"
	"github.com/sensu/sensu-go/backend/store/v2/wrap"
	"github.com/stretchr/testify/mock"
)

var _ storev2.Interface = new(Store)

type Store struct {
	mock.Mock
}

func (s *Store) CreateOrUpdate(req storev2.ResourceRequest, w *storev2.Wrapper) error {
	args := s.Called(req, w)
	return args.Error(0)
}

func (s *Store) UpdateIfExists(req storev2.ResourceRequest, w *storev2.Wrapper) error {
	args := s.Called(req, w)
	return args.Error(0)
}

func (s *Store) CreateIfNotExists(req storev2.ResourceRequest, w *storev2.Wrapper) error {
	args := s.Called(req, w)
	return args.Error(0)
}

func (s *Store) Get(req storev2.ResourceRequest) (*storev2.Wrapper, error) {
	args := s.Called(req)
	return args.Get(0).(*storev2.Wrapper), args.Error(1)
}

func (s *Store) Delete(req storev2.ResourceRequest) error {
	args := s.Called(req)
	return args.Error(0)
}

func (s *Store) List(req storev2.ResourceRequest, pred *store.SelectionPredicate) (wrap.List, error) {
	args := s.Called(req, pred)
	return args.Get(0).(wrap.List), args.Error(1)
}

func (s *Store) Exists(req storev2.ResourceRequest) (bool, error) {
	args := s.Called(req)
	return args.Get(0).(bool), args.Error(1)
}
