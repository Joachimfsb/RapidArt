window.addEventListener('load', function () {

    // Close all dropdowns
    window.addEventListener('click', function() {
        document.querySelectorAll(".dropdown-options").forEach(opt => {
            opt.classList.add("hide");
        });
    });

    // Toggle dropdown
    const dropdownBtn = document.querySelectorAll(".dropdown-btn");
    dropdownBtn.forEach(btn => {
        btn.addEventListener('click', function(e) {
            e.stopPropagation();
            let root = e.target.parentNode;
            while (!root.classList.contains("dropdown")) root = root.parentNode; // Find root of dropdown

            const dropdownOpts = root.querySelector(".dropdown-options"); // Find options
            dropdownOpts.classList.toggle("hide"); // Toggle display of dropdown options
        });
    });

    // Select option
    const dropdownOpt = document.querySelectorAll(".dropdown-options a");
    dropdownOpt.forEach(opt => {
        opt.addEventListener('click', function(e) {
        
            // Find dropdown-options
            let opts = e.target.parentNode;
            while (!opts.classList.contains("dropdown-options")) opts = opts.parentNode;

            // Remove current selected
            const allOpts = opts.querySelectorAll("a"); // List of actual options
            allOpts.forEach(o => {
                if (o.classList.contains("selected")) {
                    o.classList.remove("selected");
                    o.querySelector("img.dropdown-icon").remove();
                }
            });

            // Add new selected
            opt.classList.add("selected");
            let img = document.createElement("img");
            img.src = "/res/icon/check-lg.svg";
            img.classList.add("dropdown-icon");
            opt.appendChild(img);

            // Update button text
            let root = e.target.parentNode;
            while (!root.classList.contains("dropdown")) root = root.parentNode; // Find root of dropdown
            let text = (opt.childNodes[0].tagName == "IMG") ? opt.childNodes[0].getAttribute("alt") : opt.textContent;
            root.querySelector(".dropdown-btn").childNodes[0].textContent = text;
        });
    });

});