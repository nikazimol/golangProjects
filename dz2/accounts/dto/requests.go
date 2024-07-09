package dto

type CreateAccountRequest struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type PatchAccountRequest struct {
	OldName string `json:"old_name"`
	NewName string `json:"new_name"`
}

type ChangeAccountRequest struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type DeleteAccountRequest struct {
	Name string `json:"name"`
}
