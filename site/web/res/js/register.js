//////////////////// STEP 1 /////////////////////
class Step1 {
    constructor() {
        this.hasBeenShown = false;
    }

    // Run only when DOM is loaded
    init() {
        // DOM
        this.step1 = document.querySelector('#step1');
        this.emailField = document.querySelector('#form-register-email');
        this.emailInfo = document.querySelector('#form-register-email-info');
        this.usernameField = document.querySelector('#form-register-username');
        this.usernameInfo = document.querySelector('#form-register-username-info');
        this.passwordField = document.querySelector('#form-register-password');
        this.passwordInfo = document.querySelector('#form-register-password-info');
        this.cpasswordField = document.querySelector('#form-register-cpassword'); // Confirm password
        this.cpasswordInfo = document.querySelector('#form-register-cpassword-info'); // Confirm password
        this.continuebtn = document.querySelector('#form-register-continue');
        this.info = document.querySelector('#form-register-info');

        // Attributes
        this.email = "";
        this.username = "";
        this.password = "";

        // Field updates
        this.emailField.addEventListener('change', function() {
            this.email = this.emailField.value;
        }.bind(this));
        this.usernameField.addEventListener('change', function(e) {
            this.username = this.usernameField.value;
        }.bind(this));
        this.passwordField.addEventListener('change', function(e) {
            this.password = this.passwordField.value;
        }.bind(this));
        // Continue button click
        this.continuebtn.addEventListener('click', this.continue.bind(this));
    }

    continue(e) {
        // Prevent sending of form by html
        e.preventDefault();

        // Validate form before sending
        if (!this.validateForm()) return;

        // Ask server to check if email and username are valid
        // Create and send an API request
        var xhr = new XMLHttpRequest();        

        // Result logic
        xhr.onreadystatechange = () => {
            if (xhr.readyState == 4) {
                // All good
                if (xhr.status == 204) {
                    GoToStep2();
                }
                // Something went wrong
                else {
                    if (xhr.responseText == "email-exists") {
                        this.emailError(true, "Email is already registered");
                        return
                    } else {
                        this.emailError();
                    } 
                    
                    if (xhr.responseText == "username-exists") {
                        this.usernameError(true, "Username already exists");
                        return;
                    } else {
                        this.usernameError();
                    }
                }
            }
        };

        xhr.open("POST", "/api/user/register/?check_email_username", true);
        xhr.send(JSON.stringify({
            email: this.email,
            username: this.username
        }));
    }

    ///////////// VISIBILITY ////////////
    show() {
        this.step1.classList.remove("hide");

        this.emailField.value = this.email;
        this.usernameField.value = this.username;
        this.passwordField.value = this.password;
        this.cpasswordField.value = this.password;

        this.hasBeenShown = true;
    }
    hide() {
        this.step1.classList.add("hide");
    }

    /////////// VALIDATION ///////////

    // Validate the form before request is sent.
    // This function returns false if there are errors, and also informs the user of the error.
    validateForm() {
        
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
        if (this.email.length == 0) {
            this.emailError(true, "Required field");
            result = false;
        } else if (validateEmail(this.email) == null) {
            this.emailError(true, "Invalid email address");
            result = false;
        } else {
            this.emailError(); // Reset errors
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
        if (this.username.length == 0) {
            this.usernameError(true, "Required field");
            result = false;
        } else if (validateUsername(this.username) == null) {
            this.usernameError(true, "Invalid username");
            result = false;
        } else {
            this.usernameError(); // Reset error
        }

        ////////////// PASSWORDS ///////////////

        // Passwords are not empty and do not match
        if (this.password.length > 0 && this.cpasswordField.value.length > 0 && this.password != this.cpasswordField.value) {
            this.cpasswordError(true);
            this.passwordError(true, "Passwords do not match")
            result = false;
        } else {
            // Check that password meets required length
            if (this.password.length < 10 && this.password.length > 0 && this.cpasswordField.value.length > 0) {
                this.cpasswordError(true);
                this.passwordError(true, "Password must be minimum 10 characters")
                result = false;

            } else {
                // Empty password field
                if (this.password.length == 0) {
                    this.passwordError(true, "Required field")
                    result = false;
                } else {
                    this.passwordError(); // Reset error
                }

                // Empty confirm password field
                if (this.cpasswordField.value.length == 0) {
                    this.cpasswordError(true, "Required field")
                    result = false;
                } else {
                    this.cpasswordError(); // Reset error
                }
            }
        }
        return result
    }


    ////////// HELPERS /////////////

    emailError(bad = false, msg = "") {
        if (bad)
            this.emailField.classList.add("red-border");
        else 
            this.emailField.classList.remove("red-border");
        this.emailInfo.textContent = msg;
    }
    
    usernameError(bad = false, msg = "") {
        if (bad)
            this.usernameField.classList.add("red-border");
        else
            this.usernameField.classList.remove("red-border");
        this.usernameInfo.textContent = msg;
    }
    
    passwordError(bad = false, msg = "") {
        if (bad)
            this.passwordField.classList.add("red-border");
        else
            this.passwordField.classList.remove("red-border");
        this.passwordInfo.textContent = msg;
    }
    
    cpasswordError(bad = false, msg = "") {
        if (bad)
            this.cpasswordField.classList.add("red-border");
        else
            this.cpasswordField.classList.remove("red-border");
        this.cpasswordInfo.textContent = msg;
    }
};


//////////////////// STEP 2 /////////////////////

class Step2 {
    constructor() {
        this.hasBeenShown = false;
    }

