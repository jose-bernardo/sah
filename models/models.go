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
    err := mysql.RegisterTLSConfig("DBConfig", &tlsConf)
    if err != nil {
        log.Fatal("Error registering TLS configuration.")
    }

    cfg := mysql.Config {
        User: os.Getenv("DB_USER"),
        Passwd: os.Getenv("DB_PASS"),
        DBName: "testdb",
        Net: "tcp",
        Addr: "192.168.2.1:3306",
        //Addr: "localhost:3306",
	TLSConfig: "DBConfig",

    }

	print(cfg.FormatDSN())
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

func NewAppointment(nhs string, date string, medicalSpecialty string) (error) {
    
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

func GetAppointmentForUser(nhs string) ([]Appointment, error){
    rows, err := DB.Query("SELECT * FROM Appointments WHERE patientNhs = ?;", nhs)

    if err != nil {
        return nil, err
    }

    defer rows.Close()

    var appointments []Appointment
    var id int
    for rows.Next() {
        var app Appointment
        if err := rows.Scan(&id, &app.Date, &app.Nhs, &app.MedicalSpecialty); err != nil {
            return appointments, err
        }

        appointments = append(appointments, app)
    }

    if err = rows.Err(); err != nil {
        return appointments, err
    }

    return appointments, nil
}

