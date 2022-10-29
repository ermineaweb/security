package infra

import (
	"log"
	"os"
)

type FileRepository struct {
	Filename  string
	Directory string
}

func NewFileRepository(filename string, directory string) FileRepository {
	return FileRepository{Filename: filename, Directory: directory}
}

func (r FileRepository) Save(data string) {
	path := "results/" + r.Directory
	os.MkdirAll(path, os.ModePerm)

	file, err := os.Create(path + "/" + r.Filename)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		log.Fatal(err)
	}
}
