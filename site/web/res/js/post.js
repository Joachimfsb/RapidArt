window.addEventListener('load', function () {

    fetchAndUpdateComments();


    // Post image click
    let DOMPostImgWrapper = document.querySelector("#post-img-wrapper");
    
    DOMPostImgWrapper.addEventListener('mousedown', showBasisCanvas);
    DOMPostImgWrapper.addEventListener('mouseup', hideBasisCanvas);
    DOMPostImgWrapper.addEventListener('touchstart', showBasisCanvas);
    DOMPostImgWrapper.addEventListener('touchend', hideBasisCanvas);
    DOMPostImgWrapper.addEventListener('touchcancel', hideBasisCanvas);


    // Like click
    document.querySelector("#like-wrapper").addEventListener('click', toggleLike);

    // Comment
    document.querySelector("#new-comment-btn").addEventListener('click', postComment);
});





////////// BASIS CANVAS ////////////
function showBasisCanvas() {
    let DOMPostImg = document.querySelector("#post-img");
    let DOMPostBasisCanvas = document.querySelector("#post-basis-canvas");

    DOMPostBasisCanvas.classList.remove("hide");
    DOMPostImg.classList.add("hide");
}

function hideBasisCanvas() {
    let DOMPostImg = document.querySelector("#post-img");
    let DOMPostBasisCanvas = document.querySelector("#post-basis-canvas");

    DOMPostImg.classList.remove("hide");
    DOMPostBasisCanvas.classList.add("hide");
}

/////////// LIKE ////////////
function toggleLike() {
    let wrapper = document.querySelector("#like-wrapper");

    var xhr = new XMLHttpRequest();

    if (wrapper.dataset.liked == "1") {
        // Unlike
        xhr.onreadystatechange = function() {
            if (this.readyState == 4 && this.status == 204) {
                // Update DOM
                wrapper.querySelector("#interaction-icon-like").src = "/res/icon/heart.svg"; // Icon
                let likeCount = wrapper.querySelector("#like-count");                        // Like count
                likeCount.textContent = parseInt(likeCount.textContent) - 1;
                likeCount.classList.remove("red", "bold");
                wrapper.title = "Like this post";                                            // Title
                wrapper.dataset.liked = "0";                                                 // Data
            }
        };
    
        xhr.open("POST", "/api/post/unlike/" + pageInfo.post_id, true);
    } else {
        // Like
        xhr.onreadystatechange = function() {
            if (this.readyState == 4 && this.status == 204) {
                // Update DOM
                wrapper.querySelector("#interaction-icon-like").src = "/res/icon/heart-fill-red.svg"; // Icon
                let likeCount = wrapper.querySelector("#like-count");                                 // Like count
                likeCount.textContent = parseInt(likeCount.textContent) + 1;
                likeCount.classList.add("red", "bold");
                wrapper.title = "Unlike this post";                                                   // Title
                wrapper.dataset.liked = "1";                                                          // Data
            }
        };
    
        xhr.open("POST", "/api/post/like/" + pageInfo.post_id, true);
    }

    xhr.send();
}


//////////// COMMENT /////////////
function fetchAndUpdateComments() {
    var xhr = new XMLHttpRequest();

    // Comment
    xhr.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            let res = xhr.responseText;

            // Replace escaped <br> with original <br>
            res = res.replaceAll("&lt;br&gt;", "<br>");


            document.querySelector("#comments").innerHTML = res;
        }
    };

    xhr.open("GET", "/post/comments/" + pageInfo.post_id, true);
    xhr.send();
}


function postComment() {

    // Fetch data
    let message = document.querySelector("#new-comment-msg").value;

    // Validate
    if (message.length > 512 || message.length < 1) {
        // Display error
        document.querySelector("#new-comment-msg").classList.add("red-border");
        return;
    } else {
        document.querySelector("#new-comment-msg").classList.remove("red-border");
    }

    var xhr = new XMLHttpRequest();

    // Like
    xhr.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 204) {
            // Update DOM
            document.querySelector("#new-comment-msg").value = "";
            let cc = document.querySelector("#comment-count");
            cc.textContent = parseInt(cc.textContent) + 1;


            // Refresh comments
            fetchAndUpdateComments();
        }
    };

    xhr.open("POST", "/api/post/comment/" + pageInfo.post_id, true);
    xhr.setRequestHeader('Content-Type', 'application/json');

    xhr.send(JSON.stringify({
        message: message
    }));
}