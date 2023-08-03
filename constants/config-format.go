package constants

// ConfigFormat config format
type ConfigFormat string

const (
	JSON       ConfigFormat = "json"
	YAML       ConfigFormat = "yaml"
	PROPERTIES ConfigFormat = "properties"
	TOML       ConfigFormat = "toml"
	INI        ConfigFormat = "ini"
)
