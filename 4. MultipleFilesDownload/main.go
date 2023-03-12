package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func downloadFile(URL string) ([]byte, error) {
	response, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, errors.New(response.Status)
	}

	var data bytes.Buffer

	_, err = io.Copy(&data, response.Body)
	if err != nil {
		return nil, err
	}

	return data.Bytes(), nil
}

func downloadMultipleFiles(urls []string) ([][]byte, error) {
	done := make(chan []byte, len(urls))
	errch := make(chan error, len(urls))

	for _, URL := range urls {
		go func(URL string) {
			byteData, err := downloadFile(URL)
			if err != nil {
				errch <- err
				done <- nil
				return
			}
			done <- byteData
			errch <- nil
		}(URL)
	}

	bytesArray := make([][]byte, 0)
	var errStr string

	for i := 0; i < len(urls); i++ {
		bytesArray = append(bytesArray, <-done)
		if err := <-errch; err != nil {
			errStr = errStr + " " + err.Error()
		}
	}

	var err error
	if errStr != "" {
		err = errors.New(errStr)
	}

	return bytesArray, err
}

func main() {
	// insert all the urls to the files in this slice
	urls := []string{
		"https://onlinejpgtools.com/images/examples-onlinejpgtools/sunflower.jpg",
		"https://upload.wikimedia.org/wikipedia/commons/thumb/3/3a/Cat03.jpg/1025px-Cat03.jpg",
		"https://ik.imagekit.io/ikmedia/backlit.jpg",
		"https://i.kym-cdn.com/photos/images/original/001/468/202/b02.jpg",
	}

	filesData, err := downloadMultipleFiles(urls)
	if err != nil {
		log.Fatal(err)
	}

	for idx, file := range filesData {
		err = ioutil.WriteFile("./"+strconv.Itoa(idx+1)+".jpg", file, 0700)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("All Files Downloaded Successfully!")
}
