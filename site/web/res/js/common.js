window.addEventListener('load', function () {

    ///////// HEADER ///////////
    // Color header button based on what page we are on
    let home = document.querySelector("#nav-home");
    let search = document.querySelector("#nav-search");
    let toplist = document.querySelector("#nav-toplist");
    let profile = document.querySelector("#nav-profile");
    
    let path = window.location.pathname;

    if (path == "/") home.classList.add("red-item");
    else if (path.startsWith("/search")) search.classList.add("red-item");
    else if (path.startsWith("/toplist")) toplist.classList.add("red-item");
    else if (path.startsWith("/profile")) profile.classList.add("red-item");

});


// Sends an API request to log the user out
function logout(e) {
    e.preventDefault();
        
    var xhr = new XMLHttpRequest();

    xhr.onreadystatechange = function () {
        if (this.readyState == 4) {
            if (this.status == 204) {
                window.location = "/";
            }        
        }
    };

    xhr.open("POST", "/api/auth/logout/", true);
    xhr.send();
}