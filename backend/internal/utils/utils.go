package utils

import (
	"strings"
	"math/rand"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func ContainString(listStr string, delimiter string, target string) bool {
	list := strings.Split(listStr, delimiter)
	for _, item := range list {
		if item == target {
			return true
		}
	}
	return false
}

func RandomString(length int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    b := make([]byte, length)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}