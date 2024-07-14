package main

import (
	"context"
	"dz-3/proto"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type Command struct {
	Port    int
	Host    string
	Cmd     string
	Name    string
	NewName string
	Amount  int
}

func (c *Command) create(conn *grpc.ClientConn) error {
	cl := proto.NewAccountsClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := cl.CreateAccount(ctx, &proto.CreateAccountRequest{Name: c.Name, Amount: int32(c.Amount)})
	if err != nil {
		panic(err)
	}

	return nil
}

func (c *Command) get(conn *grpc.ClientConn) error {
	cl := proto.NewAccountsClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := cl.GetAccount(ctx, &proto.GetAccountRequest{Name: c.Name})
	if err != nil {
		panic(err)
	}

	fmt.Printf("response account name: %s and amount: %d", res.GetName(), res.GetAmount())

	return nil
}

func (c *Command) delete(conn *grpc.ClientConn) error {
	cl := proto.NewAccountsClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := cl.DeleteAccount(ctx, &proto.DeleteAccountRequest{Name: c.Name})
	if err != nil {
		panic(err)
	}

	return nil
}

func (c *Command) patchAmount(conn *grpc.ClientConn) error {
	cl := proto.NewAccountsClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := cl.PatchAmount(ctx, &proto.PatchAmountRequest{Name: c.Name, Amount: int32(c.Amount)})
	if err != nil {
		panic(err)
	}

	return nil
}

func (c *Command) patchName(conn *grpc.ClientConn) error {
	cl := proto.NewAccountsClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := cl.PatchName(ctx, &proto.PatchNameRequest{OldName: c.Name, NewName: c.NewName})
	if err != nil {
		panic(err)
	}

	return nil
}

func main() {
	portVal := flag.Int("port", 8080, "server port")
	hostVal := flag.String("host", "0.0.0.0", "server host")
	cmdVal := flag.String("cmd", "", "command to execute")
	nameVal := flag.String("name", "", "name of account")
	newNameVal := flag.String("new_name", "", "new name of account")
	amountVal := flag.Int("amount", 0, "amount of account")

	flag.Parse()

	cmd := Command{
		Port:    *portVal,
		Host:    *hostVal,
		Cmd:     *cmdVal,
		Name:    *nameVal,
		NewName: *newNameVal,
		Amount:  *amountVal,
	}

	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", cmd.Host, cmd.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = conn.Close()
	}()

	if err := do(cmd); err != nil {
		panic(err)
	}
}

func do(cmd Command) error {
	conn, err := grpc.NewClient("0.0.0.0:4567", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = conn.Close()
	}()
	switch cmd.Cmd {
	case "create":
		if err := cmd.create(conn); err != nil {
			return fmt.Errorf("error in creating account: %w", err)
		}

		return nil
	case "get":
		if err := cmd.get(conn); err != nil {
			return fmt.Errorf("error in getting account: %w", err)
		}

		return nil
	case "delete":
		if err := cmd.delete(conn); err != nil {
			return fmt.Errorf("error in deleting account: %w", err)
		}

		return nil
	case "patch-name":
		if err := cmd.patchName(conn); err != nil {
			return fmt.Errorf("error in patching name: %w", err)
		}

		return nil
	case "patch-amount":
		if err := cmd.patchAmount(conn); err != nil {
			return fmt.Errorf("error in patching amount: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unknown command %s", cmd.Cmd)
	}
}
