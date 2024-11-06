package lib

import "time"

var JwtSecret = []byte("baba-yaga-secret")

var ExpirationTime = time.Now().Add(time.Hour * 24 * 7).Unix()

var Port = ":4000"

var ApiPrefix = "/api/v1"

var URL = "http://localhost:4000"

var FrontEndProxyURL = "http://localhost:3000"

var DbUrl = "mongodb://localhost:27017"