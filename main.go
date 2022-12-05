package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  //"html/template"
  "os"
  "fmt"
)

const DB_NAME = "testdb"
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"

type Appointment struct {
    Nhs, MedicalSpecialty, Date string
}

func main() {

    DB_USER := os.Getenv("DB_USER")
    DB_PASS := os.Getenv("DB_PASS")

    db, err := sql.Open("mysql", DB_USER + ":" + DB_PASS + "@/" + DB_NAME)
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    r := gin.Default()
    r.LoadHTMLGlob("templates/*.html")

    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", gin.H{})
    })
    
    r.GET("/api/new_appointment", func(c *gin.Context) {
        c.HTML(http.StatusOK, "new_appointment.html", gin.H{})
    })

    r.POST("/api/new_appointment", func(c *gin.Context) {
        name := c.PostForm("name")
        nhs := c.PostForm("nhs")
        date := c.PostForm("date")
        medicalSpecialty := c.PostForm("medicalSpecialty")

        fmt.Printf("name: %s, nhs: %s\n", name, nhs); 

        tx, err := db.Prepare("INSERT INTO Appointments (date, patientNhs, medicalSpecialty) VALUES ( ?, ?, ? );")
        if err != nil {
            panic(err.Error())
        }
        defer tx.Close()

        _, err = tx.Exec(date, nhs, medicalSpecialty)
        if err != nil {
            panic(err.Error())
        }
    })

    r.GET("/api/appointments/:nhs", func(c *gin.Context) {

        nhs := c.Param("nhs")

        fmt.Printf("getting appointments for: %s\n", nhs);

        tx, err := db.Prepare("SELECT patientNhs, medicalSpecialty, date FROM Appointments WHERE patientNhs = ?;")
        if err != nil {
            panic(err.Error())
        }
        defer tx.Close()

        rows, err := tx.Query(nhs)
        if err != nil {
            panic(err.Error())
        }

        var appointments []Appointment
        for rows.Next() {
            var appoint Appointment

            err := rows.Scan(&appoint.Nhs, &appoint.MedicalSpecialty, &appoint.Date)
            if err != nil {
                panic(err.Error())
            }

            fmt.Println(appoint);

            appointments = append(appointments, appoint)
        }

        if rows.Err() != nil {
            panic(err.Error())
        }
        
        c.HTML(http.StatusOK, "appointments.html", gin.H{"Appointments": appointments})
    })


    r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
