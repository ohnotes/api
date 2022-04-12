package services

import (
    "strings"

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
    var note Note
    
    err := c.BindJSON(&response)
    if err != nil {
        c.JSON(400, gin.H {
            "message": "An error was occurred during note update.",
        })

        return
    }
    
    db.Notes.FindOne(db.Ctx, bson.M{"id": response.ID}).Decode(&note)
    
    if len(c.Request.Header["Authorization"]) != 0 {
        token := strings.Split(c.Request.Header["Authorization"][0], " ")[1]

        if note.Owner != token {
            c.JSON(403, forbidden)

            return
        }

    } else {
        c.JSON(403, forbidden)

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
