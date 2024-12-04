# RapidArt site
This Go project serves as the frontend and backend source code of the RapidArt platform.

## Table of contents

[TOC]

## Deployment 🚀
Deployment of this web-service can be done either by running the code directly or using the provided docker-compose script (recommended).

### Step 1. Configure setup
Before we run the web service it is important to set up the configuration of the web server correctly.

In the directory `/site/configs` you can find a file named `config.json.template`. Open this file in a text editor and fill in your values.

**Explainations:**
```
{
  "server": {
    "host": "localhost",     // What ip should the server bind itself to. Use "localhost" if testing locally and "0.0.0.0" if you wish to allow remote connections.
    "port": "8080"           // Should be 8080 unless altered in docker-files
  },
  "database": {
    "url": "localhost:3306", // URL and port of database
    "db": "rapidart",        // Database name (should be "rapidart" unless changed manually)
    "user": "rapidserver",   // Username to connect to the db with ("rapidserver" should be used here most of the time)
    "pass": "iloveart"       // Change this to the custom password you set to the "rapidserver" database-user.
  }
}
```

Once your values has been set, rename the file to `config.json`.

Your configuration should thus be ready. Continue to the **Step 2. docker-compose** section or further down to **Step 2. Running directly**.

### Step 2. docker-compose (option 1)
Once the configuration is set up, you can run the build and run docker container for the web-service.

**Prerequisites:**
* You will both need [docker](https://docs.docker.com/get-started/get-docker/) and [docker compose](https://docs.docker.com/compose/install/) installed on your machine.

**Build:**

To build and start the web server on your own server/machine, navigate to this directory in the terminal and run this command: `docker-compose up --build`.

Take a look at the log messages and verify that the app did not return any errors. If you get an error, there might be something wrong with your configuration.

If all went well, you may wish to stop the container by pressing `ctrl + c` and restarting it in detached mode using `docker-compose up --build -d`.

**NOTE:** If you get error messages when starting the docker containers, you may need `sudo` in front of the commands.

### Step 2. Running directly (option 2)
To run the web-service directly, follow these steps:

1. Ensure the configuration is set up correctly.
2. Make sure [Go](https://go.dev/dl/) (minumum version 1.22) is installed on your computer/server.
3. Open a terminal and navigate to this directory (`/site`). Then execute the command `go run cmd/rapidart/main.go`.

This should start the server. If you see any error messages, you might have an error in your configuration.

## Implementation 🖥️

### Services
The web-application contains only one service (executable) located in `cmd/rapidart/`.

### Architecture
The web-application is structured in a 4+1 layered architecture:
1. Presentation (Frontend) - HTML/CSS/JS (`web/`)
2. Interface (Backend) - Go - handlers (`internal/handlers/`)
3. Domain (Backend) - Go - various packages (`internal/*`)
4. Data connection (Backend) - Go - database package (`internal/database/`)
5. **(External)** Database (Database) - MySQL/MariaDB


### File / package structure
The code is structured in the following way:
* `cmd/` - Executables
* `configs/` - Configuration of the server
* `internal/` - Internal components of the application
  * `auth/` - Authentication management
  * `basismanager/` - Manages basiscanvases and basisgalleries
  * `config/` - Reads and stores the server configuration
  * `consts/` - Application wide constants
  * `crypto/` - Provides cryptographic related functions
  * `database/` - Database interface. Other components uses this package to send CRUD operations to the database
  * `handlers/` - Manages the web-servers endpoints
    * `web/` - Web related endpoints: Gathers information and generates html files from templates. 
    * `api/` - API interface endpoints: Responds and takes action on user initiated requests.
    * `middleware/` - Functions that are run before handlers. Mostly used to ensure user is authenticated.
  * `models/` - Data models
  * `post/` - Manages posts including likes, comments and reports.
  * `user/` - Manages everything to do with a user, ex. login, profile mgmt., follows, etc.
  * `util/` - Various utility (helper) functions. 
* `test/` - Test related helpers
* `web/` - Templated and static web-resources

## Endpoints 🔌
The following endpoints are made available by the server. 

The endpoints are defined in the file `internal/handlers/routes.go`.

Some endpoints are marked with **AUTH** meaning that they require the user to be authenticated (redirects to `/login` if not authenticated), while some are marked **NO AUTH** meaning they require the user to NOT be authenticated (redirects to `/` if they are authenticated).  

### Web endpoints
* **NO AUTH** GET `/login/` - Login page
* **NO AUTH** GET `/register/` - Register page
* **AUTH** GET `/` - Front page
* **AUTH** GET `/profile/{username}` - Profile page of self or another user.
* **AUTH** GET `/drawing/` - Drawing page
* **AUTH** GET `/post/{post_id}` - Shows a single post
* **AUTH** GET `/search/` - Search for user page
* **AUTH** GET `/toplist/` - Toplist page
* **Web components:**
  * **AUTH** GET `/comments/{post_id}` - Returns a list of comments
  * `/top/{type}` - Top posts/users component
    * **AUTH** GET `/top/posts?{since=time}&{basiscanvas=id}`
      * `since` is optional and represents the top posts *since* a given date
      * `basiscanvas` is optional, and if specified, the given posts are filtered on the basiscavas
    * **AUTH** GET `/top/users?{:metric=string}`
      * Gets the most liked users by a given metric (`likes` or `followers`).

### API endpoints
* `/api/post/` - Post related APIs
  * **AUTH** POST `/api/post/comment/{:id}` - Posts a new comment to the given post
  * **AUTH** POST `/api/post/like/{:id}` - Likes the given post
  * **AUTH** POST `/api/post/unlike/{:id}` - Unlikes the given post
  * **AUTH** POST `/api/post/report/{:id}` - Report a post
* `/api/user/` - User related APIs
  * **AUTH** POST `/api/user/follow/{:UserId}/{:Value}`
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
  * **AUTH** POST `/api/auth/logout/`
* `/api/img/` - Contains images that are fetched from the DB. (Note that some images require authentication)
  * **AUTH** GET `/api/img/basiscanvas/?id={:BasisCanvasId}` - Fetches a single BasisCanvas by its ID
  * **AUTH** GET `/api/img/user/profile-pic/?userid={:UserId}` - Fetches a user's profile picture by user ID
  * **AUTH** GET `/api/img/post/?{:post_id=int}` - Fetches the post image of a given post
* `/api/search/`
  * **AUTH** POST `/api/search/users/` - Searches for a user and returns the best matches.

## Testing 🧪
To test the code, use the built in functionality in Go.

Enter this directory in a terminal and run `go test ./...`.