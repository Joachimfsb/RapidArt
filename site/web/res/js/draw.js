const canvas = document.getElementById("canvas");
const fixedWidth = 700;
const fixedHeight = 600;

let context = canvas.getContext("2d");
let draw_color = "black";
let draw_width = parseInt(document.getElementById("brush-size-input").value);
let is_drawing = false;

// Timer variables
let timerDuration = 5 * 60; // 5 minutes in seconds
let timerInterval = null;
let timerStarted = false;

let restore_array = [];
let index = -1;

let previous_color = draw_color;
let fillMode = false;  // Fill mode toggle

// Eraser functionality
document.getElementById("eraser-icon").addEventListener("click", () => {
    previous_color = draw_color;  // Save the current color
    draw_color = "white";  // Set eraser color
    previewCircle.style.backgroundColor = draw_color;  // Update preview
});

// Pencil functionality
document.getElementById("pencil-icon").addEventListener("click", () => {
    draw_color = previous_color;  // Restore previous color
    previewCircle.style.backgroundColor = draw_color;  // Update preview
});

canvas.addEventListener("click", (event) => {
    if (fillMode) {
        context.fillStyle = draw_color;  // Set fill color to current drawing color
        context.fillRect(0, 0, canvas.width, canvas.height);  // Fill entire canvas
        restore_array.push(context.getImageData(0, 0, canvas.width, canvas.height));  // Save state for undo
        index += 1;
        fillMode = false;  // Turn off fill mode after filling
    }
});

document.getElementById("fill-icon").addEventListener("click", () => {
    fillMode = true;
});

// Resizing function with scale adjustments for accuracy
function resizeCanvas() {
    const container = document.querySelector('.canvas-container');
    const displayWidth = container.clientWidth;
    const aspectRatio = fixedWidth / fixedHeight;
    const displayHeight = displayWidth / aspectRatio;

    canvas.width = fixedWidth;
    canvas.height = fixedHeight;
    canvas.style.width = `${displayWidth}px`;
    canvas.style.height = `${displayHeight}px`;

    if (restore_array.length > 0) {
        context.putImageData(restore_array[index], 0, 0);
    } else {
        initializeCanvasBackground();
    }
}

// Initial canvas background
function initializeCanvasBackground() {
    context.fillStyle = "white";
    context.fillRect(0, 0, fixedWidth, fixedHeight);
}

// Calculate the scale factor
function getScaleFactor() {
    const displayWidth = parseInt(window.getComputedStyle(canvas).width);
    const displayHeight = parseInt(window.getComputedStyle(canvas).height);
    return {
        x: canvas.width / displayWidth,
        y: canvas.height / displayHeight
    };
}

// Drawing functions
function start(event) {
    if (!timerStarted) {
        startTimer();
        timerStarted = true;
    }
    is_drawing = true;
    context.beginPath();

    const scale = getScaleFactor();
    const x = (event.offsetX || event.touches[0].clientX - canvas.getBoundingClientRect().left) * scale.x;
    const y = (event.offsetY || event.touches[0].clientY - canvas.getBoundingClientRect().top) * scale.y;
    context.moveTo(x, y);

    clickDraw(x, y);
    event.preventDefault();
}

function clickDraw(x, y) {
    context.lineTo(x, y);
    context.strokeStyle = draw_color;
    context.lineWidth = draw_width;
    context.lineCap = "round";
    context.lineJoin = "round";
    context.stroke();
}

function draw(event) {
    if (is_drawing) {
        const scale = getScaleFactor();
        const x = (event.offsetX || event.touches[0].clientX - canvas.getBoundingClientRect().left) * scale.x;
        const y = (event.offsetY || event.touches[0].clientY - canvas.getBoundingClientRect().top) * scale.y;

        context.lineTo(x, y);
        context.strokeStyle = draw_color;
        context.lineWidth = draw_width;
        context.lineCap = "round";
        context.lineJoin = "round";
        context.stroke();
    }
    event.preventDefault();
}

function stop(event) {
    if (is_drawing) {
        context.stroke();
        context.closePath();
        is_drawing = false;
    }
    event.preventDefault();

    if (event.type !== 'mouseout') {
        restore_array.push(context.getImageData(0, 0, canvas.width, canvas.height));
        index += 1;
    }
}

// Tool button actions
function undo_last() {
    if (index <= 0) {
        clear_canvas();
    } else {
        index -= 1;
        restore_array.pop();
        context.putImageData(restore_array[index], 0, 0);
    }
}

function clear_canvas() {
    context.fillStyle = "white";
    context.clearRect(0, 0, canvas.width, canvas.height);
    context.fillRect(0, 0, canvas.width, canvas.height);
    restore_array = [];
    index = -1;

}

