package main

import (
	"log"

	"github.com/signintech/gopdf"
)

func main() {

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{
		PageSize: *gopdf.PageSizeA4,
	})
	pdf.AddPage()

	err := pdf.AddTTFFont("ipaexg", "./IPAexfont00401/ipaexg.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}

	err = pdf.SetFont("ipaexg", "", 24)
	if err != nil {
		log.Print(err.Error())
		return
	}

	pdf.Cell(nil, "こんにちは！　なんとPDFを作成できました！")
	pdf.WritePdf("./pdf/hello.pdf")
}
