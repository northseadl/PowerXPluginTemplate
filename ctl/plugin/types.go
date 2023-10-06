package plugin

type Info struct {
	Name string `yaml:"name"`
}

type BuildInfo struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Version     string `yaml:"version"`
	Backend     string `yaml:"backend"`
}

type Route struct {
	Name string `yaml:"name" json:"name"`
	Path string `yaml:"path" json:"path"`
	Meta struct {
		Icon   string `yaml:"icon" json:"icon"`
		Locale string `yaml:"locale" json:"locale"`
	} `yaml:"meta" json:"meta"`
}
