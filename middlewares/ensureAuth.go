package middlewares

import (
    "strings"

    db "github.com/ohnotes/api/database"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
)

func EnsureAuth(c *gin.Context) {
    if len(c.Request.Header["Authorization"]) != 0 {
        token := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
        
        if err := db.Users.FindOne(db.Ctx, bson.M{"token": token}).Err(); err != nil {
            c.JSON(403, gin.H {
                "message": "Invalid token.",
            })

            return
        }

        c.Set("token", token)
        c.Next()
    
    } else {
        c.JSON(403, gin.H {
            "message": "You don't have permission to do that.",
        })
       
        c.Abort()
        return
    }
}
