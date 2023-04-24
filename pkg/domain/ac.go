package domain

type AC struct {
	Enabled     bool    `json:"enabled"`
	Temperature float32 `json:"temperature"`
	Humidity    float32 `json:"humidity"`
}
