# RapidArt site
This Go project serves as the frontend and backend source code of the RapidArt platform.

# Table of contents

[TOC]

## Deployment
Deployment of this web-service can be done either by running the code directly or using the provided docker-compose script (recommended).

### 1. Configure setup
Before we run the web service it is important to set up the configuration of the web server correctly.

In the directory `/site/configs` you can find a file named `config.json.template`. Open this file in a text editor and fill in your values.

**Explainations:**
```
{
  "server": {
    "host": "localhost", // What ip should the server bind itself to. Use "localhost" if testing locally and "0.0.0.0" if you wish to allow remote connections.
    "port": "8080" // Should be 8080 unless altered in docker-files
  },
  "database": {
    "url": "localhost:3306", // URL and port of database
    "db": "rapidart", // Database name (should be "rapidart" unless changed manually)
    "user": "rapidserver", // Username to connect to the db with ("rapidserver" should be used here most of the time)
    "pass": "iloveart" // Change this to the custom password you set to the "rapidserver" database-user.
  }
}
```

Once your values has been set, rename the file to `config.json`.

### docker-compose
Once the configuration is set up, you can run the build and run docker container for the web-service.

**Prerequisites:**
* You will both need [docker](https://docs.docker.com/get-started/get-docker/) and [docker compose](https://docs.docker.com/compose/install/) installed on your machine.

**Build:**

To build and start the web server on your own server/machine, navigate to this directory in the terminal and run this command: `docker-compose up --build`.

Take a look at the log messages and verify that the app did not return any errors. If you get an error, there might be something wrong with your configuration.

If all went well, you may wish to stop the container by pressing `ctrl + c` and restarting it in detached mode using `docker-compose up --build -d`.

**NOTE:** If you get error messages when starting the docker containers, you may need `sudo` in front of the commands.

### Running directly
To run the web-service directly, follow these steps:

1. Ensure the configuration is set up correctly.
2. Make sure [Go](https://go.dev/dl/) (minumum version 1.22) is installed on your computer/server.
3. Open a terminal and navigate to this directory (`/site`). Then execute the command `go run cmd/rapidart/main.go`.

This should start the server. If you see any error messages, you might have an error in your configuration.

## Web endpoints
* `/top/` - Top posts/users api
  * **BASIC AUTH** GET `/top/posts?{since=time}&{basiscanvas=id}`
    * `since` is optional and represents the top posts *since* a given date
    * `basiscanvas` is optional, and if specified, the given posts are filtered on the basiscavas
  **BASIC AUTH** GET `/top/users?{:metric=string}`
    * Gets the most liked users by a given metric (`likes` or `followers`).

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
* `/api/img/` - Contains images that are fetched from the DB. (Note that some images require authentication)
  * **BASIC AUTH** GET `/api/img/basiscanvas/?id={:BasisCanvasId}` - Fetches a single BasisCanvas by its ID
  * **BASIC AUTH** GET `/api/img/user/profile-pic/?userid={:UserId}` - Fetches a user's profile picture by user ID