package main
 
import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"bytes"
	"io"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func lineCounter(r io.Reader) (int, error) {
    buf := make([]byte, 32*1024)
    count := 0
    lineSep := []byte{'\n'}

    for {
        c, err := r.Read(buf)
        count += bytes.Count(buf[:c], lineSep)

        switch {
        case err == io.EOF:
            return count, nil

        case err != nil:
            return count, err
        }
    }
}

func parseFile(input string, output string, Count int){
	fmt.Printf("LINE COUNT AT  %v \n", Count)
	inputFile, err1 := os.Open(input)
	check(err1)
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	outputFile, err3 := os.Create(output)
	w := bufio.NewWriter(outputFile)
	check(err3)
	n := 0
	lines := 0
	linebuff := 0
	lineid := 1
	fastarec := 0
	for scanner.Scan() {
		n ++
		lines ++
		if lines == 1000000 {
			linebuff ++ 
			fmt.Printf("PROCESSED %v  MILLION READS\n", linebuff)
			lines = 0
		}
		if lineid == 4 {
			lineid = 1

		} else if lineid == 3 {
			lineid ++
		} else if lineid == 2 {
			lineid ++
			//fmt.Println(scanner.Text())
			fasta := scanner.Text()
			w.WriteString(fasta+"\n")
			fastarec ++
		} else {
			//fmt.Println(scanner.Text())
			lineid ++
			fastaID := scanner.Text()
			fastaID = strings.Replace(fastaID, "@", ">", 1)
			w.WriteString(fastaID+"\n")
		}
	}
	fmt.Printf("Processed lines %v \n", n)
	fmt.Printf("Processed fasta records %v \n", fastarec)
	w.Flush()
}

func main() {
	// Open an input file, exit on error.
	input:= os.Args[1]
	output:= os.Args[2]
	inputFile, err1 := os.Open(input)
	check(err1)
	defer inputFile.Close()
	Count, err2:= lineCounter(inputFile)
	check(err2)
	parseFile(input, output, Count)
	
}