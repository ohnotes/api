package services

import (
    db "github.com/ohnotes/api/database"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
)

type DestructiveResponse struct {
    Destructive bool `json:"destructive"`
    Turns int `json:"turns"`
}

func UpdateDestructive(c *gin.Context) {
    var note DestructiveResponse

    id := c.Param("id")

    if err := db.Notes.FindOne(db.Ctx, bson.M{"id": id}).Decode(&note); err != nil {
        c.JSON(404, notFound)

        return
    }

    if note.Destructive {
        if note.Turns < 1 {
            _, err := db.Notes.DeleteOne(db.Ctx, note)
            if err != nil {
                c.JSON(400, gin.H {
                    "message": "Failed to delete note.",
                })

                return
            }
        
        } else {
            update, err := db.Notes.UpdateOne(db.Ctx, bson.M{"id": id}, bson.M{"$set": bson.M{"turns": note.Turns - 1}})
            if err != nil || update.MatchedCount == 0 {
                c.JSON(400, gin.H {
                    "message": "Failed to update note.",
                })

                return
            }
        }
    }
}
