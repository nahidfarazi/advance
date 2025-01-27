package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)
func Check(e error){
	if e!=nil {
		fmt.Println(e)
		return
	}
}
type Data struct{
	URL string
	fileName string
}
func (d *Data) download()error{
	res, err := http.Get(d.URL)
	Check(err)
	defer res.Body.Close()

	if res.StatusCode != 200{
		return errors.New("received non 200 response code")
	}
	file,err := os.Create(d.fileName)
	Check(err)
	defer file.Close()

	_,err = io.Copy(file, res.Body)
	Check(err)
	return nil
}

func main() {
	//var url,filename string
	// fmt.Printf("enter your url: ")
	// fmt.Scanf("%d",&url)
	// fmt.Printf("enter your file name: ")
	// fmt.Scanf("%d",&filename)
	data := &Data{
		URL:"https://www.facebook.com/watch/?v=1525382954801931",
		fileName: "sol4.mp4",
	}
	data.download()
}