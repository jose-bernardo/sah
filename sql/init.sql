DROP DATABASE IF EXISTS testdb;
CREATE DATABASE testdb;

USE testdb;

CREATE TABLE Patients (
    -- id INT NOT NULL AUTO_INCREMENT,
    nhs VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    -- birthday DATE NOT NULL,
    PRIMARY KEY (nhs)
);

CREATE TABLE MedicalSpecialty (
    name VARCHAR(255) NOT NULL,
    PRIMARY KEY (name)
);

CREATE TABLE Appointments (
    id INT NOT NULL AUTO_INCREMENT,
    `date` DATETIME NOT NULL,
    patientNhs VARCHAR(255) NOT NULL,
    medicalSpecialty VARCHAR(255) NOT NULL,
    FOREIGN KEY (patientNhs)
        REFERENCES Patients (nhs),
    FOREIGN KEY (medicalSpecialty)
        REFERENCES MedicalSpecialty (name),
    PRIMARY KEY (id)
);
