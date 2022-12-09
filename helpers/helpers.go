package helpers

import (
    "strings"
    "golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
    hash, err := bcrypt.GenerateFromPassword([]byte(password), 8)
    if err != nil {
        return "", err
    }

    return string(hash), nil
}

func CheckPassword(hash string, password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    if err != nil {
        return false
    }
    return true
}

func EmptyRegisterParams(username string, name string, nhs string, password string) bool {
    return strings.Trim(username, " ") == "" || strings.Trim(name, " ") == "" || strings.Trim(nhs, " ") == "" || strings.Trim(password, " ") == ""
}

func EmptyUserOrPass(username string, password string) bool {
    return strings.Trim(username, " ") == "" || strings.Trim(password, " ") == ""
} 
