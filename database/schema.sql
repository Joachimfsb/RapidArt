--
-- Database: rapidart
--
DROP DATABASE IF EXISTS rapidart;
CREATE DATABASE rapidart DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_danish_ci;
USE rapidart;



/* *********** TABLES ************ */
drop table if exists `User`;
drop table if exists `Post`;
drop table if exists `BasisGallery`;
drop table if exists `BasisCanvas`;
drop table if exists `Like`;
drop table if exists `Comment`;
drop table if exists `Report`;
drop table if exists `Follow`;


CREATE TABLE `User` (
    UserId INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    Username VARCHAR(50) NOT NULL,
    Email VARCHAR(255) NOT NULL,
    DisplayName VARCHAR(70) NOT NULL,
    PasswordHash VARCHAR(256) NOT NULL,
    PasswordSalt VARCHAR(16) NOT NULL,
    CreationDateTime DateTime NOT NULL,
    Role ENUM ('user', 'moderator', 'admin') NOT NULL,
    Bio VARCHAR(255),
    ProfilePicture LONGBLOB
);


CREATE TABLE `Post` (
    PostId INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    UserId INT UNSIGNED NOT NULL,
    BasisCanvasId INT UNSIGNED NOT NULL,
    Image BLOB NOT NULL,
    Caption VARCHAR(255),
    TimeSpentDrawing INT UNSIGNED NOT NULL, -- Milliseconds
    CreationDateTime DateTime NOT NULL,
    Active BOOL NOT NULL DEFAULT 1
);

CREATE TABLE `BasisGallery` (
    BasisGalleryId INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    StartDateTime DateTime NOT NULL,
    EndDateTime DateTime NOT NULL
);

CREATE TABLE `BasisCanvas` (
    BasisCanvasId INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    BasisGalleryId INT UNSIGNED NOT NULL,
    Type VARCHAR(50) NOT NULL,
    Image BLOB NOT NULL
);

CREATE TABLE `Like` (
    UserId INT UNSIGNED NOT NULL,
    PostId INT UNSIGNED NOT NULL,
    PRIMARY KEY (UserId, PostId)
);

CREATE TABLE `Comment` (
    UserId INT UNSIGNED NOT NULL,
    PostId INT UNSIGNED NOT NULL,
    Message VARCHAR(512) NOT NULL,
    CreationDateTime DateTime NOT NULL,
    PRIMARY KEY (UserId, PostId)
);

CREATE TABLE `Report` (
    UserId INT UNSIGNED NOT NULL,
    PostId INT UNSIGNED NOT NULL,
    Message VARCHAR(512) NOT NULL,
    CreationDateTime DateTime NOT NULL,
    PRIMARY KEY (UserId, PostId)
);

CREATE TABLE `Follow` (
    FollowerUserId INT UNSIGNED NOT NULL,
    FolloweeUserId INT UNSIGNED NOT NULL,
    PRIMARY KEY (FollowerUserId, FolloweeUserId)
);



/* *********** FOREIGN KEYS ********** */

-- BasisCanvas.BasisGalleryId -> BasisGallery.BasisGalleryId
ALTER TABLE `BasisCanvas`
ADD CONSTRAINT FK_BasisCanvas_BasisGallery
FOREIGN KEY (BasisGalleryId) REFERENCES `BasisGallery`(BasisGalleryId)
ON DELETE RESTRICT;


-- Post.BasisCanvasId -> BasisCanvas.BasisCanvasId
ALTER TABLE `Post`
ADD CONSTRAINT FK_Post_BasisCanvas
FOREIGN KEY (BasisCanvasId) REFERENCES `BasisCanvas`(BasisCanvasId)
ON DELETE RESTRICT;

-- Post.UserId -> User.UserId
ALTER TABLE `Post`
ADD CONSTRAINT FK_Post_User
FOREIGN KEY (UserId) REFERENCES `User`(UserId)
ON DELETE RESTRICT;


-- Like.UserId -> User.UserId
ALTER TABLE `Like`
ADD CONSTRAINT FK_Like_User
FOREIGN KEY (UserId) REFERENCES `User`(UserId)
ON DELETE RESTRICT;

-- Like.PostId -> Post.PostId
ALTER TABLE `Like`
ADD CONSTRAINT FK_Like_Post
FOREIGN KEY (PostId) REFERENCES `Post`(PostId)
ON DELETE RESTRICT;


-- Comment.UserId -> User.UserId
ALTER TABLE `Comment`
ADD CONSTRAINT FK_Comment_User
FOREIGN KEY (UserId) REFERENCES `User`(UserId)
ON DELETE RESTRICT;

-- Comment.PostId -> Post.PostId
ALTER TABLE `Comment`
ADD CONSTRAINT FK_Comment_Post
FOREIGN KEY (PostId) REFERENCES `Post`(PostId)
ON DELETE RESTRICT;


-- Report.UserId -> User.UserId
ALTER TABLE `Report`
ADD CONSTRAINT FK_Report_User
FOREIGN KEY (UserId) REFERENCES `User`(UserId)
ON DELETE RESTRICT;

-- Report.PostId -> Post.PostId
ALTER TABLE `Report`
ADD CONSTRAINT FK_Report_Post
FOREIGN KEY (PostId) REFERENCES `Post`(PostId)
ON DELETE RESTRICT;

-- Follow.FollowerUserId -> User.UserId
ALTER TABLE `Follow`
ADD CONSTRAINT FK_Follow_User1
FOREIGN KEY (FollowerUserId) REFERENCES `User`(UserId)
ON DELETE RESTRICT;

-- Follow.FolloweeUserId -> User.UserId
ALTER TABLE `Follow`
ADD CONSTRAINT FK_Follow_User2
FOREIGN KEY (FolloweeUserId) REFERENCES `User`(UserId)
ON DELETE RESTRICT;



/* ************** INDEXES ************** */
CREATE INDEX I_Post1 ON `Post` (UserId);
CREATE INDEX I_Post2 ON `Post` (BasisCanvasId);
CREATE INDEX I_Post3 ON `Post` (CreationDateTime);
CREATE INDEX I_BasisCanvas ON `BasisCanvas` (BasisGalleryId);
CREATE INDEX I_BasisGallery1 ON `BasisGallery` (StartDateTime);
CREATE INDEX I_BasisGallery2 ON `BasisGallery` (EndDateTime);
CREATE UNIQUE INDEX I_User1 ON `User` (Username);
CREATE UNIQUE INDEX I_User2 ON `User` (Email);



/* ************** VIEWS ************* */
