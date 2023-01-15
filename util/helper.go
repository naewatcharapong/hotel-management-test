package util

import (
	"fmt"
	"os"
	"strings"
)

func ArrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

func WriteOutputFile(content string) {
	file, errToCreate := os.Create("./data/output.txt")
	if errToCreate != nil {
		panic("Can't to create the output file.")
	}
	defer file.Close()
	_, errToWriteFile := file.WriteString(content)
	if errToWriteFile != nil {
		panic("Can't to write the output file.")
	}
	fmt.Println("Create output file successfully.")
}
