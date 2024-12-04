# RapidArt platform ✏️

Welcome to the repository for the RapidArt platform. This repository contains all the source code and information needed to deploy the RapidArt platform.

For project details, see the **wiki**.

This README contains an overview of the platform as a whole, as well as deployment instructions.

For implementation details, see the *README.md* files in each relevant sub-directory.

## What is RapidArt? 🔍
RapidArt is a social-media type platform where users are challenged to, in a limited amount of time, create drawings based on a set of basis-canvases. The drawings can then be shared with other users who can react and respond to them.

## Platform overview 📑
The RapidArt platform consists of two primary components, the web-service (written in Go) and the database (MySQL based). In addition, a third component (Python script) is required that runs as a daily worker job, which creates the basis-canvases that are renewed each day.

### Web service
The web service is written in Go and can be found in its entirety in the `/site` directory.

It can optionally be run either as a docker container or as a standalone executable. See the corresponding README.md for more details.

### Database
The platform is designed to be used with a MySQL based database. This means you can opt to use which-ever database implementation you wish, as long as it implements the MySQL protocol. Primarily, **MariaDB is recommended** as this is the implementation that has been tested.

The database schema and its required users can be found in the `/database` directory.

### Worker script
A worker script (Python) that is required to run once-a-day can be found in the `/scripts` directory. This is, as previously stated, a script that generates the daily basis-canvases that renew each day.


## Deployment 🚀
Deployment of this application can be done in multiple ways.

The quickest solution is to simply use the provided **docker-compose** scripts. However, it is also possible to **manually install a database** and **run the Go web-server directly**.

### Step 1. Database
Begin by installing and setting up the database on your server/workstation.

Go to the **Deployment** section in the [README file](./database/) found in `/database` and follow the steps.

Once completed continue here.

### Step 2. Web service
Once the database is ready and set up with your personal passwords, you can begin setting up the web-service.

Go to the **Deployment** section in the [README file](./site/) found in `/site` and follow the steps.

Once completed continue here.

### Step 3. Worker script
With both the database and web service set up, the only step remaining is setting up the worker script that creates a set of basis-canvases each day.

Please read the [README file](./scripts/) found in `/scripts`, install all prerequisites and follow the steps shown under the Deployment section.

Once completed, everything should be set up and working.

Have fun! ⭐