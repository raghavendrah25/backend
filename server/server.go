package server

import (
	"net/http"

	"github.com/Rhymond/go-money"
	"github.com/gin-gonic/gin"
	"github.com/raghavendrah25/backend/internal/category"
	"github.com/raghavendrah25/backend/internal/product"
)

type Server struct {
	engine *gin.Engine
}

type Config struct{}

func New(config Config) *Server {
	engine := gin.Default()

	router := &Server{
		engine: engine,
	}

	engine.GET("/products", router.Products)
	engine.GET("/categories", router.Categories)
	return router
}

func (s *Server) Categories(c *gin.Context) {
	cat := []category.Category{
		{
			ID:          "42",
			Name:        "Book",
			Description: "A book",
		},
		{
			ID:          "41",
			Name:        "Book",
			Description: "A New book",
		},
	}
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, cat)
}

func (s *Server) Products(c *gin.Context) {
	products := []product.Product{
		{
			ID:               "42",
			Name:             "Book",
			Description:      "A book",
			PriceVATExcluded: money.New(100, money.EUR),
			PriceVAT:         nil,
		},
		{
			ID:               "41",
			Name:             "Book",
			Description:      "A book",
			PriceVATExcluded: money.New(800, money.EUR),
			PriceVAT:         nil,
		},
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, products)
}

func (s *Server) Run() error {
	err := s.engine.RunTLS(":8080", "cert.pem", "key.pem")
	if err != nil {
		return err
	}
	return nil
}
