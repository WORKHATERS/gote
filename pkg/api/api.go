package api

import (
	"gote/pkg/types"
	"net/http"
)

const URL = "https://api.telegram.org/bot"

type TGResponse[T any] struct {
	Ok          bool                      `json:"ok"`
	Result      T                         `json:"result"`
	Description string                    `json:"description,omitempty"`
	ErrorCode   int                       `json:"error_code,omitempty"`
	Parameters  *types.ResponseParameters `json:"parameters,omitempty"`
}

type API struct {
	token  string
	client *http.Client
}

func New(token string) *API {
	return &API{
		token:  token,
		client: &http.Client{},
	}
}
