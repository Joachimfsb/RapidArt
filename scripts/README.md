# Scripts

## Table of contents

[TOC]

## 1. BasisCanvas Generation Script
This is the script used to create BasisGalleries and BasisCanvases each day. 
It's written in Python, and uses multiple additional libraries.

### Requiremnets to run the script
You must have installed Python on the machine.

After installing Python you have to install all the additional libraries/tools.

#### **GhostScript**
GhostScript is a tool used to help Pillow maintain and convert PostScript to images.

You can install it by this command on a Ubuntu machine: `sudo apt-get install ghostscript`

#### **XVFB**
Turtle uses a graphic display to show the user the creation of the drawings, these displays
aren't available on servers or our Linux-VMs. Without a graphic display the Turtle-code won't
run, therefore we asked ChatGPT for a way around this, and it recommended using XVFB for creating
a virtual display.

You can install it by this command on a Ubuntu machine: `sudo apt-get install xvfb`

### Python Packages and libraries:
#### PythonTurtle
Used to draw the actual lines. This is quite a central package in the script.
#### **Python Packages and libraries:**
**PythonTurtle**

Used to actually draw the lines, this is quite a central package in the script.

You can install it by this command: `pip3 install PythonTurtle`

**Pillow**

Used to manage images (Convert, resize and save, etc.).

You can install it by this command: `pip3 install Pillow`

**SQL-connector**

Used to establish a connection and perform queries to the database.

You can install it by this command: `pip3 install mysql-connector`

### Deployment
Before you set up the script to auto-run, you have to update the database connection configuration directly in the file `basis_canvas.py`. Ensure to use the correct IP-address, database name, username and password for your database.

To get this script to run once every day, it is recommended to set up a cron job (or similar). An example configuration could be:
`0 1 * * * xvfb-run python3 /home/ubuntu/prog2052-prosjekt/scripts/basis_canvas.py`

## 2. Load and/or stress test
This script was used to try loading the webservers with many requests to see how they handle many requests at 
once, and messure the response time of the application.

### Requirements
This test requires locust which is a tool to perform load tests in using Python,
you can install it by using this command: pip install `locust`.

### Usage
To use the application simply type this command in the terminal: `locust -f load_test.py --host=<URL>`. 
Where URL is the URL you want to test. When you've done this it will start a service on port 8089 on localhost, 
open it in your browser and configure the test. 
