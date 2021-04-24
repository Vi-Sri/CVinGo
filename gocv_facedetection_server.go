package main

import (
	"fmt"
	"html/template"
	"image"
	"image/color"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"gocv.io/x/gocv"
)

var (
	err      error
	webcam   *gocv.VideoCapture
	frame_id int
)

var buffer = make(map[int][]byte)
var frame []byte
var mutex = &sync.Mutex{}
var faceblur bool

func main() {
	host := "127.0.0.1:3000"
	faceblur, err = strconv.ParseBool(os.Args[1])
	if err != nil {
		fmt.Printf("Cant parse faceblur flag")
		return
	}
	// open webcam
	webcam, err = gocv.VideoCaptureDevice(0)


	if err != nil {
		fmt.Printf("Error opening capture device: \n")
		return
	}
	defer webcam.Close()

	// start capturing
	go detectFace()

	fmt.Println("Capturing. Open http://" + host)

	// start http server
	http.HandleFunc("/video", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "multipart/x-mixed-replace; boundary=frame")
		data := ""
		for {
			fmt.Println("Frame ID: ", frame_id)
			mutex.Lock()
			data = "--frame\r\n  Content-Type: image/jpeg\r\n\r\n" + string(frame) + "\r\n\r\n"
			mutex.Unlock()
			time.Sleep(33 * time.Millisecond)
			w.Write([]byte(data))
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("index.html")
		t.Execute(w, "index")
	})

	log.Fatal(http.ListenAndServe(host, nil))
}

func detectFace() {
	img := gocv.NewMat()
	defer img.Close()
	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("Device closed\n")
			return
		}
		if img.Empty() {
			continue
		}
		frame_id++

		blue_rect := color.RGBA{0, 0, 255, 0}
		classifier := gocv.NewCascadeClassifier()
		defer classifier.Close()
		if !classifier.Load("haarcascade_frontalface_default.xml") {
			fmt.Println("Error reading cascade file")
			return
		}
		rects := classifier.DetectMultiScale(img)
		for _, r := range rects {
			if faceblur {
				imgFace := img.Region(r)
				// blur face
				gocv.GaussianBlur(imgFace, &imgFace, image.Pt(75, 75), 0, 0, gocv.BorderDefault)
				imgFace.Close()
			} else {
				gocv.Rectangle(&img, r, blue_rect, 3)
			}
		}

		gocv.Resize(img, &img, image.Point{}, float64(0.5), float64(0.5), 0)
		frame, _ = gocv.IMEncode(".jpg", img)
	}
}