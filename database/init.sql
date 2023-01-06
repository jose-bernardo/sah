DROP DATABASE IF EXISTS testdb;
CREATE DATABASE testdb;

USE testdb;

CREATE TABLE Patients (
    nhs VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,

    PRIMARY KEY (nhs)
);


CREATE TABLE MedicalSpecialty (
    name VARCHAR(255) NOT NULL,

    PRIMARY KEY (name)
);


CREATE TABLE Employee (
    id INT NOT NULL AUTO_INCREMENT,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    medicalSpecialty VARCHAR(255) NOT NULL,

    FOREIGN KEY (medicalSpecialty)
        REFERENCES MedicalSpecialty (name),
    PRIMARY KEY (id)
);


CREATE TABLE Appointments (
    id INT NOT NULL AUTO_INCREMENT,
    `date` DATETIME NOT NULL,
    patientNhs VARCHAR(255) NOT NULL,
    medicalSpecialty VARCHAR(255) NOT NULL,
    room INT,
    doctorId INT,
    state BOOLEAN NOT NULL DEFAULT 0,

    FOREIGN KEY (patientNhs)
        REFERENCES Patients (nhs),
    FOREIGN KEY (doctorId)
        REFERENCES Employee (id),
    FOREIGN KEY (medicalSpecialty)
        REFERENCES MedicalSpecialty (name),

    PRIMARY KEY (id)
);

-- Add medical specialties
INSERT INTO MedicalSpecialty (name) VALUES ("Cardiology");
INSERT INTO MedicalSpecialty (name) VALUES ("Dermatology");
INSERT INTO MedicalSpecialty (name) VALUES ("Orthopedy");
