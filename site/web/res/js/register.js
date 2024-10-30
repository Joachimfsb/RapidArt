class Step1Data {
    constructor(email, username, password) {
        this.email = email;
        this.username = username;
        this.password = password;
    }

    populateStep1() {

        const emailField = document.querySelector('#form-register-email');
        const usernameField = document.querySelector('#form-register-username');
        const passwordField = document.querySelector('#form-register-password');
        const cpasswordField = document.querySelector('#form-register-cpassword'); // Confirm password
    
        emailField.value = this.email;
        usernameField.value = this.username;
        passwordField.value = this.password;
        cpasswordField.value = this.password;
    }
};
class Step2Data {
    constructor(displayname, profilePic) {
        this.displayname = displayname;
        this.profilePic = profilePic;
    }

    populateStep2() {

        const displaynameField = document.querySelector('#form-register-displayname');
        
        displaynameField.value = this.displayname;
    }
};

let step1Data = null;
let step2Data = null;

window.addEventListener('load', function () {
    const step1 = document.querySelector('#step1');
    const step2 = document.querySelector('#step2');
    const email = document.querySelector('#form-register-email');
    const username = document.querySelector('#form-register-username');
    const password = document.querySelector('#form-register-password');
    const displayname = document.querySelector('#form-register-displayname');
    const continuebtn = document.querySelector('#form-register-continue');
    const profilepic = document.querySelector('#register-profile-pic');
    const editprofilepicbtn = document.querySelector('#register-profile-pic-edit');
    const backbtn = document.querySelector('#form-register-back');
    const createaccountbtn = document.querySelector('#form-register-createaccount');
    const info1 = document.querySelector('#form-register-info1');
    const info2 = document.querySelector('#form-register-info2');

    

    // Continue button click
    continuebtn.addEventListener('click', function(e) {

        // Prevent sending of form by html
        e.preventDefault();

        // Validate form before sending
        if (!validateStep1()) return

        // Save to data structure
        step1Data = new Step1Data(email.value, username.value.toLowerCase(), password.value);

        // Ask server to check if email and username are valid
        // Create and send an API request
        var xhr = new XMLHttpRequest();        

        // Result logic
        xhr.onreadystatechange = function () {
            if (this.readyState == 4) {
                // All good
                if (this.status == 204) {
                    // Go to step 2
                    displayname.value = capitalizeFirstLetter(step1Data.username);

                    step1.classList.add("hide");
                    step2.classList.remove("hide");

                    if (step2Data != null) {
                        // User has been to step2 already
                        step2Data.populateStep2();

                    } else {
                        step2Data = new Step2Data(displayname.value, null);
                    }

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
                }
            }
        };

        xhr.open("POST", "/api/user/register/?check_email_username", true);
        xhr.send(JSON.stringify({
            email: step1Data.email,
            username: step1Data.username
        }));
    });

    // Continue button click
    backbtn.addEventListener('click', function(e) {

        // Prevent sending of form by html
        e.preventDefault();

        // Store displayname
        step2Data.displayname = displayname.value;

        // Go to step 1
        step1.classList.remove("hide");
        step2.classList.add("hide");

        // Populate values
        step1Data.populateStep1();
    });

    // Continue button click
    createaccountbtn.addEventListener('click', function(e) {

        // Prevent sending of form by html
        e.preventDefault();

        // Validate form before sending
        if (!validateStep2()) return

        // Save to data structure
        step2Data.displayname = displayname.value;

        // Create and send an API request
        var xhr = new XMLHttpRequest();

        const fileReader = new FileReader();
        

        // Result logic
        xhr.onreadystatechange = function () {
            if (this.readyState == 4) {
                // Logged in successfully
                if (this.status == 204) {
                    window.location = "/"; // Redirect to login
                }
                // Something went wrong
                else {
                    info2.textContent = "Something went wrong, please try again later";
                }
            }
        };

        xhr.open("POST", "/api/user/register/", true);
        if (step2Data.profilePic) {
            fileReader.onload = (evt) => {
                let fileContent = evt.target.result.split(",")[1]; // Remove base64 prefix

                xhr.send(JSON.stringify({
                    email: step1Data.email,
                    username: step1Data.username,
                    password: step1Data.password,
                    displayname: step2Data.displayname,
                    profile_pic: fileContent
                }));
            }
            fileReader.readAsDataURL(step2Data.profilePic);
        } else {
            xhr.send(JSON.stringify({
                email: step1Data.email,
                username: step1Data.username,
                password: step1Data.password,
                displayname: step2Data.displayname,
                profile_pic: null
            }));
        } 
    });



    // Edit profile picture button (pencil) click
    // Pre-create file upload button
    var fileUpload = document.createElement("INPUT");
    fileUpload.setAttribute("type", "file");

    editprofilepicbtn.addEventListener('click', function(e) {
        if (editprofilepicbtn.dataset.mode == "add") {
            // Add profile picture mode
            fileUpload.click(); // Trigger file upload
        } else {
            // Delete profile picture mode
            step2Data.profilePic = null;

            profilepic.src = "/res/img/default-profile-img.png";

            // Change x to pencil
            editprofilepicbtn.src = "/res/icon/pencil-fill.svg";
            editprofilepicbtn.title = "Add profile picture";
            editprofilepicbtn.dataset.mode = "add";
        }
    });
    fileUpload.addEventListener('change', function(e) {
        let file = fileUpload.files[0];

        // If file was uploaded
        if (file) {
            // Store
            step2Data.profilePic = file;

            // Set preview to display uploaded image
            profilepic.src = URL.createObjectURL(file);
            profilepic.onload = () => {
                URL.revokeObjectURL(profilepic.src);
            };

            // Change pencil to x
            editprofilepicbtn.src = "/res/icon/x-lg.svg";
            editprofilepicbtn.title = "Remove profile picture";
            editprofilepicbtn.dataset.mode = "del";
        }
    });
});


