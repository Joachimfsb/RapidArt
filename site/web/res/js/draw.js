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
    const basisImage = document.getElementById('basis');
    const drawingCanvas = document.getElementById('canvas');

    // Basis bildet er lastet før lagring
    if (!basisImage.complete) {
        basisImage.onload = save_as_png;
        return;
    }

    // Visuell størrelse av canvas
    const displayWidth = drawingCanvas.clientWidth;
    const displayHeight = drawingCanvas.clientHeight;

    // Midlertidig canvas hvor basiscanvas og interactivecanvas er kombinert
    const tempCanvas = document.createElement('canvas');
    tempCanvas.width = displayWidth;
    tempCanvas.height = displayHeight;
    const tempContext = tempCanvas.getContext('2d');

    // Tegner interaktiv tegne kanvas på temp canvas
    tempContext.drawImage(drawingCanvas, 0, 0, displayWidth, displayHeight);

    // Henter aspectratio for å sørge for riktig størrelse ved lagring for basisbildet
    const basisAspectRatio = basisImage.naturalWidth / basisImage.naturalHeight;

    // Størrelse for basisbilde
    let basisWidth, basisHeight;
    if (displayWidth / displayHeight > basisAspectRatio) {
        basisHeight = displayHeight;
        basisWidth = basisHeight * basisAspectRatio;
    } else {
        basisWidth = displayWidth;
        basisHeight = basisWidth / basisAspectRatio;
    }

    // Posisjon for å ha basisbilde i midten
    const basisX = (displayWidth - basisWidth) / 2;
    const basisY = (displayHeight - basisHeight) / 2;

    // Tegner basisbildet til slutt på topp, med riktig størrelse
    tempContext.drawImage(basisImage, basisX, basisY, basisWidth, basisHeight);

    // Lagrer bilde som PNG
    const dataUrl = tempCanvas.toDataURL("image/png");
    const link = document.createElement('a');
    link.href = dataUrl;
    link.download = 'drawing.png';
    link.click();
}

// Funksjon for lagring til database (inneholder en del midlertidige funksjonaliteter for testing)
function save_to_database() {
    const basisImage = document.getElementById('basis');
    const drawingCanvas = document.getElementById('canvas');

    if (!basisImage.complete) {
        basisImage.onload = save_to_database;
        return;
    }

    const displayWidth = drawingCanvas.clientWidth;
    const displayHeight = drawingCanvas.clientHeight;

    const tempCanvas = document.createElement('canvas');
    tempCanvas.width = displayWidth;
    tempCanvas.height = displayHeight;
    const tempContext = tempCanvas.getContext('2d');

    tempContext.drawImage(drawingCanvas, 0, 0, displayWidth, displayHeight);

    const basisAspectRatio = basisImage.naturalWidth / basisImage.naturalHeight;

    let basisWidth, basisHeight;
    if (displayWidth / displayHeight > basisAspectRatio) {
        basisHeight = displayHeight;
        basisWidth = basisHeight * basisAspectRatio;
    } else {
        basisWidth = displayWidth;
        basisHeight = basisWidth / basisAspectRatio;
    }

    const basisX = (displayWidth - basisWidth) / 2;
    const basisY = (displayHeight - basisHeight) / 2;

    tempContext.drawImage(basisImage, basisX, basisY, basisWidth, basisHeight);

    // Konverterer til base64 (PNG format og fjerner header)
    const mergedDataUrl = tempCanvas.toDataURL("image/png");
    const mergedImageData = mergedDataUrl.replace(/^data:image\/png;base64,/, "");

    // Henter basiscanvasID og gjør til int
    const basisCanvasId = parseInt(basisImage.dataset.basisCanvasId, 10);

    // Id sjekk
    if (!basisCanvasId || basisCanvasId === 0) {
        alert("Invalid BasisCanvasId");
        return;
    }

    // Mildertidig user/tid brukt verdier
    const userId = 1;
    const timeSpentDrawing = 1;

    // Objekt some skal sendes til database (userid, basiscanvasid, bilde, caption, timespent)
    const postData = {
        image_data: mergedImageData,
        basis_canvas_id: basisCanvasId,
        user_id: userId,
        caption: '',
        time_spent_drawing: timeSpentDrawing,
    };

    // XMLHttprequest for å sende dataen
    const xhr = new XMLHttpRequest();
    xhr.open("POST", "/api/save_post", true); // Post request til en api endpoint (f.eks. save_post)
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.onreadystatechange = function() {
        if (xhr.readyState === 4) {
            if (xhr.status === 200) { // Alt går
                alert("Post saved to the database!");
            } else { // Feil
                alert("Failed to save post. Status: " + xhr.status);
                console.log("Error Response:", xhr.responseText);
            }
        }
    };
    xhr.send(JSON.stringify(postData)); // Sender dataen
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
