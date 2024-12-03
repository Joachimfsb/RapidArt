package glob

// File system
const (
	WEB_DIR    = "web/"
	RES_DIR    = WEB_DIR + "res/"
	HTML_DIR   = WEB_DIR + "html/"
	CONFIG_DIR = "configs/"
)

// error messages
const (
	InvalidNameOrPass    = "invalid username or password"     //tested
	UserNotFound         = "user not found"                   //tested
	PictureNotFound      = "picture not found"                //tested
	SqlAttempt           = "sql injection attempt discovered" //tested
	UsernameAlreadyExist = "username already exists"          //tested
	EmailAlreadyExist    = "email already registered"         //tested
	ScanFailed           = "scanning from database failed"    //tested
	NoGallery            = "no rows in database"
)

const NumberOfReportsBeforeDeactivatePost = 10
const SessionExpirationDays = 90
const DefaultProfilePicturePath = "/res/img/default-profile-img.png"
