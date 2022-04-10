package services

import (
    "os"
    "time"

    db "github.com/ohnotes/api/database"
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt"
    "github.com/google/uuid"
)

type UserCreate struct {
    ID string `json:"id" bson:"id"`
    Token string `json:"token" bson:"token"`
    CreatedAt int `json:"createdAt" bson:"createdAt"`
}

func GenerateService(c *gin.Context) {
    id := uuid.New().String()

    gen := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims {
        "user": id,
        "iat": time.Now().UnixMicro(),
    })

    token, err := gen.SignedString([]byte(os.Getenv("SECRET")))
    if err != nil {
        c.JSON(400, gin.H {
            "message": "An error was occurred during generate.",
        })

        return
    }

    _, err = db.Users.InsertOne(db.Ctx, UserCreate {
        ID: id,
        Token: token,
        CreatedAt: int(time.Now().UnixMicro()),
    })
    if err != nil {
        c.JSON(400, gin.H {
            "message": "An error was occurred during generate.",
        })

        return
    }

    c.JSON(200, gin.H {
        "id": id,
        "token": token,
    })
}
