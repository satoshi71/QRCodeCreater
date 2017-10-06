package main

import (
	"fmt"
	"image/png"
	"os"
	"io/ioutil"
	"bufio"
	"encoding/xml"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

type Conf struct {
	Out  string
	ListFile string
	Size int
}

var conf Conf
var codeList []string

func main() {
	//パラメーターをxmlファイルから取得する
	getConf()
	//QRコードにする文字のリストを読み込む
	getList()

	createCode()

	//fmt.Println("Filename: " + conf.ListFile)
	fmt.Println(len(codeList))
}

func getConf(){
	data, _ := ioutil.ReadFile("conf.xml")
	err := xml.Unmarshal(data, &conf)
	if err != nil { panic(err) }
}

func getList() {
	fp, err := os.Open(conf.ListFile)
	if err != nil { panic(err) }
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		codeList = append(codeList, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func createCode(){
	for _, code := range codeList {
		if code=="" { continue }
		qrCode, _ := qr.Encode(code, qr.M, qr.Auto)
		qrCode, _ = barcode.Scale(qrCode, conf.Size, conf.Size)
		file, _ := os.Create(conf.Out + code + ".png")
		defer file.Close()
		png.Encode(file, qrCode)
	}
}
//github.com/boombuler/barcode