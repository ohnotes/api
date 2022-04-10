package database

import (
    "context"
    "os"
    
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var (
    opts = options.Client().ApplyURI(os.Getenv("MONGO"))
    Ctx = context.TODO()
    client, _ = mongo.Connect(Ctx, opts)
    Notes = client.Database("ohnotes").Collection("notes")
    Users = client.Database("ohnotes").Collection("users")
)
