package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// APIURL representa a URL para comunicação com a API
	APIURL = ""
	// Porta onde a aplicação web esta rodando
	Porta = 0
	// HasKey é utilizado para autenticar o cookie
	HasKey []byte
	// BlockKey é utilizada para criptografar os dados do cookie
	BlockKey []byte
	// GeneratedKey é utilizada para definir se ira criar as chaves hashkey e blockkey
	GeneratedKey bool
)

// Carregar inicializa as variáveis de ambiente
func Carregar() {
	var erro error
	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("APP_PORTA"))
	if erro != nil {
		log.Fatal(erro)
	}

	APIURL = os.Getenv("API_URL")
	HasKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
	GeneratedKey, _ = strconv.ParseBool(os.Getenv("GENERATED_KEY"))
}
