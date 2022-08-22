package banco

import (
	"api/src/config"
	"database/sql"

	_ "github.com/lib/pq" // Diver de conexao com Postgres
)

// Conectar abre conexao com o bd
func Conectar() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.StringConexaoBanco)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
