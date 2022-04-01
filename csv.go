package main

import (
	"encoding/csv"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func read_dir() []fs.FileInfo {
	files, err := ioutil.ReadDir("data")
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func get_paths(files []fs.FileInfo) []string {
	var fs []string
	for _, f := range files {
		thepath, err := filepath.Abs(filepath.Dir(f.Name()))
		if err != nil {
			log.Fatal(err)
		}
		if strings.Contains(f.Name(), ".csv") {
			fs = append(fs, string(thepath)+string("/data/")+string(f.Name()))
		}
	}
	return fs
}

func read_csv(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	return records
}

func work_records(records [][]string) {
	sum := 0
	for _, r := range records {
		if r[12] == "member" {
			sum += 1
		}
	}
	result := fmt.Sprintf("the file had %v member rides in it", sum)
	fmt.Println(result)
}

func main() {
	start := time.Now()
	fs := read_dir()
	paths := get_paths(fs)
	fmt.Println(paths)
	for _, p := range paths {
		rcrds := read_csv(p)
		work_records(rcrds)
	}
	duration := time.Since(start)
	fmt.Println(duration)
}
