package database


var productList []Product

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

func Store(p Product) Product {
	p.ID = len(productList) + 1
	productList = append(productList, p)
	return p
}

func List() []Product {
	return productList
}

func Get(productID int) *Product {
	for _, product := range productList {
		if product.ID == productID {
			return &product
		}
	}
	return nil
}

func Update(product Product) {
	for idx, p := range productList {
		if p.ID == product.ID {
			productList[idx] = product
		}
	}
}

func Delete(productID int) {
	var tempList []Product

	for _, p := range productList {
		if p.ID != productID {
			tempList = append(tempList, p)
		}
	}
	productList = tempList

}

func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange contains vitamin C",
		Price:       350,
		ImgUrl:      "https://www.allrecipes.com/thmb/y_uvjwXWAuD6T0RxaS19jFvZyFU=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/GettyImages-1205638014-2000-d0fbf9170f2d43eeb046f56eec65319c.jpg",
	}

	prd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Apple has lot of minerals",
		Price:       200,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQD7ANGmYuBLltBJmAFrhfX5p0p9UdzJqfiYlVPIIXplE_RAXFsdTeuru13CUP_L4Nv1ew&usqp=CAU",
	}

	prd3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "Bana has high carb",
		Price:       200,
		ImgUrl:      "https://img.freepik.com/free-vector/vector-ripe-yellow-banana-bunch-isolated-white-background_1284-45456.jpg?semt=ais_incoming&w=740&q=80",
	}

	prd4 := Product{
		ID:          4,
		Title:       "Mango",
		Description: "Mango is the best",
		Price:       200,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQX_LYndiiyKxXWCxRmyp8hW5bCaTiM_PF45Q&s",
	}

	prd5 := Product{
		ID:          5,
		Title:       "Grape",
		Description: "Grapes are sour",
		Price:       200,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRjT-J4hVoo8EfO_s5QauWDFW5QqXAza7NKrA&s",
	}

	productList = append(productList, prd1)
	productList = append(productList, prd2)
	productList = append(productList, prd3)
	productList = append(productList, prd4)
	productList = append(productList, prd5)

}
