package config

var AllowedOrigins = []string{
	"http://localhost:3000",
	"http://localhost:5173",
}

func GetAllowedOrigins() []string {
	return AllowedOrigins
}