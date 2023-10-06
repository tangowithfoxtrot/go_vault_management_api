package schemas

import (
	"time"
)

type StatusResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Object   string `json:"object"`
		Template struct {
			ServerURL string    `json:"serverUrl"`
			LastSync  time.Time `json:"lastSync"`
			UserEmail string    `json:"userEmail"`
			UserID    string    `json:"userId"`
			Status    string    `json:"status"`
		} `json:"template"`
	} `json:"data"`
}
