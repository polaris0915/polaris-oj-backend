package main

import (
	"polaris-oj-backend/router"
)

func main() {

	r := router.Router()
	r.Run(":8000")
}
