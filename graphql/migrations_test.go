package graphql

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setupDb(t *testing.T) {
	_db, err := gorm.Open("sqlite3", "test.db")
	assert.Nil(t, err)
	db = _db
	db.LogMode(false)
	migrate()
}

func cleanupDb(t *testing.T) {
	assert.Nil(t, db.Close())
	db = nil
}

func TestDbCreatesAdmin(t *testing.T) {
	setupDb(t)
	db.DropTableIfExists(&User{})
	assert.False(t, db.Where("email = ? AND password = ?", "admin@mail.com", encryptPass(conf.AdminPass)).First(&User{}).RecordNotFound())

	cleanupDb(t)
}

func TestDbDoesntCreatesAdminIfExists(t *testing.T) {
	setupDb(t)
	db.DropTableIfExists(&User{})
	db.AutoMigrate(&User{})
	db.Create(&User{
		Email: "admin@mail.com",
		Password: "old admin pass",
	})
	assert.False(t, db.Where("email = ? AND password = ?", "admin@mail.com", "old admin pass").First(&User{}).RecordNotFound())

	cleanupDb(t)
}
