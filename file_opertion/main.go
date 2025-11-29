package main

import (
	"errors"
	"fmt"
	"os"
)

func fileExists(path string) (bool, error) {
	_, err := os.Lstat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}



func CreateFile(path string)error{
	file,err:=os.Create(path)
	if err != nil{
		return err
	}
    defer file.Close()

	
	return nil

}

func main() {
	// file exist
	path := "sample.txt"
	filexist,_:=fileExists(path)
	fmt.Println("file exist:",filexist)

}
