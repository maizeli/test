package service

import (
	"math/rand"
)

func GetName() string {
	if rand.Int() > 10000 {
		return "fix fix fix"
	}
	return "Hello,WorldV3!"
}
