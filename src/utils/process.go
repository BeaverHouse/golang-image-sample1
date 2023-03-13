package utils

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/noelyahan/impexp"
	"github.com/noelyahan/mergi"
)

func ChangeImage(inDir string, outDir string) {
	files, err := ioutil.ReadDir(inDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".png") || strings.HasSuffix(f.Name(), ".jpg") {
			imgFile, err := os.Open(inDir + "/" + f.Name())
			if err != nil {
				log.Fatal(err)
				return
			}

			var inputImg image.Image
			if strings.HasSuffix(f.Name(), ".png") {
				inputImg, err = png.Decode(imgFile)
			} else if strings.HasSuffix(f.Name(), ".jpg") {
				inputImg, err = jpeg.Decode(imgFile)
			}
			if err != nil {
				fmt.Println("??")
				log.Fatal(err)
				return
			}

			logoImg, err := mergi.Import(impexp.NewURLImporter("https://upload.wikimedia.org/wikipedia/commons/thumb/2/2f/Google_2015_logo.svg/368px-Google_2015_logo.svg.png"))
			if err != nil {
				log.Fatal(err)
				return
			}

			res, _ := mergi.Watermark(logoImg, inputImg, image.Pt(0, 200))

			out, err := os.Create(outDir + "/" + f.Name())
			if err != nil {
				log.Fatal(err)
				return
			}
			if strings.HasSuffix(f.Name(), ".png") {
				err = png.Encode(out, res)
			} else if strings.HasSuffix(f.Name(), ".jpg") {
				err = jpeg.Encode(out, res, &jpeg.Options{Quality: 80})
			}
			if err != nil {
				log.Fatal(err)
				return
			}

		}
	}

	fmt.Println("Converted")
}
