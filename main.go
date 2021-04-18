package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"./handlers"
	"./middlewares"
	"./services"
)

const SessionTokenLength = 16
const SessionDuration = 10 * time.Minute

var cache services.Cache

func main() {
	Init()

	router := http.NewServeMux()
	fs := http.FileServer(http.Dir("./static"))
	router.Handle("/static/", http.StripPrefix("/static/", fs))

	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/signin", handlers.Signin)
	router.Handle("/home", middlewares.RequireAuthentication(http.HandlerFunc(handlers.Home)))
	router.Handle("/logout", middlewares.RequireAuthentication(http.HandlerFunc(handlers.Logout)))

	log.Fatal(http.ListenAndServe(":8000", router))
}

func Init() {
	rand.Seed(time.Now().UnixNano())
	services.SessionDuration = SessionDuration
	cache = services.NewCache()
	middlewares.Cache = cache
	handlers.Cache = cache
	handlers.SessionDuration = SessionDuration
	handlers.SessionTokenLength = SessionTokenLength
}
