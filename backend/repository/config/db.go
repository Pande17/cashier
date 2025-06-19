package repository

import (
	"fmt"
	"os"
	"projek/toko-retail/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MysqlDB contains the GORM DB instance for MySQL
type MysqlDB struct {
	DB *gorm.DB
}

var Mysql MysqlDB

// OpenDB opens a connection to the MySQL database and performs migrations
func OpenDB() (*gorm.DB, error) {
	// Fetch database credentials from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Create the connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser, dbPass, dbHost, dbPort, dbName)

	fmt.Println("Connection String:", connString)

	// Open MySQL connection
	mysqlConn, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	Mysql = MysqlDB{
		DB: mysqlConn,
	}

	// Perform migrations
	if err := autoMigrate(mysqlConn); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return mysqlConn, nil
}

// autoMigrate migrates the database schema to match the model definitions
func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.Model{},
		&model.Barang{},
		&model.Penjualan{},
		&model.Diskon{},
		&model.Histori{},
	)
}
