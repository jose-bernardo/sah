package models

import (
    "database/sql"
    "sah/helpers"
    "log"
    "os"

    "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

type User struct {
    Email, Name, Password string
}

type Appointment struct {
    Nhs, MedicalSpecialty, Date string
}

func ConnectDB() {
    tlsConf := helpers.CreateTLSConf()
    err := mysql.RegisterTLSConfig("custom", &tlsConf)
    if err != nil {
        log.Fatal("Error registering TLS configuration.")
    }

    cfg := mysql.Config {
        User: os.Getenv("DB_USER"),
        Passwd: os.Getenv("DB_PASS"),
        DBName: "testdb",
        Net: "tcp",
        //Addr: "192.168.2.1:3306",
        Addr: "localhost:3306",
    }


    DB, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal("Error connecting to the database.")
    }
}

func GetUserPass(username string) (string, error) {
    tx, err := DB.Prepare("SELECT password FROM Patients WHERE username = ?;")
    if err != nil {
        return "", err
    }
    defer tx.Close()

    rows, err := tx.Query(username)
    if err != nil {
        return "", err
    }

    password := ""
    if rows.Next() {
        err := rows.Scan(&password)
        if err != nil {
            return "", err
        }
    }

    if rows.Err() != nil {
        return "", err
    }

    return password, nil
}

// TODO check for errors
func ValidRegister(username string, nhs string) (bool, error) {
    tx, err := DB.Prepare("SELECT username, nhs FROM Patients WHERE username = ? or nhs = ?;")
    if err != nil {
        return false, err
    }
    defer tx.Close()

    rows, err := tx.Query(username, nhs)
    if err != nil {
        return false, err
    }

    if rows.Next() {
        return false, nil
    }

    return true, nil
}

func RegisterUser(username string, name string, nhs string, password string) (error) {
    tx, err := DB.Prepare("INSERT INTO Patients (username, name, nhs, password) VALUES ( ?, ?, ?, ? );")
    if err != nil {
        return err
    }
    defer tx.Close()

    _, err = tx.Exec(username, name, nhs, password)
    if err != nil {
        return err
    }

    return nil
}

func NewAppointment(date string, nhs string, medicalSpecialty string) (error) {
    tx, err := DB.Prepare("INSERT INTO Appointments (date, patientNhs, medicalSpecialty) VALUES ( ?, ?, ? );")
    if err != nil {
        return err
    }
    defer tx.Close()

    _, err = tx.Exec(date, nhs, medicalSpecialty)
    if err != nil {
        return err
    }

    return nil
}
