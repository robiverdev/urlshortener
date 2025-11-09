package main

import "math/rand"

func generateShortCode() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" // All possible characters for the code
	const length = 6                                                                 // Empty byte array of length 6

	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = charset[rand.Intn(len(charset))] // Random number between 0 and charset length
	}
	return string(result) // Convert byte string and return
}