// Function to format the timer display
function formatTime(seconds) {
    const minutes = Math.floor(seconds / 60);
    const remainingSeconds = seconds % 60;
    return `${minutes.toString().padStart(2, '0')}:${remainingSeconds.toString().padStart(2, '0')}`;
}

// Start the countdown timer
function startTimer() {
    const timerDisplay = document.getElementById("timer");
    timerDisplay.textContent = formatTime(timerDuration);
    startTime = Date.now();

    timerInterval = setInterval(() => {
        timerDuration--;
        timerDisplay.textContent = formatTime(timerDuration);

        if (timerDuration <= 0) {
            clearInterval(timerInterval);
            save_to_database();
        }
    }, 1000);
}

/*
// function to save locally
function save_as_png() {
    const basisImage = document.getElementById('basis');
    const drawingCanvas = document.getElementById('canvas');

    if (!basisImage.complete) {
        basisImage.onload = save_as_png;
        return;
    }

    const tempCanvas = document.createElement('canvas');
    tempCanvas.width = fixedWidth;
    tempCanvas.height = fixedHeight;
    const tempContext = tempCanvas.getContext('2d');

    tempContext.drawImage(drawingCanvas, 0, 0, fixedWidth, fixedHeight);

    const basisAspectRatio = basisImage.naturalWidth / basisImage.naturalHeight;
    let basisWidth, basisHeight;

    if (fixedWidth / fixedHeight > basisAspectRatio) {
        basisHeight = fixedHeight;
        basisWidth = basisHeight * basisAspectRatio;
    } else {
        basisWidth = fixedWidth;
        basisHeight = basisWidth / basisAspectRatio;
    }

    const basisX = (fixedWidth - basisWidth) / 2;
    const basisY = (fixedHeight - basisHeight) / 2;

    tempContext.drawImage(basisImage, basisX, basisY, basisWidth, basisHeight);

    const dataUrl = tempCanvas.toDataURL("image/png");
    const link = document.createElement('a');
    link.href = dataUrl;
    link.download = 'drawing.png';
    link.click();
}*/

// Function for saving to database
function save_to_database() {
    disableExitWarning()
    const basisImage = document.getElementById('basis');
    const drawingCanvas = document.getElementById('canvas');

    if (!basisImage.complete) {
        basisImage.onload = save_to_database;
        return;
    }

    const tempCanvas = document.createElement('canvas');
    tempCanvas.width = fixedWidth;
    tempCanvas.height = fixedHeight;
    const tempContext = tempCanvas.getContext('2d');

    tempContext.drawImage(drawingCanvas, 0, 0, fixedWidth, fixedHeight);

    const basisAspectRatio = basisImage.naturalWidth / basisImage.naturalHeight;
    let basisWidth, basisHeight;

    if (fixedWidth / fixedHeight > basisAspectRatio) {
        basisHeight = fixedHeight;
        basisWidth = basisHeight * basisAspectRatio;
    } else {
        basisWidth = fixedWidth;
        basisHeight = basisWidth / basisAspectRatio;
    }

    const basisX = (fixedWidth - basisWidth) / 2;
    const basisY = (fixedHeight - basisHeight) / 2;

    tempContext.drawImage(basisImage, basisX, basisY, basisWidth, basisHeight);

    const mergedDataUrl = tempCanvas.toDataURL("image/png");
    const mergedImageData = mergedDataUrl.replace(/^data:image\/png;base64,/, "");

    const basisCanvasId = parseInt(basisImage.dataset.basisCanvasId, 10);
    if (!basisCanvasId || basisCanvasId === 0) {
        alert("Ugyldig BasisCanvasId");
        return;
    }

    const timeSpentDrawing = Date.now() - startTime;

    const postData = {
        image_data: mergedImageData,
        basis_canvas_id: basisCanvasId,
        caption: '',
        time_spent_drawing: timeSpentDrawing,
    };

    const xhr = new XMLHttpRequest();
    xhr.open("POST", "/api/save-post", true);
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.onreadystatechange = function() {
        if (xhr.readyState === 4) {
            // STATUS OK
            if (xhr.status === 200) {
                // Get returned id
                id = parseInt(xhr.responseText);
                // Check that returned id is a number
                if (!isNaN(id)) {
                    // Redirect to created post
                    window.location = "/post/?post_id=" + id;
                } else {
                    // Not a number
                    alert("Something went wrong, could not save post!");
                }
            } else {
                alert("Something went wrong, could not save post!");
            }
        }
    };
    xhr.send(JSON.stringify(postData));
}

// Change color from color buttons
function change_color(color) {
    draw_color = color;
    previewCircle.style.backgroundColor = draw_color;
}

