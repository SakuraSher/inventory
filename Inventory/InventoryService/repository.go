// database interaction logic
package inventory

import (
	"gorm.io/gorm"
)

type InventoryRepository struct {
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) *InventoryRepository {
	return &InventoryRepository{db: db}
}

func (r *InventoryRepository) FindByID(id uint) (*Item, error) {
	var item Item
	result := r.db.First(&item, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &item, nil
}

func (r *InventoryRepository) Save(item *Item) error {
	return r.db.Save(item).Error
}

func (r *InventoryRepository) UpdateQuantity(id uint, delta int) error {
	return r.db.Model(&Item{}).Where("id = ?", id).UpdateColumn("quantity", gorm.Expr("quantity + ?", delta)).Error
}

func (r *InventoryRepository) Delete(id uint) error {
	return r.db.Delete(&Item{}, id).Error
}

func (r *InventoryRepository) List() ([]Item, error) {
	var items []Item
	result := r.db.Find(&items)
	return items, result.Error
}
