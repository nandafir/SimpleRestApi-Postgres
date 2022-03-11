package product

import (
	"anaconda/database"

	"anaconda/utils"
)

type repository struct {
	db *database.Postgres
}

var (
	insertProduct = `
	insert into products (
		product_name, 
		category, 
		barcode, 
		price, 
		stock, 
		create_date, 
		created_by) 
	values (
		:product_name, 
		:category, 
		:barcode, 
		:price, 
		:stock, 
		now(), 
		:created_by)
		`
	selectProduct = `
	select 
		id, 
		product_name , 
		category , 
		barcode , 
		price , 
		stock , 
		create_date , 
		update_date , 
		updated_by  
	from products p	
	`
)

func (r repository) SubmitProduct(param ProductModel) error {
	tx, err := r.db.GetActiveDB().Beginx()
	if err != nil {
		utils.ErrorLog("SubmitSurvey InsertProduct", err)
		return err
	}

	_, err = tx.NamedExec(insertProduct, param)
	if err != nil {
		utils.ErrorLog("SurveyRepository UpsertSurvey NamedExec Survey", err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r repository) GetProducts() (*ProductModels, error) {
	dest := &ProductModels{}
	err := r.db.GetActiveDB().Select(dest, selectProduct+" ORDER BY product_name ASC")
	if err != nil {
		utils.ErrorLog("ProductRepository GetProdcut Select", err)
		return nil, err
	}
	return dest, nil
}
