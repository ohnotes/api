package services

import "github.com/gin-gonic/gin"

var notFound gin.H = gin.H {
    "message": "Note was not founded.",
}

var forbidden gin.H = gin.H {
    "message": "You don't have permission to do that.",
}
