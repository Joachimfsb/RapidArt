window.addEventListener('load', function () {

    let DOMPostImgWrapper = document.querySelector("#post-img-wrapper");
    
    DOMPostImgWrapper.addEventListener('mousedown', showBasisCanvas);
    DOMPostImgWrapper.addEventListener('mouseup', hideBasisCanvas);
    DOMPostImgWrapper.addEventListener('touchstart', showBasisCanvas);
    DOMPostImgWrapper.addEventListener('touchend', hideBasisCanvas);
    DOMPostImgWrapper.addEventListener('touchcancel', hideBasisCanvas);
});






function showBasisCanvas() {
    let DOMPostImg = document.querySelector("#post-img");
    let DOMPostBasisCanvas = document.querySelector("#post-basis-canvas");

    DOMPostImg.classList.add("hide");
    DOMPostBasisCanvas.classList.remove("hide");
}

function hideBasisCanvas() {
    let DOMPostImg = document.querySelector("#post-img");
    let DOMPostBasisCanvas = document.querySelector("#post-basis-canvas");

    DOMPostImg.classList.remove("hide");
    DOMPostBasisCanvas.classList.add("hide");
}