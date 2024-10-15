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