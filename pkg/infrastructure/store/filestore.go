package pkg

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	tconfig "github.com/de-bugsBunny/go-ndm/tool/config"
)

type FileStorage struct {
	filePath string
	fileName string
}

func NewFileStorage(filePath string, fileName string) FileStorage {
	return FileStorage{
		filePath: filePath,
		fileName: fileName,
	}
}

func (f FileStorage) Store(data string) error {
	file, err := os.Create(f.filePath + f.fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(data)
	return nil

}

func (f FileStorage) Clean(interval int) error {
	files, err := ioutil.ReadDir(f.filePath)
	if err != nil {
		log.Fatal(err)
	}

	if tconfig.GetAppConfigInstance().Debug {
		fmt.Println(f.filePath)
	}
	for _, file := range files {
		if tconfig.GetAppConfigInstance().Debug {
			fmt.Println(file.Name())
		}
		if time.Now().Sub(file.ModTime()).Minutes() > float64(interval) {
			if tconfig.GetAppConfigInstance().Debug {
				fmt.Println("Removing file :" + file.Name())
			}
			err := os.Remove(f.filePath + file.Name())
			if err != nil {
				if tconfig.GetAppConfigInstance().Debug {
					fmt.Println(err)
				}
			}
		}
	}
	return nil

}
