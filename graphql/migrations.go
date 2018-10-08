package graphql

import (
	"log"
)

func migrate() {
	db.AutoMigrate(&User{})
	if db.Where("email = ?", "admin").First(&User{}).RecordNotFound() {
		log.Println("Creating admin user")
		db.Create(&User{
			Email: "admin@mail.com",
			Password: encryptPass(conf.AdminPass),
			Phone: "5554441122",
		})
	}

}
