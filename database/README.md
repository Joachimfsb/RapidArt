# RapidArt Database
This README provides information on the RapidArt Database implementation and deployment.


## Implementation 💾
The implementation is created with the MySQL protocol and tested with MariaDB.

All needed tables are included in the central database named `rapidart`. The schema with all required tables and attributes can be found in the file `schema.sql`.

The project comes with a set of users that each have their purpose. These can be found in the `users.sql` file and are required for everything to work correctly.

## Deployment 🚀
There are two ways of deploying the database. Either using the provided docker compose script, or by manually installing and importing the provided `.sql` files.

**General prerequisites:**
* Before you deploy the database, regardless of method, it is highly recommended you update both a MariaDB Root password and each individual passwords for each users in the `users.sql` file. Note down these passwords.
  * **USING THE DEFAULT PASSWORDS LEAVES YOU AND YOUR USERS SUSCEPTIBLE TO DATA INTRUSION AND DATA LOSS** 

### docker-compose
The docker compose script automatically sets up two containers, MariaDB and PHPMyAdmin. 

**Prerequisites:**
* You will both need [docker](https://docs.docker.com/get-started/get-docker/) and [docker compose](https://docs.docker.com/compose/install/) installed on your machine.

**Steps:**
1. Create your own MariaDB Root password in the `compose.yml`. Do this by changing the value after `MARIADB_ROOT_PASSWORD=`. **It is important to change this password to protect against attacks.**
2. To start the database using docker you have to navigate to this directory in the terminal. Once completed, execute this command: `docker-compose up -d`. This will start the database.
3. Once you've done this you can navigate to phpMyAdmin in your browser, which is located on `<Databaseserver Machine public IP>:8080`.
4. Log in using your MariaDB root account.
5. Go to the import section and import the `schema.sql` file.
6. Import the `users.sql` file **(REMEMBER TO SET YOUR OWN PASSWORDS)**.
7. Your database should be ready for production/testing.

**NOTE:** If you get error messages when starting the docker containers, you may need `sudo` in front of the commands.

### Manual install

**Prerequisites:**
* You will both need to install [MariaDB](https://mariadb.org/download/) and optionally [PHPMyAdmin](https://www.phpmyadmin.net/downloads/) (if you wish a GUI).

**Steps:**
1. Log in to MariaDB using either the terminal or the PHPMyAdmin gui.
2. (GUI: Go to the import section and) import the `schema.sql` file.
3. Import the `users.sql` file **(REMEMBER TO SET YOUR OWN PASSWORDS)**.
4. Your database should be ready for production/testing.
