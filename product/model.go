package product

import (
	"time"

	"github.com/google/uuid"
	// "xapiens.id/shared/utils"
)

type ProductModel struct {
	ID          uuid.UUID  `db:"id"`
	ProductName string     `db:"product_name"`
	Category    string     `db:"category"`
	Barcode     string     `db:"barcode"`
	Price       float64    `db:"price"`
	Stock       float64    `db:"stock"`
	CreateDate  time.Time  `db:"create_date"`
	CreatedBy   string     `db:"created_by"`
	UpdateDate  *time.Time `db:"update_date"`
	UpdatedBy   string     `db:"updated_by"`
}

type ProductModels []ProductModel
