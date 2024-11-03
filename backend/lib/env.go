package lib

import "time"

var JwtSecret = []byte("baba-yaga-secret")

var ExpirationTime = time.Now().Add(time.Hour * 24 * 7).Unix()

var Port = ":4000"

var DbUrl = "mongodb://localhost:27017"