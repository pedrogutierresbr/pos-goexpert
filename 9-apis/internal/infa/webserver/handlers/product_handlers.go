package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/pedrogutierresbr/pos-goexpert/9-apis/internal/dto"
	"github.com/pedrogutierresbr/pos-goexpert/9-apis/internal/entity"
	"github.com/pedrogutierresbr/pos-goexpert/9-apis/internal/infa/database"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProduct
	err := json.NewDecoder(r.Body).Decode(&product) //decodifica e salva no product
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	p, err := entity.NewProduct(product.Name, product.Price) // cria um novo produto e salva no p
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.ProductDB.Create(p) // salva o produto criado no db
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
