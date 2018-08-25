package config

/*
Note: We are actually using TOML.

We convert TOML to JSON, then we unmarshal to struct from JSON.
TOML can unmarshal, but it's does not preserve the input of the struct, like JSON does.
*/

type Config struct {
	Http  Http   `json:"http"`
	Tls   *Tls   `json:"tls"`
	Mimes []Mime `json:"mime"`
}

type Http struct {
	Debug     bool   `json:"debug"`
	Address   string `json:"address"`
	CacheTime int    `json:"cacheTime"`
}

type Tls struct {
	Cert string `json:"cert"`
	Key  string `json:"key"`
}

type Mime struct {
	Extension string `json:"extension"`
	Type      string `json:"type"`
}
