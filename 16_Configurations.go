package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/caarlos0/env"
)

/*
- идентификаторы подключения к ресурсам
- креды
- значения зависимые от среды развертывания
- конфа: файлы || переменные окружения
- confita lib: for multiple backends
- viper

logging:
- zap is kinda cool
*/

/* caarlos PRODUCTION=true HOSTS="host1:host2:host3" go run ./16_Configurations.go */
type config struct {
	Home         string        `env:"HOME"`
	Port         int           `env:"PORT" envDefault:"3000"`
	IsProduction bool          `env:"PRODUCTION"`
	Hosts        []string      `env:"HOSTS" envSeparator:":"`
	Duration     time.Duration `env:"DURATION"`
	TempFolder   string        `env:"TEMP_FOLDER" envDefault:"${HOME}/tmp" envExpand:"true"`
}

func init() {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("init started")
}

func main() {

	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("%+v\n", err)
	}
	fmt.Println(cfg)

	// // Println writes to the standard logger.
	// log.Println("main started")
	// // Fatalln is Println() followed by a call to os.Exit(1)
	// log.Fatalln("fatal message")
	// // Panicln is Println() followed by a call to panic()
	// log.Panicln("panic message")

	f, err := os.OpenFile("info.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.Println("This is a test log entry")

}
