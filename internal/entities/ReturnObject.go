package entities

// Untuk enkapsulasi output JSON. Strukturnya message, httpcode dan payload data
type ReturnObject struct {
	Status  uint16 `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
	Errors  any    `json:"errors,omitempty"`
	Meta    any    `json:"meta,omitempty"`
}
