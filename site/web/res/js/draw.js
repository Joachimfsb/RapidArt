const canvas = document.getElementById("canvas");
const fixedWidth = 1000;  // fast bredde som blir lagret
const fixedHeight = 600;  // fast høyde som blir lagret

// ulike kontekst, standard farge og tegne størrelse. 
let context = canvas.getContext("2d");
let draw_color = "black";
let draw_width = "2";
let is_drawing = false; // brukes som bool til å se om brukeren tegner

// Array som lagrer siste tegna linje/... (brukes til undo)
let restore_array = []; 
let index = -1;

// Endrer visuell størrelse på canvas basert på vindu størrelse, beholder opplæsning
function resizeCanvas() {
    const displayWidth = window.innerWidth > 768 ? 800 : window.innerWidth - 60; 
    const displayHeight = window.innerWidth > 768 ? 600 : 400;

    // Setter til fast størrelse (const i starten 1000x600 nå)
    canvas.width = fixedWidth;
    canvas.height = fixedHeight;

    // Visuell størrelse blir endret i CSS
    canvas.style.width = displayWidth + "px";
    canvas.style.height = displayHeight + "px";

    // Fyll canvas bakgrunn 
    context.fillStyle = "white";
    context.fillRect(0, 0, fixedWidth, fixedHeight);
}

// Ved resize/load endrer størrelse
window.addEventListener('load', resizeCanvas);
window.addEventListener('resize', resizeCanvas);

// Funksjon som henter skalafaktor mellom visuell og faktsik oppløsning (hentet fra CHATGPT)
function getScaleFactor() {
    const displayWidth = parseInt(window.getComputedStyle(canvas).width);
    const displayHeight = parseInt(window.getComputedStyle(canvas).height);
    // Returnerer skala som brukes for å justere mus/touch posisjon
    return {
        x: canvas.width / displayWidth,
        y: canvas.height / displayHeight
    };
}

// Funksjon for ny tegnehandling 
function start(event) {
    is_drawing = true;
    const scale = getScaleFactor();
    context.beginPath();
    context.moveTo(event.offsetX * scale.x, event.offsetY * scale.y);  // Adjust for scaling
    event.preventDefault();
}

// Tegne funksjon når musen beveger seg
function draw(event) {
    if (is_drawing) { // sjekk om man tegner
        const scale = getScaleFactor();
        context.lineTo(event.offsetX * scale.x, event.offsetY * scale.y);  
        context.strokeStyle = draw_color;
        context.lineWidth = draw_width;
        context.lineCap = "round";
        context.lineJoin = "round";
        context.stroke();
    }
    event.preventDefault();
}

// Funksjon som stopper når man løfter museklikk eller musen går av canvaset
function stop(event) {
    if (is_drawing) {
        context.stroke();
        context.closePath();
        is_drawing = false;
    }
    event.preventDefault();

    if (event.type != 'mouseout') {
        restore_array.push(context.getImageData(0, 0, canvas.width, canvas.height));
        index += 1;
    }
}

// Angre siste handling (bruker index og array til å huske siste)
function undo_last() {
    if (index <= 0) {
        clear_canvas();
    } else {
        index -= 1;
        restore_array.pop();
        context.putImageData(restore_array[index], 0, 0);
    }
}

// Fjerner alt på canvaset
function clear_canvas() {
    context.fillStyle = "white";
    context.clearRect(0, 0, canvas.width, canvas.height);
    context.fillRect(0, 0, canvas.width, canvas.height);

    restore_array = [];
    index = -1;
}

// Lagrer hva enn som er på canvaset som en PNG på brukerens datamaskin
function save_as_png() {
    const dataUrl = canvas.toDataURL("image/png", 1.0);  // Lagrer som 1000x600
    const link = document.createElement('a');
    link.href = dataUrl;
    link.download = 'drawing.png';
    link.click();
}

// Test funksjon for lagring til database, lagres som BLOB (CHATGPT)
function save_to_database() {
    const dataUrl = canvas.toDataURL("image/png"); // Data URL for bildet
    const imageData = dataUrl.replace(/^data:image\/(png|jpg);base64,/, ""); // Fjerner base64-header

    //XMLHttpRequest for å sende dataen til serveren
    const xhr = new XMLHttpRequest();
    xhr.open("POST", "http://localhost:8080/save_image", true); //POST forrespørsel
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    xhr.onreadystatechange = function() { // Håndterer respons serverens respons
        if (xhr.readyState === 4) {
            console.log("XHR Status: " + xhr.status);
            console.log("Server Response: " + xhr.responseText);

            if (xhr.status === 200) {
                alert("Image saved to database!");
            } else {
                alert("Error saving image: " + xhr.status);
            }
        }
    };
    xhr.send("image_data=" + encodeURIComponent(imageData)); //Sender bildet til servern
}

// Ulike event listeners for handlinger
canvas.addEventListener("touchstart", start, false); // Start tegning ved touch
canvas.addEventListener("touchmove", draw, false); // Tegn når det beveges
canvas.addEventListener("mousedown", start, false); // STart ved museklikk
canvas.addEventListener("mousemove", draw, false); // Tegn når musen beveger seg
 
canvas.addEventListener("touchend", stop, false); // Stopp ved ingen touch
canvas.addEventListener("mouseup", stop, false); // Stopp ved mouse-up
canvas.addEventListener("mouseout", stop, false); // Stopp hvis musen går utenfor canvaset

// Lagrer PNG (brukes til at bildet fungerer som knapp)
document.getElementById("save-as-png-btn").addEventListener("click", save_as_png);

// Lagrer til database (brukes til at bildet fungerer som knapp)
document.getElementById("save-to-database-btn").addEventListener("click", save_to_database);

// Fargevelger
document.getElementById("color-picker-icon").addEventListener("click", function() {
    document.getElementById("color-picker").click();  // Trigger the hidden color picker input
});
