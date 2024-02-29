package dataSeed

import (
	"fmt"
	"time"

	categoryEntity "github.com/JubaerHossain/golang-ddd/internal/category/domain/entity"
	"github.com/JubaerHossain/golang-ddd/internal/core/entity"
	"github.com/bxcodec/faker/v3"
	"gorm.io/gorm"
)

var categories = []categoryEntity.Category{
	{Name: "Starter", Status: entity.Active},
	{Name: "Kabab Section", Status: entity.Active},
	{Name: "Curry Section", Status: entity.Active},
	{Name: "Rice & Biriyani", Status: entity.Active},
	{Name: "Special Item", Status: entity.Active},
	{Name: "Thali & Set Meal", Status: entity.Active},
	{Name: "Juice Bar Section", Status: entity.Active},
	{Name: "Soft Drinks", Status: entity.Active},
	{Name: "Floor Rent", Status: entity.Active},
}

func CategorySeed(db *gorm.DB, num int) error {

	if err := db.Exec("DELETE FROM categories").Error; err != nil {
		return err
	}

	for _, category := range categories {
		category.CreatedAt = time.Now()
		category.UpdatedAt = time.Now()
		if err := db.Create(&category).Error; err != nil {
			return err
		}
	}

	for i := 0; i < num; i++ {
		var category categoryEntity.Category
		category.Name = faker.FirstName() + " " + faker.TimeString()
		category.Status = entity.Active
		category.CreatedAt = time.Now()
		category.UpdatedAt = time.Now()

		if err := db.Create(&category).Error; err != nil {
			return err
		}
	}

	fmt.Println("Category seeding completed")

	return nil
}
