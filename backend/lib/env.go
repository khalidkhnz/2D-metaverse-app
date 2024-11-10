package lib

import "time"

var JwtSecret = []byte("baba-yaga-secret")

var ExpirationTime = time.Now().Add(time.Hour * 24 * 7).Unix()

var Port = ":4000"

var ApiPrefix = "/api/v1"

var Base = "localhost:4000"

var URL = "http://" + Base

var FrontEndProxyURL = "http://localhost:3000"

var DbUrl = "mongodb+srv://khalidkhnz:khalidkhnz@khalid-cluster.ttkcc.mongodb.net/metaverse?retryWrites=true&w=majority&appName=khalid-cluster"