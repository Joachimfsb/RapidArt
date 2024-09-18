# `web` directory
The `web` directory contains all HTML/CSS/JS files.

The directory is structured in 2 folders:
* `res`
  * This folder contains static **res**ources that are hosted under the url `http://domain/res/`. This means that a file named `test.js` can automatically be reached at the url `http://domain/res/test.js` without any setup.
  * This folder can contain any web resource file like CSS, JS, PNG, etc. 
* `html`
  * This folder contains **HTML** files (both static and template files). These files are only accessible through a handler that specifically targets each file.