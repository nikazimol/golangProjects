package main

import (
	"context"
	"dz-3/accounts/models"
	"dz-3/proto"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"sync"
)

type server struct {
	proto.AccountsServer
	accounts map[string]*models.Account
	guard    *sync.RWMutex
}

func (s *server) New() *server {
	return &server{
		accounts: make(map[string]*models.Account),
		guard:    &sync.RWMutex{},
	}
}

func (s *server) CreateAccount(ctx context.Context, req *proto.CreateAccountRequest) (*proto.CreateAccountResponse, error) {
	if len(req.GetName()) == 0 {
		return nil, errors.New("empty name")
	}

	s.guard.Lock()

	if _, ok := s.accounts[req.GetName()]; ok {
		s.guard.Unlock()

		return nil, errors.New("account already exists")
	}

	s.accounts[req.GetName()] = &models.Account{
		Name:   req.GetName(),
		Amount: int(req.GetAmount()),
	}

	s.guard.Unlock()

	response := &proto.CreateAccountResponse{Response: "ok"}
	return response, nil
}

func (s *server) GetAccount(ctx context.Context, req *proto.GetAccountRequest) (*proto.GetAccountResponse, error) {
	name := req.GetName()

	s.guard.RLock()

	account, ok := s.accounts[name]

	s.guard.RUnlock()

	if !ok {
		return nil, errors.New("account not found")
	}

	response := &proto.GetAccountResponse{
		Name:   account.Name,
		Amount: int32(account.Amount),
	}
	return response, nil
}

func (s *server) DeleteAccount(ctx context.Context, req *proto.DeleteAccountRequest) (*proto.DeleteAccountResponse, error) {
	if len(req.GetName()) == 0 {
		return nil, errors.New("empty name")
	}

	s.guard.Lock()

	if _, ok := s.accounts[req.GetName()]; !ok {
		s.guard.Unlock()

		return nil, errors.New("account not found")
	}

	delete(s.accounts, req.GetName())

	s.guard.Unlock()

	response := &proto.DeleteAccountResponse{Response: "ok"}

	return response, nil
}

func (s *server) PatchAccount(ctx context.Context, req *proto.PatchAmountRequest) (*proto.PatchAmountResponse, error) {

	if len(req.GetName()) == 0 {
		return nil, errors.New("empty name")
	}

	s.guard.Lock()

	if _, ok := s.accounts[req.GetName()]; !ok {
		s.guard.Unlock()

		return nil, errors.New("account not found")
	}

	s.accounts[req.GetName()].Amount = int(req.GetAmount())

	s.guard.Unlock()

	response := &proto.PatchAmountResponse{Response: "ok"}

	return response, nil
}

func (s *server) PatchName(ctx context.Context, req *proto.PatchNameRequest) (*proto.PatchNameResponse, error) {

	if len(req.GetOldName()) == 0 || len(req.GetNewName()) == 0 {
		return nil, errors.New("empty name")
	}

	s.guard.Lock()

	if _, ok := s.accounts[req.GetOldName()]; !ok {
		s.guard.Unlock()

		return nil, errors.New("account not found")
	}

	if _, ok := s.accounts[req.GetNewName()]; ok {
		s.guard.Unlock()

		return nil, errors.New("account already exists")
	}

	s.accounts[req.GetNewName()] = &models.Account{
		Name:   req.GetNewName(),
		Amount: s.accounts[req.GetOldName()].Amount,
	}

	delete(s.accounts, req.GetOldName())

	s.guard.Unlock()

	response := &proto.PatchNameResponse{Response: "account name changed"}

	return response, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 4567))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	newServer := &server{
		accounts: make(map[string]*models.Account),
		guard:    &sync.RWMutex{},
	}
	proto.RegisterAccountsServer(s, newServer)
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
