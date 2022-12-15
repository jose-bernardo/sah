package helpers

import (
    "strings"
    "golang.org/x/crypto/bcrypt"
    "crypto/tls"
    "crypto/x509"
    "log"
    "io/ioutil"
)

func CreateTLSConf() tls.Config {
    rootCertPool := x509.NewCertPool()
    pem, err := ioutil.ReadFile("./certs/ca.pem")
    if err != nil {
        log.Fatal(err)
    }

    if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
        log.Fatal("Failed to open PEM.")
    }

    clientCert := make([]tls.Certificate, 0, 1)
    certs, err := tls.LoadX509KeyPair("./certs/client-cert.pem", "./certs/client-key.pem")
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
