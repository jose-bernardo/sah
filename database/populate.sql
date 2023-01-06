-- Add medical specialties
INSERT INTO MedicalSpecialty (name) VALUES ("Cardiology");
INSERT INTO MedicalSpecialty (name) VALUES ("Dermatology");
INSERT INTO MedicalSpecialty (name) VALUES ("Orthopedy");

INSERT INTO Patients (nhs, email, name, password) VALUES ("123456789", "pexas@yahoo.copm", "José","password");
INSERT INTO Patients (nhs, email, name, password) VALUES ("987654321", "NNate@sah.com.pt", "Nuno","password2");
INSERT INTO Patients (nhs, email, name, password) VALUES ("123987456", "gas@xd.pt", "Gonçalo","password3");

INSERT INTO Employee (id, password, name, medicalSpecialty) VALUES (1, "ola", "Maria", "Cardiology");
INSERT INTO Employee (id, password, name, medicalSpecialty) VALUES (2, "adeus" ,"Carolina","Orthopedy");

INSERT INTO Appointments (id, `date`, patientNhs, medicalSpecialty, room, doctorId) VALUES (1, "2023-02-24 10:00:00" ,"123987456", "Orthopedy", 26,2);
