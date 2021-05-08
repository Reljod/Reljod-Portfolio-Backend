package environ

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

const projectDirName = "Reljod-Portfolio-Backend"

type Environ struct {
	Port     string
	LogLevel string
}

var DatabaseUrl string

func (env *Environ) Setup() {

	re := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

	err := godotenv.Load(os.ExpandEnv(string(rootPath) + "/config.env"))

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	env.Port = os.Getenv("PORT")
	env.LogLevel = os.Getenv("LOG")
	DatabaseUrl = os.Getenv("DATABASE_URL")
}
