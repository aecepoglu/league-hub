package graphql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"
	"testing"
)

func setupTestDb(t *testing.T) {
	_db, err := gorm.Open("sqlite3", "test.db")
	assert.Nil(t, err)
	db = _db
	db.LogMode(false)
}

func cleanupTestDb(t *testing.T) {
	assert.Nil(t, db.Close())
	db = nil
}

func TestMigrationsCreatesAdmin(t *testing.T) {
	setupTestDb(t)
	db.DropTableIfExists(&User{})
	assert.Nil(t, migrate())
	assert.False(t, db.Where("email = ? AND password = ?", "admin@mail.com", encryptPass(conf.AdminPass)).First(&User{}).RecordNotFound())

	cleanupTestDb(t)
}

func TestMigrationsDoesntCreatesAdminIfExists(t *testing.T) {
	setupTestDb(t)
	db.DropTableIfExists(&User{})
	db.AutoMigrate(&User{})
	db.Create(&User{
		Email:    "admin@mail.com",
		Password: "old admin pass",
	})
	assert.Nil(t, migrate())
	assert.NotEqual(t, "old admin pass", encryptPass("old admin pass"))
	assert.False(t, db.Where("email = ? AND password = ?", "admin@mail.com", "old admin pass").First(&User{}).RecordNotFound())
	assert.True(t, db.Where("email = ? AND password = ?", "admin@mail.com", encryptPass(conf.AdminPass)).First(&User{}).RecordNotFound())
	cleanupTestDb(t)
}

func TestMigrationsCreatesDefaultSports(t *testing.T) {
	setupTestDb(t)
	db.DropTableIfExists(&Sport{})
	assert.Nil(t, migrate())
	var count int
	assert.Nil(t, db.Model(&Sport{}).Count(&count).Error)
	assert.Equal(t, 2, count)
	cleanupTestDb(t)
}
