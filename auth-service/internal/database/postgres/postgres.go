package postgres

import (
	"fmt"
	"log"

	"github.com/vinit-jpl/restaurant-management/auth-service/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDb struct {
	cfg *config.Config
}

func NewPostgresDb(psqlCfg *config.Config) *PostgresDb {
	return &PostgresDb{cfg: psqlCfg}
}

func (p *PostgresDb) Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		p.cfg.DBHost,
		p.cfg.DBUser,
		p.cfg.DBPassword,
		p.cfg.DBName,
		p.cfg.DBPort,
		p.cfg.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	log.Println("Connecting to Postgres database...")

	if err != nil {
		log.Fatalf("Failed to connect to Postgres database: %v", err)
		return nil, err
	}

	return db, nil
}
