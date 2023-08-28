package utils

import "math/rand"

var base62Characters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func GenerateBase62String(length int) string {
	var output string
	for len(output) < length {
		output += string([]rune(base62Characters)[rand.Intn(62)])
	}
	return output
}
