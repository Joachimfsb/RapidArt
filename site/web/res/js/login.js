const loginButton = document.getElementById('login');

loginButton.addEventListener('mouseover', function() {
    loginButton.src = '../../../images/login_hover.png';
});

loginButton.addEventListener('mouseout', function() {
    loginButton.src = '../../../images/login.png';
});
