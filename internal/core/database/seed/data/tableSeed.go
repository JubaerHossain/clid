package dataSeed

import (
	"fmt"
	"strconv"
	"time"

	"github.com/JubaerHossain/golang-ddd/internal/core/entity"
	tableEntity "github.com/JubaerHossain/golang-ddd/internal/table/domain/entity"
	"gorm.io/gorm"
)

func TableSeed(db *gorm.DB, num int) error {
	
	if err := db.Exec("DELETE FROM tables").Error; err != nil {
		return err
	}

	for i := 0; i < num; i++ {
		var table tableEntity.Table
		table.Name = "Table " + strconv.Itoa(i+1)
		table.Status = entity.Active
		table.CreatedAt = time.Now()
		table.UpdatedAt = time.Now()

		if err := db.Create(&table).Error; err != nil {
			return err
		}
	}

	fmt.Println("Tables seeded")

	return nil
}
