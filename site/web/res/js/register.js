
window.addEventListener('load', function () {
    const email = document.querySelector('#form-login-email');
    const emailInfo = document.querySelector('#form-login-email-info');
    const username = document.querySelector('#form-login-username');
    const usernameInfo = document.querySelector('#form-login-username-info');
    const password = document.querySelector('#form-login-password');
    const submit = document.querySelector('#form-login-submit');
    const info = document.querySelector('#form-login-info');

    // Submit button click
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
                    window.location = "/login/"; // Redirect to login
                }        
                // Something went wrong
                else {
                    if (xhr.responseText == "email-exists") {
                        // Notify user
                        email.classList.add("red-border");
                        emailInfo.textContent = "Email is already registered";
                        return
                    } else {
                        email.classList.remove("red-border");
                        emailInfo.textContent = "";
                    } 
                    
                    if (xhr.responseText == "username-exists") {
                        
                        username.classList.add("red-border");
                        usernameInfo.textContent = "Username already exists";
                        return;
                    } else {
                        username.classList.remove("red-border");
                        usernameInfo.textContent = "";
                    }

                    info.textContent = "Something went wrong, please try again later";
                }
            }
        };

        xhr.open("POST", "/api/user/register/", true);
        xhr.send(JSON.stringify({
            email: email.value,
            username: username.value,
            password: password.value
        }));
    });
});



// Validate the form before request is sent.
// This function returns false if there are errors, and also informs the user of the error.
function validateForm() {
    const email = document.querySelector('#form-login-email');
    const emailInfo = document.querySelector('#form-login-email-info');
    const username = document.querySelector('#form-login-username');
    const usernameInfo = document.querySelector('#form-login-username-info');
    const password = document.querySelector('#form-login-password');
    const passwordInfo = document.querySelector('#form-login-password-info');
    const cpassword = document.querySelector('#form-login-cpassword'); // Confirm password
    const cpasswordInfo = document.querySelector('#form-login-cpassword-info'); // Confirm password
    const submit = document.querySelector('#form-login-submit');
    const info = document.querySelector('#form-login-info');

    let good = true;

    ////////////// EMAIL /////////////////

    const validateEmail = (email) => {
        return String(email)
            .toLowerCase()
            .match(
                /^[a-z0-9!#$%&'*+/=?^_{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?/
        );
    };

    // Empty email name field
    if (email.value.length == 0) {
        email.classList.add("red-border");
        emailInfo.textContent = "Required field";
        good = false;
    } else if (validateEmail(email.value) == null) {
        email.classList.add("red-border");
        emailInfo.textContent = "Invalid email address";
        good = false;
    } else {
        email.classList.remove("red-border");
        emailInfo.textContent = "";
    }

    ////////////// USERNAME ////////////////

    const validateUsername = (username) => {
        return String(username)
            .toLowerCase()
            .match(
                /^[a-zA-Z0-9]+/
        );
    };

    // Empty username field
    if (username.value.length == 0) {
        
        username.classList.add("red-border");
        usernameInfo.textContent = "Required field";
        good = false;
    } else if (validateUsername(username.value) == null) {

        username.classList.add("red-border");
        usernameInfo.textContent = "Invalid username";
        good = false;
    } else {
        username.classList.remove("red-border");
        usernameInfo.textContent = "";
    }

    ////////////// PASSWORDS ///////////////

    // Passwords are not empty and do not match
    if (password.value.length > 0 &&
        cpassword.value.length > 0 &&
        password.value != cpassword.value) {



        password.classList.add("red-border");
        cpassword.classList.add("red-border");
        cpasswordInfo.textContent = "";
        passwordInfo.textContent = "Passwords do not match";
        good = false;
    } else {
        // Check that password meets required length
        if (password.value.length < 10 && password.value.length > 0 && cpassword.value.length > 0) {
            cpassword.classList.add("red-border");
            cpasswordInfo.textContent = "";
            password.classList.add("red-border");
            passwordInfo.textContent = "Password must be minimum 10 characters";
            good = false;
        } else {
            // Empty password field
            if (password.value.length == 0) {

                password.classList.add("red-border");
                passwordInfo.textContent = "Required field";
                good = false;
            } else {
                password.classList.remove("red-border");
                passwordInfo.textContent = "";
            }

            // Empty confirm password field
            if (cpassword.value.length == 0) {

                cpassword.classList.add("red-border");
                cpasswordInfo.textContent = "Required field";
                good = false;
            } else {
                cpassword.classList.remove("red-border");
                cpasswordInfo.textContent = "";
            }

        }
        
    }

    return good
}