// Data (with defaults)
let selectedType = "top-posts"
let selectedBasisCanvas = "all"
let selectedSince = new Date(new Date().getTime() - (24*60*60*1000)) // One day ago

window.addEventListener('load', function () {

    updateList();

    // Type selector
    document.querySelectorAll("#dropdown-type .dropdown-options a").forEach(opt => {
        opt.addEventListener('click', function(e) {
            selectedType = e.target.dataset.val;

            if (selectedType == "top-posts") {
                document.querySelector("#dropdown-basiscanvas").classList.remove("hide");
                document.querySelector("#dropdown-since").classList.remove("hide");
            } else {
                document.querySelector("#dropdown-basiscanvas").classList.add("hide");
                document.querySelector("#dropdown-since").classList.add("hide");
            }

            updateList();
        });
    });

    // BasisCanvas selector
    document.querySelectorAll("#dropdown-basiscanvas .dropdown-options a").forEach(opt => {
        opt.addEventListener('click', function(e) {
            selectedBasisCanvas = (e.target.tagName == "IMG") ? e.target.parentNode.dataset.val : e.target.dataset.val;
            updateList();
        });
    });

    // BasisCanvas selector
    document.querySelectorAll("#dropdown-since .dropdown-options a").forEach(opt => {
        opt.addEventListener('click', function(e) {
            switch (e.target.dataset.val) {
                case "hour":
                    selectedSince = new Date(new Date().getTime() - (60*60*1000));
                    break;
                case "day":
                    selectedSince = new Date(new Date().getTime() - (24*60*60*1000));
                    break;
                case "week":
                    selectedSince = new Date(new Date().getTime() - (7*24*60*60*1000));
                    break;
                case "month":
                    selectedSince = new Date(new Date().getTime() - (30*24*60*60*1000));
                    break;
                case "year":
                    selectedSince = new Date(new Date().getTime() - (365*24*60*60*1000));
                    break;
                case "all":
                    selectedSince = null;
                    break;
            }

            updateList();
        });
    });

});

function updateList() {
    // Prep data
    let query = [];
    if (selectedSince != null) query.push("since=" + selectedSince.toISOString());
    if (selectedBasisCanvas != "all") query.push("basiscanvas=" + selectedBasisCanvas);
    
    // Fetch new list from server
    var xhr = new XMLHttpRequest();

    // Follow request
    xhr.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            // Update list
            document.querySelector("#results").innerHTML = xhr.responseText;
        }
    };

    if (selectedType == "top-posts") {
        xhr.open("GET", "/api/top/posts?" + query.join("&"), true);
        xhr.send();
    } else if (selectedType == "top-users") {
        xhr.open("GET", "/api/top/users", true);
        xhr.send();
    }
}