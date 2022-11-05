package db

import (
	"github.com/hugobally/mimiko_api/internal/db/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Client struct {
	*gorm.DB
}

func NewClient(verbose bool) *Client {
	level := logger.Warn
	if verbose {
		level = logger.Info
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(level),
	})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.Map{}, &models.User{}, &models.Knot{}, &models.Link{})

	return &Client{
		db,
	}
}

func (c *Client) FindUser(id uint) (*models.User, error) {
	var u models.User

	res := c.First(&u, id)
	if res.Error != nil {
		return nil, res.Error
	}

	return &u, nil
}

func (c *Client) FindMap(id uint) (*models.Map, error) {
	var m models.Map

	res := c.First(&m, id)
	if res.Error != nil {
		return nil, res.Error
	}

	return &m, nil
}
