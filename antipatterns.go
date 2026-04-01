package main

import (
	"crypto/des"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

// Command injection — user input passed directly to shell
func runUserCommand(input string) error {
	cmd := exec.Command("sh", "-c", input)
	return cmd.Run()
}

// Hardcoded credentials
const (
	dbUser     = "admin"
	dbPassword = "super_secret_p@ssw0rd!"
	apiKey     = "sk-live-4eC39HqLyjWDarjtT1zdp7dc"
)

// DES is deprecated and insecure
func encryptDES(key, plaintext []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	ciphertext := make([]byte, len(plaintext))
	block.Encrypt(ciphertext, plaintext)
	return ciphertext, nil
}

// Double MD5 is still insecure
func doubleMD5(data string) string {
	first := md5.Sum([]byte(data))
	second := md5.Sum(first[:])
	return hex.EncodeToString(second[:])
}

// World-writable file permissions
func writeInsecureFile(path string, data []byte) error {
	return os.WriteFile(path, data, 0777)
}

// Error value ignored
func ignoreErrors() {
	f, _ := os.Open("/etc/passwd")
	f.Close()

	strconv.Atoi("not_a_number")

	fmt.Fprintf(os.Stdout, "hello")
}

// Empty critical section
func emptyBranch(x int) string {
	if x > 10 {
		// TODO: handle this case
	} else if x > 5 {
		return "medium"
	} else {
		return "low"
	}
	return ""
}

// Potential slice out of bounds
func unsafeSliceAccess(s []int) int {
	return s[0] + s[len(s)-1]
}

// Comparing floats with == (imprecise)
func floatCompare(a, b float64) bool {
	return a == b
}
