package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	u1 := uuid.New()
	fmt.Printf("uuid v4: %s\n", u1)
}
