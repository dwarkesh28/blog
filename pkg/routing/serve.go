package routing

import (
	"blog/pkg/config"
	"fmt"
	"log"
)

func Serve() {
	r := GetRouter()
	configs := config.Get()
	err := r.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port)) //listen and serve on 0.0.0.0:8080 for windows "loclahost:8000"

	if err != nil {
		log.Fatal("error in routing")
		return
	}
}
