package database

import (
	"fmt"
	"golink/shared/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() error {
	var err error
	db, err = gorm.Open(sqlite.Open("payment_links.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	// Auto-migrate the schema
	err = db.AutoMigrate(&models.PaymentLinkModel{})
	if err != nil {
		return fmt.Errorf("failed to migrate database schema: %w", err)
	}

	return nil
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func SavePaymentLink(link models.PaymentLinkModel) error {
	// Connect to the SQLite database (creates file if it doesn't exist)
	db = GetDB()

	// Save the object to the database
	result := db.Create(&link)
	if result.Error != nil {
		return fmt.Errorf("failed to save payment link: %w", result.Error)
	}

	return nil
}

func GetPaymentLinkById(id string) (*models.PaymentLinkModel, error) {
	// Connect to the SQLite database (creates file if it doesn't exist)
	db = GetDB()

	var link models.PaymentLinkModel

	// Save the object to the database
	result := db.First(&link, "id = ?", id)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to find payment link with id %s: %w", id, result.Error)
	}

	return &link, nil
}

func UpdatePaymentLink(id string, updates map[string]interface{}) error {
	// Connect to the database
	db := GetDB()

	// Attempt to update the record
	result := db.Model(&models.PaymentLinkModel{}).Where("id = ?", id).Updates(updates)
	if result.Error != nil {
		return fmt.Errorf("failed to update payment link with id %s: %w", id, result.Error)
	}

	// Check if the record was actually updated
	if result.RowsAffected == 0 {
		return fmt.Errorf("no payment link found with id %s", id)
	}

	return nil
}

func GetAllPaymentLinks() ([]models.PaymentLinkModel, error) {
	// Connect to the database
	db := GetDB()

	// Slice to store the result
	var links []models.PaymentLinkModel

	// Query the database to retrieve all records
	result := db.Find(&links)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to retrieve payment links: %w", result.Error)
	}

	return links, nil
}
