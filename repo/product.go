package repo

type Product struct {
	ID          int     `json:"id"` // = it is called tag
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	ImageURL    string  `json:"imageUrl"`
}

type ProductRepo interface {
	Get(productID int) (*Product, error)
	Create(product Product) (*Product, error)
	Update(product Product) (*Product, error)
	Delete(productID int) error
	List() ([]*Product, error)
}

type productRepo struct {
	productList []*Product
}

// Constructor or Constructor function. It creates an object of struct.
func NewProductRepo() ProductRepo {
	repo := &productRepo{}
	generateProduct(repo)
	return repo
}

func (r *productRepo) Get(productID int) (*Product, error) {
	for _, value := range r.productList {
		if productID == value.ID {
			return value, nil
		}
	}
	return nil, nil
}

func (r *productRepo) Create(product Product) (*Product, error) {
	product.ID = len(r.productList) + 1
	r.productList = append(r.productList, &product)
	return &product, nil
}

func (r *productRepo) Update(product Product) (*Product, error) {
	for intdex, value := range r.productList {
		if value.ID == product.ID {
			r.productList[intdex] = &product
		}
	}
	return nil, nil
}

func (r *productRepo) Delete(productID int) error {
	var tempList []*Product

	for _, value := range r.productList {
		if value.ID != productID {
			tempList = append(tempList, value)
		}
	}

	r.productList = tempList
	return nil
}

func (r *productRepo) List() ([]*Product, error) {
	return r.productList, nil
}

func generateProduct(r *productRepo) {
	// create struct instance/object
	products := []*Product{
		{
			ID:          1,
			Title:       "Orange",
			Description: "Orange is very jucy. I don't like it.",
			Price:       120,
			ImageURL:    "https://www.croq-kilos.com/sites/default/files/styles/1920px/public/2025-03/26505-3228-tout-savoir-sur-lorange-un-tresor-de-vitamines-et-bien-plus-encore.jpg",
		},
		{
			ID:          2,
			Title:       "Apple",
			Description: "Apple is red. I feel boring eating apple.",
			Price:       210,
			ImageURL:    "https://cdn.britannica.com/60/5760-050-FCD7CDA2/Apples-Malus-domestica.jpg",
		},
		{
			ID:          3,
			Title:       "Banana",
			Description: "Banana is yellow. I like it's smell.",
			Price:       40,
			ImageURL:    "https://images.contentstack.io/v3/assets/bltcedd8dbd5891265b/blt66b8f1f45d70de9e/6665ec0febbbaa7d99a9f533/4440854-Banana-hero.jpg?q=70&width=3840&auto=webp",
		},
		{
			ID:          4,
			Title:       "Mango",
			Description: "Mango is so yummy. I love mango.",
			Price:       145,
			ImageURL:    "https://blog2.pittmandavis.com/wp-content/uploads/2023/07/MangoFinal.jpg",
		},
		{
			ID:          5,
			Title:       "Grapes",
			Description: "Grapes is testy. I like it.",
			Price:       370,
			ImageURL:    "https://img.ksl.com/slc/2513/251331/25133147.jpg?filter=kslv2/responsive_story_lg",
		},
	}

	// append instance/object
	r.productList = append(r.productList, products...)
}
