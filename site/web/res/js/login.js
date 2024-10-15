
window.addEventListener('load', function () {
    const username = document.querySelector('#form-login-username');
    const password = document.querySelector('#form-login-password');
    const submit = document.querySelector('#form-login-submit');

    submit.addEventListener('click', function(e) {

        e.preventDefault();
        
        var xhr = new XMLHttpRequest();

        xhr.onreadystatechange = function () {
            if (this.readyState == 4) {
                if (this.status == 204) {
                    window.location = "/";
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

