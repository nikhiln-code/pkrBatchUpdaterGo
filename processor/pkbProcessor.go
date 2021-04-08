package processor

import (
	"bufio"
	"fmt"
	"strconv"

	//"net/http"
	log "github.com/sirupsen/logrus"
	"nordea/npl/norway/pkrPublisher/v2/helper"
	"os"
)

func ProcessEPK() {
	helper.InitiateAndCreateLogfile("epk")
	log.Info("processing the opa file")
	inputFile, err := os.Open("test/resources/epk-input.txt")
	if err != nil {
		log.Error("error opening the input file")
		fmt.Println("error opening the input file")
	}
	inputFileSanner := bufio.NewScanner(inputFile)
	inputFileSanner.Split(bufio.ScanLines)

	for inputFileSanner.Scan() {
		//fmt.Println("input file:"+ inputFileSanner.Text() +" valid json: " + strconv.FormatBool(helper.CheckValidJson(inputFileSanner.Text())))
		log.Info("input file:" + inputFileSanner.Text() + " valid json: " + strconv.FormatBool(helper.CheckValidJson(inputFileSanner.Text())))
		//http.Patch()
	}

	inputFile.Close()
	helper.ExitLogger()
}
