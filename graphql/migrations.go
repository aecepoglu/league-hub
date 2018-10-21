package graphql

import (
	"log"
)

func migrate() error {
	db.AutoMigrate(
		&User{},
		&Sport{},
	)

	var admin User
	if db.Where("email = ?", "admin@mail.com").First(&admin).RecordNotFound() {
		log.Println("Creating admin user with pass " + conf.AdminPass)
		admin.Email = "admin@mail.com"
		admin.Phone = "5554441122"
		admin.Password = encryptPass(conf.AdminPass)

		err := db.Create(&admin).Error

		if err != nil {
			return err
		}
	}

	var numSports int
	err := db.Model(&Sport{}).Count(&numSports).Error
	if err != nil {
		return err
	}

	if numSports < 1 {
		log.Println("Creating default sports")
		for _, name := range []string{"Squash", "Tennis"} {
			err := db.Create(&Sport{
				Name:      name,
				TeamSize:  1,
				CreatedBy: admin,
			}).Error

			if err != nil {
				return err
			}
		}
	}

	return nil
}
