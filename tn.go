package main

import (
    "bufio"
    "fmt"
    "io"
    "log"
    "os"
    "strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func normalize(line string) string {
// step 1 - Trim trailing spaces
    s := strings.TrimSpace(line)
// step 2 - Convert TAB to spaces

    /* implementing  */

    return s
}

func main() {

	if(len(os.Args) != 2) {
		fmt.Printf("Usage: %s <file name>\n",os.Args[0])
		return
	}
	filename := os.Args[1]

	f, err := os.Open(filename)
	check(err)

	bf := bufio.NewReader(f)
	for {
		switch line, err := bf.ReadString('\n'); err {
		case nil:
			tn := normalize(line)
			fmt.Printf("%s\n", tn)
		case io.EOF:
			if line > "" {
			// last line of file missing \n, but still valid
				fmt.Println(line)
			}
			return
		default:
		log.Fatal(err)
        }
    }
    f.Close()
}

