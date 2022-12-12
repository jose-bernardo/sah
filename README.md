# SIRS Project

## Database setup
### Create database keys and certificate
` cd setup `

Create CA certificate  
` openssl genrsa 2048 > ca-key.pem `  
` openssl req -new -x509 -nodes -days 3600 -key ca-key.pem -out ca.pem` 

Create server certificate, remove passphrase, and sign it.  
server-cert.pem = public key, server-key.pem = private key  
` openssl req -newkey rsa:2048 -days 3600 -nodes -keyout server-key.pem -out server-req.pem `  
` openssl rsa -in server-key.pem -out server-key.pem `  
` openssl x509 -req -in server-req.pem -days 3600 -CA ca.pem -CAkey ca-key.pem -set_serial 01 -out server-cert.pem -extfile config.cnf -extensions v3_ca `  

Create client certificate, remove passphrase, and sign it.  
client-cert.pem = public key, client-key.pem = private key  
` openssl req -newkey rsa:2048 -days 3600 -nodes -keyout client-key.pem -out client-req.pem `  
` openssl rsa -in client-key.pem -out client-key.pem `  
` openssl x509 -req -in client-req.pem -days 3600 -CA ca.pem -CAkey ca-key.pem -set_serial 01 -out client-cert.pem `  

Move certificates and key to data folder  
` mv ca.pem server-cert.pem server-key.pem /var/lib/mysql/ `  

Distribute ca.pem, client-cert.pem and client-key.pem to clients  

### Allow remote connections
Edit /etc/mysql/mysql.conf.d/mysqld.cnf file
```
[mysqld]
ssl_ca=ca.pem
ssl_cert=server-cert.pem
ssl_key=server-key.pem
tls_version=TLSv1.3
require_secure_transport=ON
bind-address = 192.168.2.1
```

### User creation
Access mysql server and execute:  
` CREATE USER 'frontOffice'@'192.168.0.100' IDENTIFIED WITH sha256_password BY 'password' REQUIRE X509 WITH MAX_USER_CONNECTIONS 3 PASSWORD EXPIRE DEFAULT; `
