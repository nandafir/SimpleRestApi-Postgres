package product

import (
	"anaconda/database"
	"anaconda/utils"
	// "xapiens.id/geosurvey/database"
	// "xapiens.id/shared/content"
	// "xapiens.id/shared/errs"
	// "xapiens.id/shared/utils"
)

type Service struct {
	repository *repository
	// contentClient       *content.Client
}

func NewService(
// contentClient *content.Client,
) *Service {
	return &Service{
		// contentClient:       contentClient,
		repository: &repository{
			db: database.New(),
		},
	}
}

func (s Service) SubmitProduct(params ProductRequest) error {

	//paramRequest to paramInsert
	p := ProductModel{
		ProductName: params.ProductName,
		Category:    params.Category,
		Barcode:     params.Barcode,
		Price:       params.Price,
		Stock:       params.Stock,
		CreatedBy:   "jhon doe",
	}

	err := s.repository.SubmitProduct(p)
	if err != nil {
		utils.ErrorLog("SQL Error on InsertProduct", err)
		return err
	}

	return nil
}

func (s Service) GetProducts() (ProductModels, error) {

	// res := ProductModels{}

	res, err := s.repository.GetProducts()
	if err != nil {
		utils.ErrorLog("SQL Error on GetProduct", err)
		return nil, err
	}

	return *res, nil
}
