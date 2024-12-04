# BasisCanvas Generation Script
This is the script used to create BasisGalleries and BasisCanvases every day. 
It's written in Python, and uses multiple additional libraries.

# Table of contents

[TOC]

## Requiremnets to run the script
You must have Python installed on the machine.

After installing Python you have to install all the additional libraries/tools.

### GhostScript
GhostScript is a tool used to help Pillow maintain and convert PostScript to images.

You can install it by this command on a Ubuntu machine: `sudo apt-get install ghostscript`

### XVFB
Turtle uses a graphic display to show the user the creation of the drawings, these displays
aren't available on servers or our Linux-VMs. Without a graphic display the Turtle-code won't
run, therefore we asked ChatGPT for a way around this, and it recommended using XVFB for creating
a virtual display.

You can install it by this command on a Ubuntu machine: `sudo apt-get install xvfb`

### Python Packages and libraries:
#### PythonTurtle
Used to draw the actual lines. This is quite a central package in the script.

You can install it by this command: `pip3 install PythonTurtle`

#### Pillow
Used to manage images (Convert, resize and save, etc.).

You can install it by this command: `pip3 install Pillow`

#### SQL-connector
Used to establish a connection and perform queries to the database.

You can install it by this command: `pip3 install mysql-connector`
## Deployment
You have to update the password in the Database connection to the correct password for the `rapidserver`, 
before you can run the code. By default this password is `iloveart`.

To get this script to run once every day, it is recommended to set up a cron job (or similar). An example configuration could be: 
`0 1 * * * xvfb-run python3 /home/ubuntu/prog2052-prosjekt/scripts/basis_canvas.py`