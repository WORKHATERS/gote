package bot

import (
	"bytes"
	"context"
	"gote/pkg/types"
	"log"
	"net/http"
)

type TGResponse[T any] struct {
	Ok          bool                      `json:"ok"`
	Result      T                         `json:"result"`
	Description string                    `json:"description,omitempty"`
	ErrorCode   int                       `json:"error_code,omitempty"`
	Parameters  *types.ResponseParameters `json:"parameters,omitempty"`
}

func requestWithContext(ctx context.Context, url string, data []byte) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Telegram API вернул статус: %s", resp.Status)
	}

	return resp, nil
}
