package store

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"

	"github.com/henrriusdev/nails/config"

	_ "modernc.org/sqlite"
)

func NewConnection(cfg config.EnvVar) (Queryable, error) {
	// Definir la ruta de la base de datos SQLite (archivo)
	dsn := fmt.Sprintf("%s.db", cfg.DBName)

	// Abrir conexión con sqlx
	connection, err := sqlx.Open("sqlite", dsn)
	if err != nil {
		log.Fatalf("Error al conectar con SQLite: %v", err)
		return nil, err
	}

	// Probar la conexión
	if err := connection.Ping(); err != nil {
		log.Fatalf("Error al hacer ping a la base de datos: %v", err)
		return nil, err
	}

	fmt.Println("Conexión exitosa a SQLite")
	return connection, nil
}
