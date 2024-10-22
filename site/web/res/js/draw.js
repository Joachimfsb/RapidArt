const canvas = document.getElementById("canvas");
const fixedWidth = 700;  // fast bredde som blir lagret
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
    const container = document.querySelector('.canvas-container');
    const displayWidth = container.clientWidth; // Bruk tilgjengelig bredde i containeren
    const aspectRatio = fixedWidth / fixedHeight; // Holder aspektforholdet konstant

    // Beregn høyden basert på bredden og aspektforholdet
    const displayHeight = displayWidth / aspectRatio;

    // Sett canvas størrelse til fast oppløsning (700x600)
    canvas.width = fixedWidth;
    canvas.height = fixedHeight;

    // Juster visuell størrelse på canvas, samtidig som aspektforholdet holdes konstant
    canvas.style.width = displayWidth + "px";
    canvas.style.height = displayHeight + "px";

    // Fyll bakgrunn på canvas (for eksempel)
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

    // Basis bildet må være lastet før lagring
    if (!basisImage.complete) {
        basisImage.onload = save_as_png;
        return;
    }

    // Lager et midlertidig canvas med fast oppløsning
    const tempCanvas = document.createElement('canvas');
    tempCanvas.width = fixedWidth;
    tempCanvas.height = fixedHeight;
    const tempContext = tempCanvas.getContext('2d');

    // Tegner innholdet fra tegnecanvas
    tempContext.drawImage(drawingCanvas, 0, 0, fixedWidth, fixedHeight);

    // Skalerer basisbildet og sentrerer det
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

    // Tegner basisbildet på toppen
    tempContext.drawImage(basisImage, basisX, basisY, basisWidth, basisHeight);

    // Lagrer resultatet som PNG
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

    const userId = 1;
    const timeSpentDrawing = 1;

    const postData = {
        image_data: mergedImageData,
        basis_canvas_id: basisCanvasId,
        user_id: userId,
        caption: '',
        time_spent_drawing: timeSpentDrawing,
    };

    const xhr = new XMLHttpRequest();
    xhr.open("POST", "/api/save_post", true);
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.onreadystatechange = function() {
        if (xhr.readyState === 4) {
            if (xhr.status === 200) {
                alert("Post lagret til databasen!");
            } else {
                alert("Kunne ikke lagre post. Status: " + xhr.status);
                console.log("Feilmelding:", xhr.responseText);
            }
        }
    };
    xhr.send(JSON.stringify(postData));
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
