package entity

type Product struct {
	ProductId     uint64
	ProductName   string
	Price         uint
	InventoryInfo []InventoryInfo
}

type InventoryInfo struct {
	InventoryId uint
	Amount      int
	// color, size
	Specification string
}

func (p *Product) CheckProductInventory() bool {
	return true
}

func (p *Product) UpdateProductInventory() bool {
	return true
}
