package repo

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
} 


type ProductRepo interface{
	Create(p Product) (*Product, error)
	Get(productID int) (*Product, error)
	List() ([]*Product, error)
	Delete(productID int) error
	Update(p Product) (*Product, error)
}

type productRepo struct{
	productList []*Product
}

func NewProductRepo() ProductRepo{

	repo := &productRepo{}

	generateInitialProducts(repo)

	return repo
}


func (r *productRepo) Create(p Product) (*Product, error){
	p.ID = len(r.productList) + 1
	r.productList = append(r.productList, &p)
	return &p, nil
}


func (r *productRepo) Get(productID int) (*Product, error){
	for _, product := range r.productList {
		if product.ID == productID {
			return product, nil
		}
	}
	return nil, nil

}


func (r *productRepo) List() ([]*Product, error){
	return r.productList, nil
}


func (r *productRepo) Delete(productID int) error{
	var tempList []*Product

	for _, p := range r.productList {
		if p.ID != productID {
			tempList = append(tempList, p)
		}
	}
	r.productList = tempList

	return nil
}


func (r *productRepo) Update(product Product) (*Product, error){
	for idx, p := range r.productList {
		if p.ID == product.ID {
			r.productList[idx] = &product
		}
	}
	return &product, nil
}


func generateInitialProducts(r *productRepo) {
	prd1 := &Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange contains vitamin C",
		Price:       350,
		ImgUrl:      "https://www.allrecipes.com/thmb/y_uvjwXWAuD6T0RxaS19jFvZyFU=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/GettyImages-1205638014-2000-d0fbf9170f2d43eeb046f56eec65319c.jpg",
	}

	prd2 := &Product{
		ID:          2,
		Title:       "Apple",
		Description: "Apple has lot of minerals",
		Price:       200,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQD7ANGmYuBLltBJmAFrhfX5p0p9UdzJqfiYlVPIIXplE_RAXFsdTeuru13CUP_L4Nv1ew&usqp=CAU",
	}

	prd3 := &Product{
		ID:          3,
		Title:       "Banana",
		Description: "Bana has high carb",
		Price:       200,
		ImgUrl:      "https://img.freepik.com/free-vector/vector-ripe-yellow-banana-bunch-isolated-white-background_1284-45456.jpg?semt=ais_incoming&w=740&q=80",
	}

	prd4 := &Product{
		ID:          4,
		Title:       "Mango",
		Description: "Mango is the best",
		Price:       200,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQX_LYndiiyKxXWCxRmyp8hW5bCaTiM_PF45Q&s",
	}

	prd5 := &Product{
		ID:          5,
		Title:       "Grape",
		Description: "Grapes are sour",
		Price:       200,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRjT-J4hVoo8EfO_s5QauWDFW5QqXAza7NKrA&s",
	}

	r.productList = append(r.productList, prd1)
	r.productList = append(r.productList, prd2)
	r.productList = append(r.productList, prd3)
	r.productList = append(r.productList, prd4)
	r.productList = append(r.productList, prd5)

}