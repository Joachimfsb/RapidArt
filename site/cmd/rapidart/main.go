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

	/*test := models.RapidUser{ //test to add user
		Username:     "Test",
		Email:        "test@test.no",
		Displayname:  "TheTester123",
		Password:     "test",
		CreationTime: time.Now().String(),
		Bio:          "TesterTesterTester",
	}*/

	//test, err := database.UserById(7)    //test for fetching from db
	//test, err := database.UserByEmail("test@test.no") //test for fetching from db
	/*atest := models.UserAuthentication{ //test of authentication
		Email:    "test@test.no",
		Password: "test",
	}*/

	// Set up routing
	handlers.ServeStaticContent()
	handlers.BindRoutes() // Bind all routes

	// Start the server
	log.Println("Service is listening om port: " + util.Config.Server.Port)
	//database.AddUser(test) //test to add user
	/*test, err := database.UserLogin(atest)
	log.Println(test.Username + " " + test.Email + " " + test.Displayname + " " + test.Password + " " + test.PasswordSalt + " " + test.Role + " " + test.Bio) //test after fetching from db */
	log.Fatal(http.ListenAndServe(util.Config.Server.Host+":"+util.Config.Server.Port, nil))

}
