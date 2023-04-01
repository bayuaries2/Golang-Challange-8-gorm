package config

import (
	"Challange-7/model"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Gorm struct {
	// db configuration
	Username string
	Password string
	Port     string
	Address  string
	Database string

	// db connection
	DB *gorm.DB
}

type GormDb struct {
	*Gorm
}

var (
	GORM *GormDb
)

func InitGorm() error {
	GORM = new(GormDb)

	GORM.Gorm = &Gorm{
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Address:  os.Getenv("POSTGRES_ADDRESS"),
		Database: os.Getenv("POSTGRES_DB"),
	}

	// connect to database
	err := GORM.Gorm.OpenConnection()
	if err != nil {
		return err
	}

	return nil
}

func (p *Gorm) OpenConnection() error {
	// init dsn
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", p.Address, p.Port, p.Username, p.Password, p.Database)

	dbConnection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "db_book_gorm.",
			SingularTable: false,
		},
	})
	if err != nil {
		return err
	}

	p.DB = dbConnection

	// migration
	err = p.DB.Debug().AutoMigrate(model.Book{})
	if err != nil {
		return err
	}

	// read migration sql files - make dummy migration data
	// queries, err := ioutil.ReadFile("./migration/migration.sql")
	// if err != nil {
	// 	return err
	// }
	// err = p.DB.Exec(string(queries)).Error
	// if err != nil {
	// 	return err
	// }

	fmt.Println("Successfully connected to database")

	return nil
}