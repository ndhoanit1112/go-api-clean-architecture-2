package product

import (
	"github.com/gin-gonic/gin"
	"github.com/ndhoanit1112/go-api-clean-architecture-2/internal/adapters/mysql/repositories"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	h := &productHandlers{}
	repo := repositories.InitProductRepo(db)

	routes := r.Group("/products")
	routes.GET("/", h.GetProducts(repo))
	routes.POST("/", h.CreateProduct(repo))
	routes.GET("/:id", h.GetProduct(repo))
}
