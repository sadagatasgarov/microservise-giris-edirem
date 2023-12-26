package main

import "github.com/sadagatasgarov/toll-calc/types"

type MemoryStore struct{}

func NewMemoryStore() *MemoryStore{
	return &MemoryStore{}
}

func (s *MemoryStore) Insert(d types.Distance) error{
	return nil
}