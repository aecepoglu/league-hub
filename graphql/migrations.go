package graphql

import (
	"log"
)

func migrate() {
	db.AutoMigrate(&User{})
	if db.Where("email = ?", "admin@mail.com").First(&User{}).RecordNotFound() {
		log.Println("Creating admin user with pass " + conf.AdminPass)
		db.Create(&User{
			Email: "admin@mail.com",
			Password: encryptPass(conf.AdminPass),
			Phone: "5554441122",
		})
	}

}
