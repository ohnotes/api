package services

import (
    "github.com/gin-gonic/gin"
    db "github.com/ohnotes/api/database"
    "go.mongodb.org/mongo-driver/bson"
)

func WipeService(c *gin.Context) {
    token := c.MustGet("token")

    _, err := db.Notes.DeleteMany(db.Ctx, bson.M{"owner": token})
    if err != nil {
        c.JSON(400, gin.H {
            "message": "Failed to wipe data.",
        })

        return
    }
    
    if err := db.Users.FindOneAndDelete(db.Ctx, bson.M{"token": token}).Err(); err != nil {
        c.JSON(400, gin.H {
            "message": "Failed to wipe data.",
        })

        return
    }
}