function validateStep2() {
    const displayname = document.querySelector('#form-register-displayname');

    if (displayname.value == "") {
        displaynameError(true, "Required field");
        return false;
    } else {
        displaynameError();
    }

    return true;
}


// Validate the form before request is sent.
// This function returns false if there are errors, and also informs the user of the error.
function validateStep1() {
    
    const email = document.querySelector('#form-register-email');
    const username = document.querySelector('#form-register-username');
    const password = document.querySelector('#form-register-password');
    const cpassword = document.querySelector('#form-register-cpassword'); // Confirm password

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
        document.querySelector('#form-register-email').classList.add("red-border");
    } else {
        document.querySelector('#form-register-email').classList.remove("red-border");
    }
    document.querySelector('#form-register-email-info').textContent = msg;
}

function usernameError(bad = false, msg = "") {
    if (bad) {
        document.querySelector('#form-register-username').classList.add("red-border");
    } else {
        document.querySelector('#form-register-username').classList.remove("red-border");
    }
    document.querySelector('#form-register-username-info').textContent = msg;
}

function passwordError(bad = false, msg = "") {
    if (bad) {
        document.querySelector('#form-register-password').classList.add("red-border");
    } else {
        document.querySelector('#form-register-password').classList.remove("red-border");
    }
    document.querySelector('#form-register-password-info').textContent = msg;
}

function cpasswordError(bad = false, msg = "") {
    if (bad) {
        document.querySelector('#form-register-cpassword').classList.add("red-border");
    } else {
        document.querySelector('#form-register-cpassword').classList.remove("red-border");
    }
    document.querySelector('#form-register-cpassword-info').textContent = msg;
}

function displaynameError(bad = false, msg = "") {
    if (bad) {
        document.querySelector('#form-register-displayname').classList.add("red-border");
    } else {
        document.querySelector('#form-register-displayname').classList.remove("red-border");
    }
    document.querySelector('#form-register-displayname-info').textContent = msg;
}

function capitalizeFirstLetter(val) {
    return String(val).charAt(0).toUpperCase() + String(val).slice(1);
}