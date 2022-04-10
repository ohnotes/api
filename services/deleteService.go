package services

import (
    "github.com/gin-gonic/gin"
    db "github.com/ohnotes/api/database"
    "go.mongodb.org/mongo-driver/bson"
)


func DeleteService(c *gin.Context) {
    id := c.Param("id")

    if err := db.Notes.FindOneAndDelete(db.Ctx, bson.M{"id": id}).Err(); err != nil {
        c.JSON(404, notFound)

        return
    }
}
