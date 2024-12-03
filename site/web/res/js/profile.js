window.addEventListener('DOMContentLoaded', function () {

    const logoutButton = document.querySelector('#logout');

    if (logoutButton != null) {
        logoutButton.addEventListener('click', function(e) {
            logout(e);
        });
    }


      
    const shareBtn = document.querySelector("button#share");
    
    if (shareBtn) {
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
    }
      



    const followBtn = document.querySelector("button#follow");
    const unfollowBtn = document.querySelector("button#unfollow");

    // Unfollow
    if (unfollowBtn) {
        unfollowBtn.addEventListener('mouseover', function(e) {
            unfollowBtn.querySelector("img").src = "/res/icon/person-fill-x-white.svg";
        });
        unfollowBtn.addEventListener('mouseleave', function(e) {
            unfollowBtn.querySelector("img").src = "/res/icon/person-fill-check-white.svg";
        });

        unfollowBtn.addEventListener('click', function(e) {
            var xhr = new XMLHttpRequest();

            // Follow request
            xhr.onreadystatechange = function() {
                if (this.readyState == 4 && this.status == 204) {
                    // Toggle buttons
                    unfollowBtn.classList.add("hide");
                    followBtn.classList.remove("hide");

                    // Reduce followers counter
                    let stat = document.querySelector("#stat-followers .stat-val");
                    stat.textContent = parseInt(stat.textContent) - 1;
                }
            };

            xhr.open("POST", "/api/user/follow/" + pageInfo.user_id + "/0", true);
            xhr.send();
        });
    }


    // Follow
    if (followBtn) {
        followBtn.addEventListener('click', function(e) {
            var xhr = new XMLHttpRequest();

            // Follow request
            xhr.onreadystatechange = function() {
                if (this.readyState == 4 && this.status == 204) {
                    // Toggle buttons
                    followBtn.classList.add("hide");
                    unfollowBtn.classList.remove("hide");

                    // Increment followers counter
                    let stat = document.querySelector("#stat-followers .stat-val");
                    stat.textContent = parseInt(stat.textContent) + 1;
                }
            };

            xhr.open("POST", "/api/user/follow/" + pageInfo.user_id + "/1", true);
            xhr.send();
        });
    }
});
