# SIRS Project

## CA  setup
Setup CA  
` cd setup `  
` generate-CA-certificate.sh `  

## Front office Server setup 
### Setup databse credentials
` export DB_USER="user" `    
` export DB_PASS="password" `    
### Generate keys  
` cd setup `    
` generate-server-certificates.sh `   
` generate-bd-client-certificates.sh `    
` mv -r bd-client ../sah/certs `    
` mv -r SAH_SERVER ../sah/certs `    

## Internal office Server setup
### Setup databse credentials
` export DB_USER="staff" `   
` export DB_PASS="password" `  
### Generate keys
` cd setup `    
` generate-internal-server-certificates.sh `    
` generate-bd-client-certificates.sh `    
` mv -r bd-client ../SAH-Backoffice/certs `    
` mv -r INTERNAL_SAH_SERVER ../SAH-Backoffice/certs `    

## Database server setup
### Generate keys
` generate-bd-certificates.sh `    
`sudo mv bd-server/* /var/lib/mysql/ `   

### Allow remote connections 
Edit /etc/mysql/mysql.conf.d/mysqld.cnf file   
```
[mysqld]
ssl_ca=ca.pem
ssl_cert=bd-server-cert.pem
ssl_key=bd-server-key.pem
tls_version=TLSv1.3
require_secure_transport=ON
bind-address = 192.168.2.1
```

### User creation
Access mysql server and execute:  
` DROP DATABASE IF EXISTS testdb; `  
` CREATE DATABASE testdb; `  
` CREATE USER ''user''@'192.168.0.100' IDENTIFIED WITH sha256_password BY 'password' REQUIRE X509 WITH MAX_USER_CONNECTIONS 500 PASSWORD EXPIRE DEFAULT; `    
` GRANT SELECT, INSERT, UPDATE ON testdb.* TO 'user'@'192.168.0.100'; `  
` CREATE USER 'staff'@'192.168.1.1' IDENTIFIED WITH sha256_password BY 'password' REQUIRE X509 WITH MAX_USER_CONNECTIONS 500 PASSWORD EXPIRE DEFAULT; `  
` GRANT SELECT, INSERT, UPDATE ON testdb.* TO 'staff'@'192.168.1.1'; `    

## Internal Doctor machine  
` generate-doctor-certificate.sh 1 `  




