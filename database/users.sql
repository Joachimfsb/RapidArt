/* ******** DATABASE USERS ********** */
-- REMEMBER TO SET YOUR OWN PASSWORDS FOR EACH USER


-- rapidadmin
CREATE USER `rapidadmin` IDENTIFIED BY 'iloveart';
GRANT ALL PRIVILEGES ON `rapidart`.* TO `rapidadmin`@`%` WITH GRANT OPTION;

-- rapidserver
CREATE USER `rapidserver` IDENTIFIED BY 'iloveart';
GRANT SELECT, INSERT, UPDATE, DELETE ON `rapidart`.* TO `rapidserver`@`%`;

-- rapidbackup
CREATE USER `rapidbackup` IDENTIFIED BY 'iloveart';
GRANT SELECT, LOCK TABLES, SHOW VIEW, TRIGGER ON `rapidart`.* TO `rapidbackup`@`%`;