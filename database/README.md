# RapidArt Database
This is the database of the RapidArt webside. We've decided to use MariaDB which is a SQL database. We've 
also made use of phpMyAdmin to help monitor and administrate over the database.

# Table of contents

[TOC]

## Deployment
Before you deploy the database, it is highly recommended you update both a MariaDB Root password and each individual passwords for the
users in `users.sql`. You can change the MariaDB Root password by changing `MARIADB_ROOT_PASSWORD` in the compose-file. All passwords will
by default be `iloveart`.

To start the database using docker you have to navigate to this directory and use this command: `docker-compose up -d`

When you've done this you can navigate to phpMyAdmin in your browser, which is located on `<Machine floating IP>:8080`. On phpMyAdmin
create a database called "RapidArt", and upload the SQL-files, which will now be accessible on `<Machine IP>:3306`