    init() {
        // DOM
        this.step2 = document.querySelector('#step2');
        this.displaynameField = document.querySelector('#form-register-displayname');
        this.displaynameInfo = document.querySelector('#form-register-displayname-info');
        this.profilePicElem = document.querySelector('#register-profile-pic');
        this.editprofilepicbtn = document.querySelector('#register-profile-pic-edit');
        this.backbtn = document.querySelector('#form-register-back');
        this.createaccountbtn = document.querySelector('#form-register-createaccount');
        this.info2 = document.querySelector('#form-register-info2');

        // Attributes
        this.displayname = "";
        this.profilePic = null;

        ////////// EVENTS /////////

        // Update field if displayname changes
        this.displaynameField.addEventListener('change', function(e) {
            this.displayname = this.displaynameField.value;
        }.bind(this));

        // Pre-create file upload button
        this.fileUploadElem = document.createElement("INPUT");
        this.fileUploadElem.setAttribute("type", "file");

        // Edit profile picture
        this.editprofilepicbtn.addEventListener('click', this.editProfilePic.bind(this));
        // File upload event
        this.fileUploadElem.addEventListener('change', this.fileUpload.bind(this));

        // Back button click
        this.backbtn.addEventListener('click', this.back.bind(this));

        // Create account button click
        this.createaccountbtn.addEventListener('click', this.createAccount.bind(this));

            
    }

    ///////////// EVENTS ///////////////
    back(e) {
        // Prevent sending of form by html
        e.preventDefault();

        GoToStep1();
    }

    createAccount(e) {
        // Prevent sending of form by html
        e.preventDefault();

        // Validate form before sending
        if (!this.validateForm()) return;

        // Create and send an API request
        var xhr = new XMLHttpRequest();

        const fileReader = new FileReader();

        // Result logic
        xhr.onreadystatechange = () => {
            if (xhr.readyState == 4) {
                // Logged in successfully
                if (xhr.status == 204) {
                    window.location = "/"; // Redirect to login
                }
                // Something went wrong
                else {
                    this.info2.textContent = "Something went wrong, please try again later";
                }
            }
        };

        xhr.open("POST", "/api/user/register/", true);
        if (this.profilePic) {
            // Profile pic is set
            fileReader.onload = (evt) => {
                let fileContent = evt.target.result.split(",")[1]; // Remove base64 prefix

                xhr.send(JSON.stringify({
                    email: step1.email,
                    username: step1.username,
                    password: step1.password,
                    displayname: this.displayname,
                    profile_pic: fileContent
                }));
            }
            fileReader.readAsDataURL(this.profilePic);
        } else {
            // Profile pic not set
            xhr.send(JSON.stringify({
                email: step1.email,
                username: step1.username,
                password: step1.password,
                displayname: this.displayname,
                profile_pic: null
            }));
        }
    }

    editProfilePic() {
        if (this.editprofilepicbtn.dataset.mode == "add") {
            // Add profile picture mode
            this.fileUploadElem.click(); // Trigger file upload
        } else {
            // Delete profile picture mode

            this.profilePic = null;
            this.profilePicElem.src = "/res/img/default-profile-img.png";

            // Change x to pencil
            this.editprofilepicbtn.src = "/res/icon/pencil-fill.svg";
            this.editprofilepicbtn.title = "Add profile picture";
            this.editprofilepicbtn.dataset.mode = "add";
        }
    }

    fileUpload(e) {
        let file = this.fileUploadElem.files[0];

        // If file was uploaded
        if (file) {
            // Store
            this.profilePic = file;

            // Set preview to display uploaded image
            this.profilePicElem.src = URL.createObjectURL(file);
            this.profilePicElem.onload = () => {
                URL.revokeObjectURL(this.profilePicElem.src);
            };

            // Change pencil to x
            this.editprofilepicbtn.src = "/res/icon/x-lg.svg";
            this.editprofilepicbtn.title = "Remove profile picture";
            this.editprofilepicbtn.dataset.mode = "del";
        }
    }


    //////////// VISIBILITY /////////////
    show() {
        this.step2.classList.remove("hide");

        this.displaynameField.value = this.displayname;

        this.hasBeenShown = true;
    }
    hide() {
        this.step2.classList.add("hide");
    }

    ////////////// VALIDATION //////////////
    validateForm() {    
        if (this.displayname == "") {
            this.displaynameError(true, "Required field");
            return false;
        } else {
            this.displaynameError();
        }
    
        return true;
    }

    /////////// HELPERS ////////////

    displaynameError(bad = false, msg = "") {
        if (bad)
            this.displaynameField.classList.add("red-border");
        else
            this.displaynameField.classList.remove("red-border");
        this.displaynameInfo.textContent = msg;
    }
};


let step1 = new Step1();
let step2 = new Step2();

window.addEventListener('DOMContentLoaded', function () {

    // RUNS ON LOAD
    // Init
    step1.init();
    step2.init();
    // Show step 1
    step1.show();
});


//////////// FUNCTIONS ////////////

function GoToStep1() {
    step1.show();
    step2.hide();
}

function GoToStep2() {
    // If first time showing step 2, set displayname = username
    if (!step2.hasBeenShown) step2.displayname = capitalizeFirstLetter(step1.username);
    step2.show();
    step1.hide();
}







//////////// HELPERS //////////////

function capitalizeFirstLetter(val) {
    return String(val).charAt(0).toUpperCase() + String(val).slice(1);
}