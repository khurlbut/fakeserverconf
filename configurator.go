package fakeserverconf

type Configuration struct {
	Path   string
	Status int
	Body   string
}

func Configurator() []*Configuration {
	c := &Configuration{
		Path:   "/",
		Status: 200,
		Body:   "Welcome to fakeserver. This is the Default configuration for Root path (/)",
	}
	configs := []*Configuration{c}

	return configs
}
