package main

import (
	"bytes"
	"dz-3/accounts/dto"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
)

type Command struct {
	Port    int
	Host    string
	Cmd     string
	Name    string
	NewName string
	Amount  int
}

func (c *Command) Do() error {
	switch c.Cmd {
	case "create":
		return c.create()
	case "get":
		return c.get()
	case "delete":
		return c.delete()
	case "patch-amount":
		return c.patchAmount()
	case "patch-name":
		return c.patchName()
	default:
		return fmt.Errorf("unknown command: %s", c.Cmd)
	}
}

func (c *Command) create() error {
	request := dto.CreateAccountRequest{
		Name:   c.Name,
		Amount: c.Amount,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	resp, err := http.Post(
		fmt.Sprintf("http://%s:%d/account/create", c.Host, c.Port),
		"application/json",
		bytes.NewReader(data),
	)
	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode == http.StatusCreated {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}

func (c *Command) get() error {
	resp, err := http.Get(
		fmt.Sprintf("http://%s:%d/account?name=%s", c.Host, c.Port, c.Name),
	)
	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("read body failed: %w", err)
		}

		return fmt.Errorf("resp error %s", string(body))
	}

	var response dto.GetAccountResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return fmt.Errorf("json decode failed: %w", err)
	}

	fmt.Printf("response account name: %s and amount: %d", response.Name, response.Amount)

	return nil
}

func (c *Command) delete() error {
	req := dto.DeleteAccountRequest{
		Name: c.Name,
	}

	_, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("JSON marshal failed: %w", err)
	}
	request, err := http.NewRequest(http.MethodDelete,
		fmt.Sprintf("http://%s:%d/account/delete?name=%s", c.Host, c.Port, c.Name),
		nil)
	if err != nil {
		return fmt.Errorf("delete failed: %w", err)
	}

	request.Header.Set("Content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		return fmt.Errorf("http delete failed: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return nil
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}

func (c *Command) patchAmount() error {
	request := dto.ChangeAccountRequest{
		Name:   c.Name,
		Amount: c.Amount,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	req, err := http.NewRequest(http.MethodPatch,
		fmt.Sprintf("http://%s:%d/account/patch-account", c.Host, c.Port),
		bytes.NewReader(data))

	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}

	req.Header.Set("Content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return fmt.Errorf("http patch_account failed: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}

func (c *Command) patchName() error {
	request := dto.PatchAccountRequest{
		OldName: c.Name,
		NewName: c.NewName,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	req, err := http.NewRequest(http.MethodPatch,
		fmt.Sprintf("http://%s:%d/account/change-account", c.Host, c.Port),
		bytes.NewReader(data))

	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}

	req.Header.Set("Content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return fmt.Errorf("http change-name failed: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}

func main() {
	portVal := flag.Int("port", 8080, "server port")
	hostVal := flag.String("host", "0.0.0.0", "server host")
	cmdVal := flag.String("cmd", "", "command to execute")
	nameVal := flag.String("name", "", "name of account")
	amountVal := flag.Int("amount", 0, "amount of account")

	flag.Parse()

	cmd := Command{
		Port:   *portVal,
		Host:   *hostVal,
		Cmd:    *cmdVal,
		Name:   *nameVal,
		Amount: *amountVal,
	}

	if err := cmd.Do(); err != nil { //do(cmd)
		panic(err)
	}
}
