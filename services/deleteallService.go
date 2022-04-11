package services

import (
    db "github.com/ohnotes/api/database"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
)

func DeleteAllService(c *gin.Context) {
    token := c.MustGet("token")

    del, err := db.Notes.DeleteMany(db.Ctx, bson.M{"owner": token})
    if del.DeletedCount == 0 {
        c.JSON(404, gin.H {
            "message": "No one note was founded.",
        })

        return
    }

    if err != nil {
        c.JSON(400, gin.H {
            "message": "Failed to delete notes",
        })
    }
}
