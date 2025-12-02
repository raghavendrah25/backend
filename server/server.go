package server

import (
	"fmt"
	"net/http"

	"github.com/Rhymond/go-money"
	"github.com/gin-gonic/gin"
	"github.com/raghavendrah25/backend/internal/category"
	"github.com/raghavendrah25/backend/internal/product"
)

type Server struct {
	engine *gin.Engine
	port   uint
}

type Config struct {
	Port uint
}

func (s *Server) CORSMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Next()
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
	c.JSON(http.StatusOK, cat)
}

func (s *Server) Products(c *gin.Context) {
	products := []product.Product{
		{
			ID:               "42",
			Name:             "Protein Powder",
			Description:      "Muscle Protein Powder",
			PriceVATExcluded: money.New(20, money.EUR),
			PriceVAT:         nil,
		},
		{
			ID:               "41",
			Name:             "AI & ML Book",
			Description:      "An Artificial Intelligence and Machine Learning International Book",
			PriceVATExcluded: money.New(50, money.EUR),
			PriceVAT:         nil,
		},
	}
	c.JSON(http.StatusOK, products)
}

func (s *Server) Run() error {
	err := s.engine.Run(fmt.Sprintf(":%d", s.port))
	if err != nil {
		return err
	}
	return nil
}

func New(config Config) *Server {
	engine := gin.Default()

	router := &Server{
		engine: engine,
		port:   config.Port,
	}
	engine.Use(router.CORSMiddleware)
	engine.GET("/products", router.Products)
	engine.GET("/categories", router.Categories)
	return router
}
