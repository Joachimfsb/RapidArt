window.addEventListener('DOMContentLoaded', function () {

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
    else if (path.startsWith("/profile")) {
        // Profile is a special case. We also need to check if the user is visiting their own profile or not
        if (pageInfo != null && pageInfo.is_self != null && pageInfo.is_self) 
            profile.classList.add("red-item");
    }

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