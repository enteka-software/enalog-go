package enalog

type Event struct {
	Project     string            `json:"project"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Push        bool              `json:"push"`
	Icon        string            `json:"icon"`
	Tags        []string          `json:"tags"`
	Meta        map[string]string `json:"meta"`
	Channels    map[string]string `json:"channels"`
}
