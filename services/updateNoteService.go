package services

import (
    db "github.com/ohnotes/api/database"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
)

type UpdateNoteResponse struct {
    ID string `json:"id"`
    Text string `json:"text"`
    LastUpdate string `json:"lastUpdate"`
}

func UpdateNoteService(c *gin.Context) {
    var response UpdateNoteResponse
    
    err := c.BindJSON(&response)
    if err != nil {
        c.JSON(400, gin.H {
            "message": "An error was occurred during note update.",
        })

        return
    }
    
    update, err := db.Notes.UpdateOne(db.Ctx, bson.M{"id": response.ID}, bson.M{"$set": bson.M{"text": response.Text}})
    if err != nil || update.MatchedCount == 0 {
        c.JSON(400, gin.H {
            "message": "An error was occurred during note update.",
        })

        return
    }
    
    update, err = db.Notes.UpdateOne(db.Ctx, bson.M{"id": response.ID}, bson.M{"$set": bson.M{"lastUpdate": response.LastUpdate}})
    if err != nil || update.MatchedCount == 0 {
        c.JSON(400, gin.H {
            "message": "An error was occurred during note update.",
        })

        return
    }

    c.JSON(200, gin.H {
        "message": "Successfully updated.",
    })
}
