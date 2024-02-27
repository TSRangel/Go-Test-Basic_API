package handlers

import (
	"encoding/json"
	"github.com/TSRangel/Go-Test-Basic_API/internal/dto"
	"github.com/TSRangel/Go-Test-Basic_API/internal/entities/product"
	"github.com/TSRangel/Go-Test-Basic_API/internal/infra/database"
	"github.com/TSRangel/Go-Test-Basic_API/pkg/tools"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type ProductHandler struct {
	DB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{DB: db}
}

// Create Product 	godoc
// @Summary Create 	product
// @Description 	Create product
// @Tags 			products
// @Accpet 			json
// @Produce 		json
// @Param 			request 	body 	dto.CreateProductInput 	true 	"product request"
// @Success 		201
// @Failure 		500 		{object} Error
// @Router 			/products 	[post]
// @Security 		ApiKeyAuth
func (ph *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var productDTO dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&productDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	newProduct, err := product.NewProduct(productDTO.Name, productDTO.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = ph.DB.Create(newProduct)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// ListProducts 	godoc
// @Summary 		List products
// @Description 	Get all products
// @Tags 			products
// @Accept 			json
// @Produce 		json
// @Param 			page query string false "page number"
// @Param 			limit query string false "limit"
// @Success 		200 {array} product.Product
// @Failure 		404 
// @Failure 		500 {object} Error
// @Router 			/products [get]
// @Security 		ApiKeyAuth
func (ph *ProductHandler) ListAllProducts(w http.ResponseWriter, r *http.Request) {
	page := chi.URLParam(r, "page")
	limit := chi.URLParam(r, "limit")
	sort := chi.URLParam(r, "sort")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 0
	}
	products, err := ph.DB.FindAll(pageInt, limitInt, sort)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// GetProduct 		godoc
// @Summary 		Get a product
// @Description 	Get a product
// @Tags 			products
// @Accpet 			json
// @Produce 		json
// @Param 			id path string true "product ID" Format(uuid)
// @Success 		200 {object} product.Product
// @Failure 		404
// @Failure 		500 {object} Error
// @Router 			/products/{id} [get]
// @Security 		ApiKeyAuth
func (ph *ProductHandler) ListProductByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	product, err := ph.DB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// UpdateProduct 	godoc
// @Summary 		Update a product
// @Description 	Update a product
// @Tags 			products
// @Accpet 			json
// @Produce 		json
// @Param 			id path string true "product ID" Format(uuid)
// @Param 			request body dto.CreateProductInput true "product request"
// @Success 		200
// @Failure 		404
// @Failure 		500 {object} Error
// @Router 			/products/{id} [put]
// @Security 		ApiKeyAuth
func (ph *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err := tools.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	searchedProduct, err := ph.DB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewDecoder(r.Body).Decode(&searchedProduct)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	err = ph.DB.Update(searchedProduct)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// DeleteProduct 	godoc
// @Summary 		Delete a product
// @Description 	Delete a product
// @Tags 			products
// @Accept 			json
// @Produce 		json
// @Param 			id path string true "product ID" Format(uuid)
// @Success 		200
// @Failure 		404
// @Failure 		500 {object} Error
// @Router 			/products/{id} [delete]
// @Security 		ApiKeyAuth
func (ph *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err := tools.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = ph.DB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = ph.DB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.WriteHeader(http.StatusOK)
}
