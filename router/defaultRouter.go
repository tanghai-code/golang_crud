package router

import (
	"github.com/gin-gonic/gin"
	"golang_crud/controller"
)

// SetAllRouter put on a gin engine to set all router
func SetAllRouter(r *gin.Engine) {
	// Set All User Entity Request Handler
	controller.SetAllUserController(r)
}
