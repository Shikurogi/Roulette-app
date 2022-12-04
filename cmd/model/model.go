package model

type User struct {
	ID       string `bson,json:"id"`
	Name     string `bson,json:"name"`
	Age      int32  `bson,json:"age"`
	Password string `bson,json:"password"`
	Token    string `bson,json:"refreshToken"`
}

type Config struct {
	CurrentDB     string `env:"CURRENT_DB" envDefault:"postgres"`
	PostgresDBURL string `env:"POSTGRES_DB_URL"`
	MongoDBURL    string `env:"MONGO_DB_URL"`
	JwtKey        []byte `env:"JWT-KEY" `
}
