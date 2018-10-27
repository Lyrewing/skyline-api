package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomNumber(start int, end int) uint {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := r.Intn((end - start)) + start
	return uint(num)
}
