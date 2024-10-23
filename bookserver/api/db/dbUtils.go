package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"xpJain.co/bookserver/models"
)

type DBServer struct {
	// DBServer is a struct that contains the database connection
	// and the database name
	DB *gorm.DB
}


var DBServerInstance *DBServer

// NewDBServer is a function that creates a new DBServer struct

func InitializeDB() {
	// Connect to the database
	ConnectionStr := "host=db user=bookuser password=test2020 dbname=bookdb port=5432 sslmode=disable TimeZone=Asia/Calcutta"
	db , err := gorm.Open(postgres.Open(ConnectionStr), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	DBServerInstance = &DBServer{
		DB: db,
	}

	fmt.Println("Database connected successfully!")

	// Create the tables
	DBServerInstance.CreateTables()
}

// GetDBServerInstance is a function that returns the DBServerInstance
func GetDBServerInstance() *DBServer {
	return DBServerInstance
}

// GetDB is a function that returns the database connection
func (d *DBServer) GetDB() *gorm.DB {
	return d.DB
}

// Create Tables is a function that creates the tables in the database
func (d *DBServer) CreateTables() {
	d.DB.AutoMigrate(&models.Book{})
	d.DB.AutoMigrate(&models.User{})
}

