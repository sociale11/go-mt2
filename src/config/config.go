package config

type Config struct {
	AppVersion string
	AppHost    string
	AppPort    int
	AppDebug   bool
}

var Settings = Config{
	AppVersion: "0.0.1",
	AppHost:    "localhost",
	AppPort:    8080,
	AppDebug:   true,
}
