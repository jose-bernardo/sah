package models

import (
    "database/sql"
    "github.com/go-sql-driver/mysql"
    "sah/helpers"
    "log"
    "os"
)

type User struct {
    Nhs, Email, Name, Password string
}

type Appointment struct {
    Nhs, MedicalSpecialty, Date string
}

var DB *sql.DB

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
        Addr: "server.localhost:3306",
        //TLSConfig: "custom",
    }


    DB, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal("Error connecting to the database.")
    }
}

func GetUserPass(nhs string) (string, error) {
    tx, err := DB.Prepare("SELECT password FROM Patients WHERE nhs = ?;")
    if err != nil {
        return "", err
    }
    defer tx.Close()

    rows, err := tx.Query(nhs)
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

func GetUser(nhs string) User {
    tx, err := DB.Prepare("SELECT nhs, email, name, password FROM Patients WHERE nhs = ?;")
    if err != nil {
        log.Fatal(err.Error())
    }
    defer tx.Close()

    rows, err := tx.Query(nhs)
    if err != nil {
        log.Fatal(err.Error())
    }

    var user User
    if rows.Next() {
        err := rows.Scan(&user.Nhs, &user.Email, &user.Name, &user.Password)
        if err != nil {
            log.Fatal(err.Error())
        }
    }

    if rows.Err() != nil {
        log.Fatal(err.Error())
    }

    return user;
}

func ValidRegister(nhs string) bool {
    tx, err := DB.Prepare("SELECT nhs FROM Patients WHERE nhs = ?;")
    if err != nil {
        log.Fatal(err.Error())
    }
    defer tx.Close()

    rows, err := tx.Query(nhs)
    if err != nil {
        log.Fatal(err.Error())
    }

    if rows.Next() {
        return false
    }

    return true
}

func RegisterUser(email string, name string, nhs string, password string) (error) {
    tx, err := DB.Prepare("INSERT INTO Patients (email, name, nhs, password) VALUES ( ?, ?, ?, ? );")
    if err != nil {
        return err
    }
    defer tx.Close()

    _, err = tx.Exec(email, name, nhs, password)
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
