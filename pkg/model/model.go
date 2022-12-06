package model

type User struct {
	ID       string `bson,json:"id"`
	Name     string `bson,json:"name"`
	Age      int32  `bson,json:"age"`
	Password string `bson,json:"password"`
	Token    string `bson,json:"Token"`
}

type Config struct {
	CurrentDB string `env:"CURRENT_DB" envDefault:"postgres"`
	JwtKey    []byte `env:"JWT-KEY" `
}
