package config

var AllowedOrigins = []string{
	"http://localhost:3000",
}

func GetAllowedOrigins() []string {
	return AllowedOrigins
}