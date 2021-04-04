package main

import(
	"nordea/npl/norway/pkrPublisher/v2/cmd"
	_ "nordea/npl/norway/pkrPublisher/v2/helper"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.Info("Executing the pkrPublisher")
	cmd.Execute()

}
