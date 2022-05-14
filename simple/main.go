package main

import (
	"dev_test/apis"
	"dev_test/engine"
)

func main() {
	start := engine.New()
	start.Post("/v1/signup", apis.Signup)
	start.Post("/v1/signin", apis.Signin)
	start.Get("/v1/profile", apis.Profile)
	start.Post("/v1/profile/update", apis.ProfileUpdate)
	start.Run(":8080")
}
