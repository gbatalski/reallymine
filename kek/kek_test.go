package kek

import (
	"encoding/hex"
	"fmt"
	"log"
	"testing"
	"time"
)

func TestKEK(t *testing.T) {
	password := "abc123"
	start := time.Now()
	passwordHash := FromPassword(password)
	elapsed := time.Since(start)
	log.Printf("Hashing password %s took %s", password, elapsed)
	fmt.Print(hex.Dump(passwordHash))
}
