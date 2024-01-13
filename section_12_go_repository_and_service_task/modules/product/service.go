package product

type (
	ProductService interface {
		CreateProduct(product *Product) (*Product, error)
		GetAllProduct() ([]Product, error)
		GetProductByID(id int) (*Product, error)
		UpdateProduct(product *Product) (*Product, error)
		DeleteProduct(product *Product, id int) (*Product, error)
	}

	productService struct {
		ProductRepository ProductRepository
	}
)

func NewProductService(pr ProductRepository) ProductService {
	return &productService{
		ProductRepository: pr,
	}
}

func (ps productService) CreateProduct(product *Product) (*Product, error) {
	return ps.ProductRepository.Create(product)
}

func (ps productService) GetAllProduct() ([]Product, error) {
	return ps.ProductRepository.GetAll()
}

func (ps productService) GetProductByID(id int) (*Product, error) {
	return ps.ProductRepository.GetByID(id)
}

func (ps productService) UpdateProduct(product *Product) (*Product, error) {
	return ps.ProductRepository.Update(product)
}

func (ps productService) DeleteProduct(product *Product, id int) (*Product, error) {
	return ps.ProductRepository.Delete(product, id)
}
