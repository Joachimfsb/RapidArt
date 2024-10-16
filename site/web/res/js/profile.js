
window.addEventListener('load', function () {
    const logoutButton = document.querySelector('#logout');

    logoutButton.addEventListener('click', function(e) {
        logout(e);
    });
});
