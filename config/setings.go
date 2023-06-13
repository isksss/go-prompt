package config

type Config struct {
	// Prompt
	Prompt Prompt `json:"prompt"`
}

type Prompt struct {
	// Prompt
	FgColor string `json:"fgcolor"`
	BgColor string `json:"bgcolor"`
}
