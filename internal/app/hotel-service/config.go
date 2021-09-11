package hotel_service

import "github.com/BurntSushi/toml"

type config struct {
	Server Server
	Storage Storage
}

type Server struct {
	Port int
}

type Storage struct {
	Url string
	MaxPoolConn int
}

func newConfig() *config {
	return &config {
		Server{
			3000,
		},
		Storage {
			"postgresql://postgres:postgres@localhost:5432/postgres",
			30,
		},
	}
}

func (c *config) loadFromToml(tomlData string) (e error) {
	_, e = toml.DecodeFile("configs/" + tomlData, c)
	return
}