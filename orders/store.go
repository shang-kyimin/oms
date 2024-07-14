package main

import "context"

type store struct {
	// Add here our MongoDB
}

func NewStore() *store {
	return &store{}
}

func (s *store) Create(context.Context) error {
	return nil
}
