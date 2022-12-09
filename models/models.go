package models

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

type User struct {
    Email, Name, Password string
}

type Appointment struct {
    Nhs, MedicalSpecialty, Date string
}

func ConnectDB() {
    DB_USER := os.Getenv("DB_USER")
    DB_PASS := os.Getenv("DB_PASS")

    DB, _ = sql.Open("mysql", DB_USER + ":" + DB_PASS + "@/testdb")

    DB.Ping()
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

func ValidRegister(username string, nhs string) bool {
    tx, err := DB.Prepare("SELECT username, nhs FROM Patients WHERE username = ? or nhs = ?;")
    if err != nil {
        return false
    }
    defer tx.Close()

    rows, err := tx.Query(username, nhs)
    if err != nil {
        return false
    }

    if rows.Next() {
        return false
    }

    return true
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
