package domain

type Products struct {
	ID       int     `db:"id"`
	Name     string  `db:"name"`
	Price    float64 `db:"price"`
	Count    int     `db:"count"`
	Type     string  `db:"type"`
	Category string  `json:"name" db:"name"`
}

type Category struct {
	Name string `json:"name" db:"name"`
}

type ProductRequest struct {
	ID       int      `json:"id",omitempty"`
	Name     string   `json:"name"`
	Price    float64  `json:"price"`
	Count    int      `json:"count"`
	Type     string   `json:"type"`
	Category Category `json:"category"`
}

type ProductsRepository interface {
	FindAll() ([]ProductRequest, error)
	FindByID(id int) (ProductRequest, error)
	FindByName(name string) (ProductRequest, error)
	Create(product ProductRequest) (ProductRequest, error)
	Remove(product ProductRequest) error
	UpdateCount(product ProductRequest) (ProductRequest, error)
}

type ProductsService interface {
	FindAll() ([]ProductRequest, error)
	FindByID(id int) (ProductRequest, error)
	FindByName(name string) (ProductRequest, error)
	Create(product ProductRequest) (ProductRequest, error)
	Remove(product ProductRequest) error
	UpdateCount(product ProductRequest) (ProductRequest, error)
}
