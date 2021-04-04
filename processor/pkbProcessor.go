package processor

import (
	"bufio"
	"fmt"
	"net/http"
	"nordea/npl/norway/pkrPublisher/v2/helper"
	log "github.com/sirupsen/logrus"
	"os"
)

func ProcessEPK(){
	helper.InitiateLogger("epk");
	log.Info("processing the opa file ");
	file, err := os.Open("epk-input.txt")
	if(err!= nil){
		log.Error("error opening the file")
		fmt.Println("error opening the file")
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan(){
		http.Patch()
	}
	helper.ExitLogger()
}
