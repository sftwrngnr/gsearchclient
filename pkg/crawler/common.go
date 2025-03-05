package crawler

type AllowedDomains struct {
	Domains []string `json:"domains"`
}

type Filterfunc func(string) error
