package repository

import (
	"fmt"
	"os"

	"cashier-machine/model" // Import model yang sesuai dengan struktur model di aplikasi Anda

	"github.com/sirupsen/logrus"
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

	// Log the connection string (don't expose password in production)
	logrus.Info("Connection String:", connString)

	// Open MySQL connection
	mysqlConn, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	Mysql = MysqlDB{
		DB: mysqlConn,
	}

	// Perform the migrations for the tables (if needed)
	err = autoMigrate(mysqlConn)
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return mysqlConn, nil
}

// autoMigrate ensures that the table schema in the database matches the model definitions
func autoMigrate(db *gorm.DB) error {
	// Explicitly define the table names for each model
	return db.AutoMigrate(
		&model.Admin{},       // Ensure admin table exists
		&model.Barang{},      // Ensure barang table exists
		&model.Invoice{},     // Ensure invoice table exists
		&model.Member{},      // Ensure member table exists
		&model.InvoiceItem{}, // Ensure invoice_item table exists
	)
}
