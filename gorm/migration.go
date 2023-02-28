package main

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "201608301400",
			Migrate: func(tx *gorm.DB) error {
				// it's a good pratice to copy the struct inside the function,
				// so side effects are prevented if the original struct changes during the time
				type Person struct {
					gorm.Model
					Name string
				}
				return tx.AutoMigrate(&Person{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("people")
			},
		},
		// add age column to persons
		{
			ID: "201608301415",
			Migrate: func(tx *gorm.DB) error {
				// when table already exists, it just adds fields as columns
				type Person struct {
					Age int
				}
				return tx.AutoMigrate(&Person{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropColumn("people", "age")
			},
		},
		{
			ID: "201608301416",
			Migrate: func(tx *gorm.DB) error {
				// when table already exists, it just adds fields as columns
				type Person struct {
					Age int
				}
				return tx.AutoMigrate(&Person{})
			},
		},
		{
			ID: "201608301417",
			Migrate: func(tx *gorm.DB) error {
				type Person struct {
					Age      int
					PersonID int
				}
				return tx.Migrator().DropTable("people")
			},
		},
		{
			ID: "201608301418",
			Migrate: func(tx *gorm.DB) error {
				type Member struct {
					Age      int
					MemberID int
				}
				return tx.Migrator().CreateTable(&Member{})
			},
		},
		// add pets table
	})

	return m.Migrate()
}
