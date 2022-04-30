package user

import (
    "testing"
)

func TestHashEncryptDecrypt(t *testing.T) {
    testcases := []struct{
        title, password string
    } {
        { "3 characters", "abc" },
        { "6 characters", "abcdef" },
        { "11 characters", "abcdefghijk" },
        { "special characters", "!@#$%^&*()_+{}:\"<>?[];',./`~\\|'" },
    }
    for _, testcase := range testcases {
        t.Run(testcase.title, func(t *testing.T) {
            encryptedPassword, err := hashPassword(testcase.password)
            if err != nil {
                t.Fatalf("users.hashPassword(%v) throw error %v", testcase.password, err)
            }
            if !checkPasswordHash(testcase.password, encryptedPassword) {
                t.Fatalf("users.checkPasswordHash(%v, %v) return false",
                    testcase.password, encryptedPassword)
            }
        })
    }
}
