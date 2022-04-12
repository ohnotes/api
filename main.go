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
        AllowHeaders: []string{"Authorization", "Content-Type"},
        AllowOrigins: []string{"https://ohnotes.vercel.app", "*"},
    }))

    app.GET("/note/:id", services.NoteService)
    app.GET("/getNotes", middlewares.EnsureAuth, services.GetNotesService)
    app.GET("/destructive/:id", services.UpdateDestructive)
    app.GET("/wipe", middlewares.EnsureAuth, services.WipeService)
    app.GET("/deleteall", middlewares.EnsureAuth, services.DeleteAllService)
    app.POST("/generate", services.GenerateService)
    app.POST("/new", middlewares.EnsureAuth, services.NewService)
    app.POST("/update/:id", services.UpdateNoteService)
    app.POST("/update/:id/settings", middlewares.EnsureAuth, middlewares.EnsureOwner, services.UpdateSettingsService)
    app.POST("/delete/:id", middlewares.EnsureAuth, middlewares.EnsureOwner, services.DeleteService)

    app.Run(":" + os.Getenv("PORT"))
}