// Create Pickr instance
const pickr = Pickr.create({
    el: '#color-picker-icon',
    theme: 'classic',
    default: '#000000',
    useAsButton: true,
    position: 'top-middle',
    swatches: ['#FF0000', '#00FF00', '#0000FF', '#FFFF00', '#FF00FF', '#00FFFF'],
    components: {
        preview: true,
        opacity: false,
        hue: true,
        interaction: {
            hex: false,
            rgba: false,
            hsla: false,
            hsva: false,
            cmyk: false,
            input: false,
            clear: false,
            save: false
        }
    }
});

// Show color picker when the icon is clicked
document.getElementById("color-picker-icon").addEventListener("click", () => {
    pickr.show();
});

// Update the drawing color and brush preview color when a color is selected
pickr.on('change', (color) => {
    draw_color = color.toHEXA().toString();
    previewCircle.style.backgroundColor = draw_color;
});

// Update brush size and preview circle size
function updateBrushSize(size) {
    draw_width = parseInt(size) || 1;
    previewCircle.style.width = `${draw_width}px`;
    previewCircle.style.height = `${draw_width}px`;
}

const brushSizeInput = document.getElementById("brush-size-input");
brushSizeInput.addEventListener("input", (event) => updateBrushSize(event.target.value));

brushSizeInput.addEventListener("keydown", function (event) {
    if (event.key === "Enter") {
        event.preventDefault();
        brushSizeInput.blur(); // Removes focus from the input field
    }
});

canvas.addEventListener("mousedown", () => {
    brushSizeInput.blur();
});
canvas.addEventListener("touchstart", () => {
    brushSizeInput.blur();
});

// Event listeners for resizing and drawing
window.addEventListener('load', () => {
    resizeCanvas();
    initializeCanvasBackground();
});
window.addEventListener('resize', resizeCanvas);

canvas.addEventListener("mousedown", start);
canvas.addEventListener("mousemove", draw);
canvas.addEventListener("mouseup", stop);
canvas.addEventListener("touchstart", start);
canvas.addEventListener("touchmove", draw);
canvas.addEventListener("touchend", stop);
canvas.addEventListener("mouseout", stop);

// Tool button actions
document.getElementById("undo-btn").addEventListener("click", undo_last);
document.getElementById("clear-btn").addEventListener("click", () => {
    if (confirm("Are you sure you want to clear the canvas? This action cannot be undone.")) {
        clear_canvas();
    }
});
document.getElementById("save-to-database-btn").addEventListener("click", () => {
    disableExitWarning();
    if (confirm("Are you sure you want to deliver this painting early?")) {
        save_to_database();
    }
});
document.getElementById("back-btn").addEventListener("click", function () {
    disableExitWarning();
    if (confirm("Are you sure you want to leave this page? Unsaved changes will be lost.")) {
        window.location.href = '/';
    }
});

// X out prevention
function showExitWarning(event) {
    event.preventDefault();
    event.returnValue = "";
}

window.addEventListener("beforeunload", showExitWarning);

function disableExitWarning() {
    window.removeEventListener("beforeunload", showExitWarning);
}

// ctrl z to undo
document.addEventListener("keydown", (event) => {
    if ((event.ctrlKey || event.metaKey) && event.key === "z") {
        event.preventDefault();
        undo_last();
    }
});

// Color fields by event listeners
document.querySelectorAll('.color-field').forEach(colorField => {
    colorField.addEventListener('click', () => {
        const color = colorField.getAttribute('data-color');
        change_color(color);
    });
});

// Add event listeners for the brush size input
document.getElementById("brush-size-input").addEventListener("input", (event) => {
    updateBrushSize(event.target.value);
});

// Brush preview setup
const previewCircle = document.createElement("div");
previewCircle.style.position = "absolute";
previewCircle.style.borderRadius = "50%";
previewCircle.style.pointerEvents = "none";
previewCircle.style.zIndex = 2;
previewCircle.style.opacity = "0.5";
previewCircle.style.border = "1px solid black";  // Set border width here
previewCircle.style.boxSizing = "border-box";    // Ensures width/height includes border
document.body.appendChild(previewCircle);

canvas.addEventListener("mousemove", (event) => {
    const x = event.clientX;
    const y = event.clientY;

    previewCircle.style.width = `${draw_width}px`;
    previewCircle.style.height = `${draw_width}px`;
    previewCircle.style.backgroundColor = draw_color;
    previewCircle.style.left = `${x - draw_width / 2}px`;
    previewCircle.style.top = `${y - draw_width / 2}px`;
    previewCircle.style.display = "block";
});

canvas.addEventListener("mouseleave", () => {
    previewCircle.style.display = "none";
});
