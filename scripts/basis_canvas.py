import turtle
import random
import math
from PIL import Image
from datetime import datetime, timedelta

# All code related to databases is commented out:


import mysql.connector

# Database setup
# Got help from this video: https://www.youtube.com/watch?v=3vsC05rxZ8c
db = mysql.connector.connect(
    host="localhost",
    user="rapidserver",
    database="rapidart",
    passwd="iloveart"
)

mycursor = db.cursor()

# Screen setup
width=700
height=600
screen = turtle.Screen()
screen.setup(width, height)
screen_width = screen.window_width() // 2
screen_height = screen.window_height() // 2

# Turtle setup
t = turtle.Turtle()
t.speed("fastest")
t.hideturtle()

# Generate BasisCanvas lines:
max_turn = 140
def basis_canvas_line(turns):
    # Calculate a random starting point on the screen:
    x = random.randint(-screen_width, screen_width)
    y = random.randint(-screen_height, screen_height)

    # Go to the starting point
    t.penup()
    t.goto(x, y)
    t.pendown()

    # Loop through and turn the correct amount of times
    while turns > 0:
        # Randomize the angle and length of the line 
        angle = random.randint(-max_turn, max_turn)
        length = random.randint(100, 500)

        heading = math.radians(t.heading() + angle)
        
        # Calculate the target position
        new_x = t.xcor() + length * math.cos(heading)
        new_y = t.ycor() + length * math.sin(heading)

        # Check if the new position is inside of the screen boundaries:
        if -screen_width < new_x < screen_width and -screen_height < new_y < screen_height:
            # Turn and move if within bounds
            t.left(angle)
            t.forward(length)

            # Deincrement the turns
            turns -= 1

# Generate BasisCanvas shape:
max_turn=180
def basis_canvas_shape():
    # Give the shape a random starting rotation
    angle = random.randint(-max_turn, max_turn)
    t.left(angle)

    # Choose a random shape
    shape = random.choice([circle, square, triangle, hexagon])
    
    # Decide a random size
    size = random.randint(50, 250)

    # Calculate a random starting point on the screen:
    min_x = (- screen.window_width() + size) // 2
    max_x = (screen.window_width() - size) // 2
    min_y = (- screen.window_height() + size) // 2
    max_y = (screen.window_height() - size) // 2

    x = random.randint(min_x, max_x)
    y = random.randint(min_y, max_y)

    # Go to the starting point
    t.penup()
    t.goto(x, y)
    t.pendown()

    # Draw the shape
    shape(size)

# Draw a circle
def circle(radius):
    t.circle(radius)

# Draw a square
def square(side_length):
    for _ in range(4):
        t.forward(side_length)
        t.left(90)

# Draw a triangle
def triangle(side_length):
    for _ in range(3):
        t.forward(side_length)
        t.left(120)

# Draw a hexagon
def hexagon(side_length):
    for _ in range(6):
        t.forward(side_length)
        t.left(300)

# Save the image as a PNG with transparent background
def save_image(num):
    # Save the canvas as an image using Ghostscript
    screen.getcanvas().postscript(file="basis_canvas_" + str(num) + ".ps")

    # Convert to PNG using Pillow
    img = Image.open("basis_canvas_" + str(num) + ".ps")
    img.save("basis_canvas_" + str(num) + ".png")

    # Make background transparent
    make_transparent("basis_canvas_" + str(num) + ".png")

    # Resize the image:
    img = Image.open("basis_canvas_" + str(num) + ".png")
    img = img.resize((width, height))
    img.save("basis_canvas_" + str(num) + ".png")

# ChatGPT helped with this function, which makes the background transparent:
def make_transparent(file):
    image = Image.open(file).convert("RGBA")

    # Get the data of the image
    data = image.getdata()

    # Define a new empty list to store the modified pixels
    new_data = []

    # Loop through each pixel
    for item in data:
        # Change white to transparent:
        if item[:3] == (255, 255, 255):  # If white:
            new_data.append((255, 255, 255, 0))  # Make transparent
        else:
            new_data.append(item)  # Keep all non-white pixel

    # Update the image, and save it
    image.putdata(new_data)
    image.save(file, "PNG")

# Draw 5 random lines and save it
for i in range(1, 6):
    t.clear()

    turns = random.randint(1, 5)

    # Draw a line:
    t.color(random.random(), random.random(), random.random())
    t.pensize(random.randint(2, 5))
    basis_canvas_line(turns)

    save_image(i)

# Draw 5 random shapes and save it
for i in range(6, 11):
    t.clear()

    # Draw a shape:
    t.color(random.random(), random.random(), random.random())
    t.pensize(random.randint(2, 5))
    basis_canvas_shape()

    save_image(i)



# Insert the BasisCanvases into the database
start_time = datetime.now().replace(hour=6, minute=0, second=0, microsecond=0)
end_time = start_time + timedelta(hours=24) # 24 hours later

# Convert times to strings
start_time_str = start_time.strftime('%Y-%m-%d %H:%M:%S')
end_time_str = end_time.strftime('%Y-%m-%d %H:%M:%S')


# Make a BasisGallery
mycursor.execute("INSERT INTO `BasisGallery` (StartDateTime, EndDateTime) VALUES (%s, %s)", (start_time_str, end_time_str))
db.commit()

# Get the generated BasisGalleryId
mycursor.execute("SELECT BasisGalleryId FROM `BasisGallery` WHERE StartDateTime = %s AND EndDateTime = %s", (start_time_str, end_time_str))
result = mycursor.fetchone()

# Check if the result actually exists (It should)
if result:
    gallery_id = result[0]
else:
    print("BasisGallery not found.")

# Insert all generated BasisCanvases:
for i in range(1, 11):
    image_path = "basis_canvas_" + str(i) + ".png"
    # I had help with reading the image as a BLOB: https://www.youtube.com/watch?v=NwvTh-gkdfs
    with open(image_path, "rb") as file:
        image_data = file.read()  # Read the binary image data

    type = "Line"
    if i >= 6:
        type = "Shape"
    
    mycursor.execute("INSERT INTO `BasisCanvas` (BasisGalleryId, Type, Image) VALUES (%s, %s, %s)", (gallery_id, type, image_data))
    db.commit()
    
# Close the database connection:
mycursor.close()
db.close()
