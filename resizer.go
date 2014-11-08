package main

import (
    "github.com/nfnt/resize"
    "github.com/golang/glog"
    "image"
    "image/png"
    "image/jpeg"
    "os"
    "flag"
    //"path"
)

var flagInFile = flag.String("infile", "", "")
var flagOutFile = flag.String("outfile", "resizer.out", "")

var flagWidth = flag.Int("width", 100, "")
var flagHeight = flag.Int("height", 100, "")

func init() {
    flag.Parse()
}

func decode(filename string) (image.Image, string) {
    //ext := path.Ext(filename)
    //glog.Info(ext)
    file, err := os.Open(filename)
    if err != nil {
        glog.Fatal(err)
    }
    defer file.Close()

    img, format, err := image.Decode(file)
    glog.Info(format)
    if err != nil {
        glog.Fatal(err)
    }
    return img, format
}

func encode(filename string, format string, img image.Image) {
    out, err := os.Create(filename)
    if err != nil {
        glog.Fatal(err)
    }
    defer out.Close()
    switch format {
    case "png":
        jpeg.Encode(out, img, nil)
    case "jpeg":
        png.Encode(out, img)
    }
}

func main() {
    img, format:= decode(*flagInFile)
    // decode jpeg into image.Image
    m := resize.Resize(uint(*flagWidth), uint(*flagHeight), img, resize.Lanczos3)
    encode(*flagOutFile, format, m)
}
