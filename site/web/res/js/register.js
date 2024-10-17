
window.addEventListener('load', function () {
    const email = document.querySelector('#form-login-email');
    const username = document.querySelector('#form-login-username');
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
                    window.location = "/"; // Redirect to login
                }        
                // Something went wrong
                else {
                    if (xhr.responseText == "email-exists") {
                        emailError(true, "Email is already registered");
                        return
                    } else {
                        emailError();
                    } 
                    
                    if (xhr.responseText == "username-exists") {
                        usernameError(true, "Username already exists");
                        return;
                    } else {
                        usernameError();
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
    const username = document.querySelector('#form-login-username');
    const password = document.querySelector('#form-login-password');
    const cpassword = document.querySelector('#form-login-cpassword'); // Confirm password

    let result = true;

    ////////////// EMAIL /////////////////

    const validateEmail = (email) => {
        return String(email)
            .toLowerCase()
            .match(
                /^(?:[a-z0-9!#$%&'*+/=?^_{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_{|}~-]+)*|"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\[(?:(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9]))\.){3}(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9])|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\])?/
        );
    };

    // Empty email name field
    if (email.value.length == 0) {
        emailError(true, "Required field");
        result = false;
    } else if (validateEmail(email.value) == null) {
        emailError(true, "Invalid email address");
        result = false;
    } else {
        emailError();
    }

    ////////////// USERNAME ////////////////

    const validateUsername = (username) => {
        return String(username)
            .toLowerCase()
            .match(
                /^[a-zA-Z0-9]+$/
        );
    };

    // Empty username field
    if (username.value.length == 0) {
        usernameError(true, "Required field");
        result = false;
    } else if (validateUsername(username.value) == null) {
        usernameError(true, "Invalid username");
        result = false;
    } else {
        usernameError();
    }

    ////////////// PASSWORDS ///////////////

    // Passwords are not empty and do not match
    if (password.value.length > 0 && cpassword.value.length > 0 && password.value != cpassword.value) {

        cpasswordError(true);
        passwordError(true, "Passwords do not match")
        result = false;
    } else {
        // Check that password meets required length
        if (password.value.length < 10 && password.value.length > 0 && cpassword.value.length > 0) {
            cpasswordError(true);
            passwordError(true, "Password must be minimum 10 characters")
            result = false;

        } else {
            // Empty password field
            if (password.value.length == 0) {
                passwordError(true, "Required field")
                result = false;
            } else {
                passwordError();
            }

            // Empty confirm password field
            if (cpassword.value.length == 0) {
                cpasswordError(true, "Required field")
                result = false;
            } else {
                cpasswordError()
            }

        }
        
    }

    return result
}

//////////// HELPERS //////////////

function emailError(bad = false, msg = "") {
    if (bad) {
        document.querySelector('#form-login-email').classList.add("red-border");
    } else {
        document.querySelector('#form-login-email').classList.remove("red-border");
    }
    document.querySelector('#form-login-email-info').textContent = msg;
}

function usernameError(bad = false, msg = "") {
    if (bad) {
        document.querySelector('#form-login-username').classList.add("red-border");
    } else {
        document.querySelector('#form-login-username').classList.remove("red-border");
    }
    document.querySelector('#form-login-username-info').textContent = msg;
}

function passwordError(bad = false, msg = "") {
    if (bad) {
        document.querySelector('#form-login-password').classList.add("red-border");
    } else {
        document.querySelector('#form-login-password').classList.remove("red-border");
    }
    document.querySelector('#form-login-password-info').textContent = msg;
}

function cpasswordError(bad = false, msg = "") {
    if (bad) {
        document.querySelector('#form-login-cpassword').classList.add("red-border");
    } else {
        document.querySelector('#form-login-cpassword').classList.remove("red-border");
    }
    document.querySelector('#form-login-cpassword-info').textContent = msg;
}