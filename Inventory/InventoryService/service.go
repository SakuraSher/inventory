package inventory

type InventoryService struct {
	repo *InventoryRepository
}

func newInventoryService(repo *InventoryRepository) *InventoryService {
	return &InventoryService{repo: repo}
}

func (s *InventoryService) GetItem(id uint) (*Item, error) {
	return s.repo.FindByID(id)
}

func (s *InventoryService) CreateItem(item *Item) error {
	return s.repo.Save(item)
}

func (s *InventoryService) DeleteItem(id uint) error {
	return s.repo.Delete(id)
}

func (s *InventoryService) ListItems() ([]Item, error) {
	return s.repo.List()
}
func (s *InventoryService) UpdateItemQuantity(id uint, delta int) error {
	return s.repo.UpdateQuantity(id, delta)
}
