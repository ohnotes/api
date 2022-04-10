package main

import (
    "os"

    "github.com/ohnotes/api/services"
    "github.com/ohnotes/api/middlewares"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

func main() {
    app := gin.Default()
    app.Use(cors.New(cors.Config {
        AllowOrigins: []string{"*"},
        AllowHeaders: []string{"Authorization", "Content-Type"},
    }))

    app.GET("/note/:id", middlewares.EnsureAuth, services.NoteService)
    app.GET("/getNotes", middlewares.EnsureAuth, services.GetNotesService)
    app.GET("/destructive/:id", services.UpdateDestructive)
    app.POST("/generate", services.GenerateService)
    app.POST("/new", middlewares.EnsureAuth, services.NewService)
    app.POST("/update/:id", services.UpdateNoteService)
    app.POST("/update/:id/settings", middlewares.EnsureAuth, middlewares.EnsureOwner, services.UpdateSettingsService)
    app.POST("/delete/:id", middlewares.EnsureAuth, middlewares.EnsureOwner, services.DeleteService)

    app.Run(":" + os.Getenv("PORT"))
}
