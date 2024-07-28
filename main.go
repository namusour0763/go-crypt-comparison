package main

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/scrypt"
)

func main() {
	password := "password" // パスワード
	iterations := 100      // 計算回数

	// MD5
	md5Duration := benchmark(iterations, func() {
		md5.Sum([]byte(password))
	})

	// SHA-256
	sha256Duration := benchmark(iterations, func() {
		sha256.Sum256([]byte(password))
	})

	// Bcrypt
	bcryptDuration := benchmark(iterations, func() {
		_, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Printf("Bcryptエラー: %v\n", err)
		}
	})

	// Scrypt
	scryptDuration := benchmark(iterations, func() {
		_, err := scrypt.Key([]byte(password), []byte("salt"), 32768, 8, 1, 32)
		if err != nil {
			fmt.Printf("Scryptエラー: %v\n", err)
		}
	})

	fmt.Printf("MD5の平均時間 (%d回): %v\n", iterations, md5Duration)
	fmt.Printf("SHA-256の平均時間 (%d回): %v\n", iterations, sha256Duration)
	fmt.Printf("Bcryptの平均時間 (%d回): %v\n", iterations, bcryptDuration)
	fmt.Printf("Scryptの平均時間 (%d回): %v\n", iterations, scryptDuration)
}

func benchmark(iterations int, hashFunc func()) time.Duration {
	var totalDuration time.Duration
	for i := 0; i < iterations; i++ {
		start := time.Now()
		hashFunc()
		totalDuration += time.Since(start)
	}
	return totalDuration / time.Duration(iterations)
}
