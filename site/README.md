# RapidArt site
This Go project serves as the frontend and backend source code of the RapidArt platform.

# Table of contents

[TOC]

## Deployment
To start the web server, open the terminal in this directory and type `go run cmd/rapidart/main.go`. 


## API endpoints
* `/api/user/` - User related APIs
 * **NO AUTH** POST `/api/user/register/`
    * POST attributes:
      * `email`
      * `username`
      * `password`
* `/api/auth/` - Authentication related APIs
  * **NO AUTH** POST `/api/auth/login/`
    * POST attributes:
      * `username`
      * `password`
  * **BASIC AUTH** POST `/api/auth/logout/`
* `/api/img/` - Contains images that are fetched from the DB. (Note that some images require authentication)
  * **BASIC AUTH** GET `/api/img/basiscanvas/?id={:BasisCanvasId}` - Fetches a single BasisCanvas by its ID
  * **BASIC AUTH** GET `/api/img/user/profile-pic/?userid={:UserId}` - Fetches a user's profile picture by user ID