package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JSONTransformer(c *gin.Context, message []ProviderMessage) {
	c.JSON(http.StatusOK, message)
}

func XMLTransformer(c *gin.Context, message []ProviderMessage) {
	c.XML(http.StatusOK, message)
}

func TSVTransformer() {

}
