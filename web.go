package random

import (
	"math/rand"
)

func Email() string {

	name := randSeq(8)
	domain := randSeq(5)

	return name + "@" + domain + ".com"

}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
