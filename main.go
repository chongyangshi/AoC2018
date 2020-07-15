package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/chongyangshi/AoC2018/solutions"
)

func main() {
	workingDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal("Cannot get current working directory.")
	}

	runSolution := flag.String("run", "", "Specify which day of solutions to be run.")
	flag.Parse()

	inputPath := path.Join(workingDirectory, "inputs", "day"+*runSolution+".txt")
	if _, err = os.Stat(inputPath); err != nil {
		log.Fatal(fmt.Sprintf("Cannot get corresponding input (day%s.txt) in %s", *runSolution, inputPath))
	}
	input, err := ioutil.ReadFile(inputPath)
	if err != nil {
		log.Fatal(fmt.Sprintf("Cannot read corresponding input (day%s.txt) in %s", *runSolution, inputPath))
	}
	inputString := string(input)

	fmt.Printf("Executing solution for Day %s \r\n", *runSolution)
	results, execTime := solutions.RunSolution(*runSolution, inputString)
	fmt.Printf("Completed in %v seconds, output: \n%s\r\n", execTime.Seconds(), results)

	return
}
