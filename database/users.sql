/* ******** DATABASE USERS ********** */
-- REMEMBER TO SET YOUR OWN PASSWORDS FOR EACH USER
-- Also smart for security reasons to limit who can access each account. (Change '%' to a specific IP)


-- rapidadmin
CREATE USER `rapidadmin`@'%' IDENTIFIED BY 'iloveart';
GRANT ALL PRIVILEGES ON `rapidart`.* TO `rapidadmin`@`%` WITH GRANT OPTION;

-- rapidserver
CREATE USER `rapidserver`@'%' IDENTIFIED BY 'iloveart';
GRANT SELECT, INSERT, UPDATE, DELETE ON `rapidart`.* TO `rapidserver`@`%`;

-- rapidbackup
CREATE USER `rapidbackup`@'%' IDENTIFIED BY 'iloveart';
GRANT SELECT, LOCK TABLES, SHOW VIEW, TRIGGER ON `rapidart`.* TO `rapidbackup`@`%`;