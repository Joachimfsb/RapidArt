package main

import (
	"log"
	"net/http"
	"path/filepath"
	"rapidart/internal/database"
	"rapidart/internal/glob"
	"rapidart/internal/handlers"
	"rapidart/internal/util"
)

func main() {
	// Initialize config
	err := util.InitializeConfig()
	if err != nil {
		log.Fatal("FATAL: Could not load config.json. Please check that the file '" + filepath.Join(glob.CONFIG_DIR, "config.json") + "' exists and is filled out properly. Error recieved: [" + err.Error() + "]")
	}
	log.Println("Config initialized")

	// Initialize database
	defer func() {
		err := database.CloseDatabase()
		if err != nil {
			log.Fatal("Unable to close the database connection")
		} else {
			log.Println("Closed database connection")
		}
	}()

	dbError := database.InitializeDatabase()
	if dbError != nil {
		log.Fatalf("Failed to initialize the database connection: %v", dbError)
	} else {
		log.Println("Database initialized")
	}

	/*// Specify the relative file name
	fileName := "tmp.png" // Adjust as necessary

	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Construct the relative path
	tempPicPath := filepath.Join(cwd, "internal", "database", fileName) //path to temporary picture

	profilePic, err := ioutil.ReadFile(tempPicPath)
	if err != nil {
		log.Println(glob.PictureNotFound)
	}

	test := models.Post{
		UserID:           7,
		BasisCanvasID:    1,
		Image:            profilePic,
		Caption:          "HALLABALLA",
		TimeSpentDrawing: 5,
	}*/

	// Set up routing
	handlers.ServeStaticContent()
	handlers.BindRoutes() // Bind all routes

	// Start the server
	log.Println("Service is listening om port: " + util.Config.Server.Port)
	//database.NewPost(test)
	log.Fatal(http.ListenAndServe(util.Config.Server.Host+":"+util.Config.Server.Port, nil))

}
