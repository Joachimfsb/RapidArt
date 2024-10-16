
window.addEventListener('load', function () {
    const username = document.querySelector('#form-login-username');
    const usernameInfo = document.querySelector('#form-login-username-info');
    const password = document.querySelector('#form-login-password');
    const passwordInfo = document.querySelector('#form-login-password-info');
    const submit = document.querySelector('#form-login-submit');
    const info = document.querySelector('#form-login-info');

    // Submit button clock
    submit.addEventListener('click', function(e) {

        // Prevent sending of form by html
        e.preventDefault();

        // Validate form before sending
        if (!validateForm()) return

        // Create and send an API request
        var xhr = new XMLHttpRequest();

        // Result logic
        xhr.onreadystatechange = function () {
            if (this.readyState == 4) {
                // Logged in successfully
                if (this.status == 204) {
                    window.location = "/";
                }        
                // Something went wrong
                else {
                    if (xhr.responseText == "bad-user") {
                        // Notify user
                        username.classList.add("red-border");
                        usernameInfo.textContent = "User does not exist";
                        return;
                    } else {
                        username.classList.remove("red-border");
                        usernameInfo.textContent = "";
                    }

                    if (xhr.responseText == "bad-pass") {
                        // Notify user
                        password.classList.add("red-border");
                        passwordInfo.textContent = "Incorrect password";
                        return;
                    } else {
                        password.classList.remove("red-border");
                        passwordInfo.textContent = "";
                    }

                    info.textContent = "Something went wrong, please try again later";
                }
            }
        };

        xhr.open("POST", "/api/auth/login/", true);
        xhr.send(JSON.stringify({
            username: username.value,
            password: password.value
        }));
    });
});




function validateForm() {
    const username = document.querySelector('#form-login-username');
    const usernameInfo = document.querySelector('#form-login-username-info');
    const password = document.querySelector('#form-login-password');
    const passwordInfo = document.querySelector('#form-login-password-info');
    const generalInfo = document.querySelector('#form-login-info');

    let good = true;

    // Empty user name field
    if (username.value.length == 0) {
        // Notify user
        username.classList.add("red-border");
        usernameInfo.textContent = "Required field";
        good = false;
    } else {
        username.classList.remove("red-border");
        usernameInfo.textContent = "";
    }

    // Empty password field
    if (password.value.length == 0) {
        // Notify user
        password.classList.add("red-border");
        passwordInfo.textContent = "Required field";
        good = false;
    } else {
        password.classList.remove("red-border");
        passwordInfo.textContent = "";
    }

    return good
}