package main

import (
	"fmt"
	"log"
	"time"

	"github.com/pixconf/pixconf/internal/autocert"
)

func main() {
	apps := []string{"api", "hub", "secrets"}
	instance := fmt.Sprintf("%x", time.Now().Unix())

	if err := autocert.MakeSelfSigned(apps, instance); err != nil {
		log.Fatal(err)
	}
}
