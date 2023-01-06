package helpers

import (
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"strings"
	"golang.org/x/crypto/bcrypt"
)


func CreateTLSConf() tls.Config {
    rootCertPool := x509.NewCertPool()
    pem, err := ioutil.ReadFile("./certs/bd-client/ca.pem")
    if err != nil {
        log.Fatal(err)
    }

    if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
        log.Fatal("Failed to open PEM.")
    }

    clientCert := make([]tls.Certificate, 0, 1)
    certs, err := tls.LoadX509KeyPair("./certs/bd-client/bd-client-cert.pem", "./certs/bd-client/bd-client-key.pem")
    if err != nil {
        log.Fatal(err)
    }

    clientCert = append(clientCert, certs)
    return tls.Config{
        RootCAs: rootCertPool,
        Certificates: clientCert,
    }
}


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


func EmptyRegisterParams(email string, name string, nhs string, password string) bool {
    return strings.Trim(email, " ") == "" || strings.Trim(name, " ") == "" || strings.Trim(nhs, " ") == "" || strings.Trim(password, " ") == ""
}


func EmptyNhsOrPass(nhs string, password string) bool {
    return strings.Trim(nhs, " ") == "" || strings.Trim(password, " ") == ""
}


func GenerateOTP() (string, error) {
    buffer := make([]byte, 6)
    _, err := rand.Read(buffer)
    if err != nil {
        return "", err
    }

    var charset = []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

    for i := 0; i < 6; i++ {
        buffer[i] = charset[int(buffer[i]%6)]
    }

    return string(buffer), nil
}


/*
func ValidateOTP(otp string, attempt string, created string) bool {
    if otp == attempt {
        expired, err := time.Parse("2006-01-02 15:04:05", created)
        if err != nil {
            return false
        }
        if time.Now().Before(expired.Add(time.Minute)) {
            return true
        }
    }
    return false
}
*/
