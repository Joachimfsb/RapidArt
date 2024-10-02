//Funksjon som henter URL parameter (CHATGPT)
function getParameterByName(name) {
    const url = window.location.href;
    name = name.replace(/[\[]/, '\\[').replace(/[\]]/, '\\]');
    const regex = new RegExp('[\\?&]' + name + '=([^&#]*)');
    const results = regex.exec(url);
    return results === null ? '' : decodeURIComponent(results[1].replace(/\+/g, ' '));
}

// henter parameteren fra URLen
let drawingId = parseInt(getParameterByName('drawing')) || 1; // Default to 1 if not specified

// setter bilde basert på parametern (antar at bilde er lagret som drawing1 2,3... (må mulig endres mot database)
const drawingImage = document.getElementById('drawingImage');
if (drawingId) {
    drawingImage.src = `/res/img/drawing${drawingId}.png`; // Assuming images are named drawing1.png, drawing2.png, etc.
} else {
    drawingImage.src = '/res/img/drawing.png'; // Default image if no drawing is specified
}

// Navigasjon (forrige innlegg)
document.getElementById('prevPost').addEventListener('click', function(event) {
    event.preventDefault();
    if (drawingId > 1) {
        drawingId--;
        window.location.href = `/post/?drawing=${drawingId}`;
    }
});

// Navigasjon (neste innlegg)
document.getElementById('nextPost').addEventListener('click', function(event) {
    event.preventDefault();
    drawingId++;
    window.location.href = `/post/?drawing=${drawingId}`;
});
