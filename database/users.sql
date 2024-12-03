/* ******** DATABASE USERS ********** */
-- REMEMBER TO SET YOUR OWN PASSWORDS FOR EACH USER
-- Also smart for security reasons to limit who can access each account. (Change '%' to a specific IP)


-- rapidadmin: Full access to rapidart database
CREATE USER `rapidadmin`@'%' IDENTIFIED BY 'iloveart';
GRANT ALL PRIVILEGES ON `rapidart`.* TO `rapidadmin`@`%` WITH GRANT OPTION;

-- rapidserver: CRUD access to rapidart database
CREATE USER `rapidserver`@'%' IDENTIFIED BY 'iloveart';
GRANT SELECT, INSERT, UPDATE, DELETE ON `rapidart`.* TO `rapidserver`@`%`;

-- rapidscript: Read/Write access to basisgallery and basiscanvas tables in rapidart database
CREATE USER `rapidscript`@'%' IDENTIFIED BY 'iloveart';
GRANT SELECT, INSERT ON `rapidart`.`BasisGallery` TO `rapidscript`@`%`;
GRANT SELECT, INSERT ON `rapidart`.`BasisCanvas` TO `rapidscript`@`%`;

-- rapidbackup: Read permissions to rapidart database
CREATE USER `rapidbackup`@'%' IDENTIFIED BY 'iloveart';
GRANT SELECT, LOCK TABLES, SHOW VIEW, TRIGGER ON `rapidart`.* TO `rapidbackup`@`%`;