window.addEventListener('load', function () {

    const logoutButton = document.querySelector('#logout');

    if (logoutButton != null) {
        logoutButton.addEventListener('click', function(e) {
            logout(e);
        });
    }

    
      
    const shareBtn = document.querySelector("button#share");
    
    // Share must be triggered by "user activation"
    if (navigator.share) {
        shareBtn.addEventListener("click", async () => {
            const displayname = document.querySelector("#displayname").textContent;

            const shareData = {
                title: "RapidArt profile",
                text: "View " + displayname + "'s profile",
                url: this.window.location,
            };

            try {
                await navigator.share(shareData);
            } catch (err) {
            }
        });
    } else {
        // Hide share button
        shareBtn.classList.add("hide");
    }
      



    const unfollowBtn = document.querySelector("button#unfollow");

    if (unfollowBtn) {
        unfollowBtn.addEventListener('mouseover', function(e) {
            unfollowBtn.querySelector("img").src = "/res/icon/person-fill-x-white.svg";
        });
        unfollowBtn.addEventListener('mouseleave', function(e) {
            unfollowBtn.querySelector("img").src = "/res/icon/person-fill-check-white.svg";
        });
    }
});
