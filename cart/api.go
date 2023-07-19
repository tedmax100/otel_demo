package cart

import (
	"otel_demo/cart/repository"
	// "otel_demo/cart/entity"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewCartController),
)

type CartController struct {
	repo repository.CartRepository
}

func NewCartController(repo repository.CartRepository) *CartController {
	return &CartController{repo: repo}
}

func (c *CartController) AddItemToCart(ctx *gin.Context) {

}

func (c *CartController) GetCart(ctx *gin.Context) {

}
