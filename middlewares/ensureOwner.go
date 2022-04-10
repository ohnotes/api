package middlewares

import (
    db "github.com/ohnotes/api/database"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
)

type Note struct {
    Owner string `json:"owner" bson:"owner"`
}

func EnsureOwner(c *gin.Context) {
    var note Note

    token := c.MustGet("token")

    if err := db.Notes.FindOne(db.Ctx, bson.M{"id": c.Param("id")}).Decode(&note); err != nil {
        c.JSON(404, gin.H {
            "message": "Note was not found",
        })
    }

    if note.Owner != token {
        c.JSON(403, gin.H {
            "message": "You don't have permission to do that.",
        })

        c.Abort()
        return
    }

    c.Next()
}
