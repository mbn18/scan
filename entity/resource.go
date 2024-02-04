package entity

import "time"

type Resource struct {
	Urn         string         `json:"urn"`
	Kind        string         `json:"type"`
	Name        string         `json:"name,omitempty"`
	Data        map[string]any `json:"data,omitempty"`
	GeneratedAt time.Time      `json:"generated_at,omitempty"`
}
