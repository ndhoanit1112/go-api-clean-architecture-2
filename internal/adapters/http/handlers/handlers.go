package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ndhoanit1112/go-api-clean-architecture-2/internal/adapters/http/handlers/product"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	product.SetupRoutes(r, db)
}
