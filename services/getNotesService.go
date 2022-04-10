package services

import (
    "github.com/gin-gonic/gin"
    db "github.com/ohnotes/api/database"
    "go.mongodb.org/mongo-driver/bson"
)

type GetNotesResponse struct {
    ID string `json:"id"`
}

func GetNotesService(c *gin.Context) {
    var notes []GetNotesResponse
    token := c.MustGet("token")

    filter, err := db.Notes.Find(db.Ctx, bson.M{"owner": token})
    if err != nil {
        c.JSON(400, gin.H {
            "message": "An error was occurred during insert.",
        })

        return
    }
    
    if err = filter.All(db.Ctx, &notes); err != nil {
        c.JSON(400, gin.H {
            "message": "An error was occurred during insert.",
        })

        return
    }

    c.JSON(200, notes)
}
