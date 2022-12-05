DROP DATABASE IF EXISTS testdb;
CREATE DATABASE testdb;

USE testdb;

CREATE TABLE Patients (
    nhs DECIMAL(9) NOT NULL,
    name VARCHAR(255) NOT NULL,
    birthday DATE NOT NULL,
    PRIMARY KEY (nhs)
);

CREATE TABLE MedicalSpecialty (
    name VARCHAR(255) NOT NULL,
    PRIMARY KEY (name)
);

CREATE TABLE Appointments (
    id INT NOT NULL AUTO_INCREMENT,
    `date` DATETIME NOT NULL,
    patientNhs DECIMAL(9) NOT NULL,
    medicalSpecialty VARCHAR(255) NOT NULL,
    FOREIGN KEY (patientNhs)
        REFERENCES Patients (nhs),
    FOREIGN KEY (medicalSpecialty)
        REFERENCES MedicalSpecialty (name),
    PRIMARY KEY (id)
);
