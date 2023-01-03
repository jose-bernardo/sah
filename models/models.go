package models

import (
	"database/sql"
	"log"
	"os"
	"sah/helpers"

	"github.com/go-sql-driver/mysql"
)

type User struct {
    Nhs, Email, Name, Password string
}

type Appointment struct {
    Nhs, MedicalSpecialty, Date string;
    Doctor sql.NullString;
    Room sql.NullInt32;
    State bool
}

type Otp struct {
    Value, Created string
}

var DB *sql.DB


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
        Addr: "localhost:3306",
	    //TLSConfig: "DBConfig",
    }

    DB, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal("Error connecting to the database.")
    }
}


func GetUser(nhs string) User {
    tx, err := DB.Prepare("SELECT nhs, email, name, password FROM Patients WHERE nhs = ?;")
    if err != nil {
        panic(err.Error())
    }
    defer tx.Close()

    rows, err := tx.Query(nhs)
    if err != nil {
        panic(err.Error())
    }

    var user User
    if rows.Next() {
        if err := rows.Scan(&user.Nhs, &user.Email, &user.Name, &user.Password); err != nil {
            panic(err.Error())
        }
    }

    if rows.Err() != nil {
        panic(err.Error())
    }

    return user;
}


func ValidRegister(nhs string, email string) bool {
    tx, err := DB.Prepare("SELECT nhs, email FROM Patients WHERE nhs = ? or email = ?;")
    if err != nil {
        panic(err.Error())
    }
    defer tx.Close()

    rows, err := tx.Query(nhs, email)
    if err != nil {
        panic(err.Error())
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


func GetMedicalSpecialties() ([]string, error){
    tx, err := DB.Prepare("SELECT name FROM MedicalSpecialty;")
    if err != nil {
        return nil, err
    }
    defer tx.Close()

    rows, err := tx.Query()
    if err != nil {
        return nil, err
    }

    var medicalSpecialties []string
    for rows.Next() {
        var ms string
        if err := rows.Scan(&ms); err != nil {
            return medicalSpecialties, err
        }

        medicalSpecialties = append(medicalSpecialties, ms)
    }

    if rows.Err() != nil {
        return medicalSpecialties, err
    }

    return medicalSpecialties, nil
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


func GetUserAppointments(nhs string) ([]Appointment, error){
    tx, err := DB.Prepare("SELECT date, a.medicalSpecialty, state, e.name, room FROM Appointments a LEFT JOIN Employee e ON doctorId = e.id WHERE patientNhs = ?;")
    if err != nil {
        return nil, err
    }
    defer tx.Close()

    rows, err := tx.Query(nhs)
    if err != nil {
        return nil, err
    }

    var appointments []Appointment
    for rows.Next() {
        var app Appointment
        if err := rows.Scan(&app.Date, &app.MedicalSpecialty, &app.State, &app.Doctor, &app.Room); err != nil {
            return appointments, err
        }

        appointments = append(appointments, app)
    }

    if rows.Err() != nil {
        return appointments, err
    }

    return appointments, nil
}


func SetOTP(nhs string, otp string) (error) {
    tx, err := DB.Prepare("REPLACE INTO Otp (patientNhs, otp, created) VALUES ( ?, ?, NOW());")
    if err != nil {
        return err
    }
    defer tx.Close()

    _, err = tx.Exec(nhs, otp)
    if err != nil {
        return err
    }

    return nil
}


/*
func GetOTP(nhs string) (Otp, error) {
    tx, err := DB.Prepare("SELECT otp, created FROM OTP WHERE patientNhs = ?;")
    if err != nil {
        return Otp{}, err
    }
    defer tx.Close()

    rows, err := tx.Query(nhs)
    if err != nil {
        return Otp{}, err
    }
 
    var otp Otp
    if rows.Next() {
        if err := rows.Scan(&otp.Value, &otp.Created); err != nil {
            return otp, err
        }
    }

    if rows.Err() != nil {
        return Otp{}, err
    }

    return otp, nil;
}


func DeleteOTP(nhs string) error {
    tx, err := DB.Prepare("DELETE FROM Otp WHERE patientNhs = ?")
    if err != nil {
        return err
    }
    defer tx.Close()

    _, err = tx.Exec(nhs)
    if err != nil {
        return err
    }

    return nil
}
*/
