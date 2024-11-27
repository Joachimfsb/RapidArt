# RapidArt site
This Go project serves as the frontend and backend source code of the RapidArt platform.

# Table of contents

[TOC]

## Deployment
To start the web server on your own machine, navigate to this directory and use this command: `docker-compose up --build`


## API endpoints
* `/api/user/` - User related APIs
  * **BASIC AUTH** POST `/api/user/follow/{:UserId}/{:Value}`
    * Value is `1` if the user should follow and `0` if the user should stop following.
  * **NO AUTH** POST `/api/user/register/{?check_email_username}`
    * POST attributes:
      * `email`
      * `username`
      * `password`
      * `displayname`
      * `profile_pic`
* `/api/auth/` - Authentication related APIs
  * **NO AUTH** POST `/api/auth/login/`
    * POST attributes:
      * `username`
      * `password`
  * **BASIC AUTH** POST `/api/auth/logout/`
* `/api/top/` - Top posts/users api
  * **BASIC AUTH** GET `/api/top/posts/?{:since=time}&{basiscanvas=id}`
    * `since` is mandatory and represents the top posts *since* a given date
    * `basiscanvas` is optional, and if specified, the given posts are filtered on the basiscavas
  **BASIC AUTH** GET `/api/top/users`
    * Gets the most liked users
* `/api/img/` - Contains images that are fetched from the DB. (Note that some images require authentication)
  * **BASIC AUTH** GET `/api/img/basiscanvas/?id={:BasisCanvasId}` - Fetches a single BasisCanvas by its ID
  * **BASIC AUTH** GET `/api/img/user/profile-pic/?userid={:UserId}` - Fetches a user's profile picture by user ID