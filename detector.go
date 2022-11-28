package main

import (
	"fmt"
	"image"
	"image/color"
	"time"

	"gocv.io/x/gocv"
)

type Result struct {
	success bool
}

func runDetector(feed chan<- bool) {

	time.Sleep(time.Millisecond * 2000)

	webcam, _ := gocv.VideoCaptureDevice(0)
	window := gocv.NewWindow("Hello")
	img := gocv.NewMat()
	gray := gocv.NewMat()

	var red = color.RGBA{255, 0, 0, 0}
	var green = color.RGBA{43, 212, 44, 0}

	low := gocv.NewScalar(0, 0, 255, 0)
	high := gocv.NewScalar(255, 255, 255, 0)

	mask := gocv.NewMat()

	targetContour := gocv.NewPointVector()

	for {
		webcam.Read(&img)

		gocv.InRangeWithScalar(img, low, high, &mask)

		gocv.CvtColor(img, &gray, 6)
		gocv.Threshold(gray, &gray, 170, 255, 0)

		contours := gocv.FindContours(gray, 3, 4)
		for i := 0; i < contours.Size(); i++ {
			contour := contours.At(i)
			approx := gocv.ApproxPolyDP(contour, 0.01*gocv.ArcLength(contour, true), true)

			for _, p := range approx.ToPoints() {
				gocv.Circle(&img, p, 2, red, 2)
			}

			area := gocv.ContourArea(approx)
			if area > 10000 {
				targetContour = approx
			}
		}

		ptsv := gocv.NewPointsVector()
		ptsv.Append(targetContour)
		gocv.Polylines(&img, ptsv, true, red, 2)
		ptsv.Close()

		_, _, _, loc := gocv.MinMaxLoc(mask)
		if loc.X != 0 && loc.Y != 0 {

			gocv.Circle(&gray, loc, 10, green, 2)

			result := gocv.PointPolygonTest(targetContour, loc, false) > 0
			select {
			case feed <- result:
				fmt.Println(result)
			default:
			}
		}

		window.IMShow(img)
		window.WaitKey(1)
	}
}

func rectContains(p *image.Point, min *image.Point, max *image.Point) bool {
	return p.X > min.X && p.X < max.X && p.Y > min.Y && p.Y < max.Y
}
