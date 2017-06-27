package helper

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"os"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func PHA(m image.Image) string {
	// Step 1: resize picture to 8*8.
	//m = resize.Resize(8, 8, m, resize.NearestNeighbor)
	m = Resize(m, m.Bounds(), 8, 8)

	// Step 2: grayscale picture.
	gray := grayscaleImg(m)

	// Step 3: calculate average value.
	avg := calAvgValue(gray)

	// Step 4: get fingerprint.
	fg := getFingerprint(avg, gray)

	return string(fg)
}

// grayscaleImg converts picture to grayscale.
func grayscaleImg(src image.Image) []byte {
	// Create a new grayscale image
	bounds := src.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	gray := make([]byte, w*h)
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			r, g, b, _ := src.At(x, y).RGBA()
			gray[x+y*8] = byte((r*30 + g*59 + b*11) / 100)
		}
	}
	return gray
}

// calAvgValue returns average value of color of picture.
func calAvgValue(gray []byte) byte {
	sum := 0
	for _, v := range gray {
		sum += int(v)
	}
	return byte(sum / len(gray))
}

// getFingerprint returns fingerprint of a picture.
func getFingerprint(avg byte, gray []byte) string {
	var buf bytes.Buffer
	for _, v := range gray {
		if avg >= v {
			buf.WriteByte('1')
		} else {
			buf.WriteByte('0')
		}
	}
	return buf.String()
}

// Pha Compare
func CompareDiff(fg1, fg2 string) int {
	diff := 0
	fbyte := []byte(fg1)
	fbyte2 := []byte(fg2)
	for i, v := range fbyte {
		if fbyte2[i] != v {
			diff++
		}
	}
	return diff
}

//PHA算法  获取图像指纹
func GetImagePha(path string) (string, error) {

	if infile, err := os.Open(path); err != nil {
		return "", err
	} else {

		// Decode picture.
		if srcImg, _, err := image.Decode(infile); err != nil {
			fmt.Println("Decode picture:", err)
		} else {

			return PHA(srcImg), err

		}

	}
	return "", errors.New("获取图片PHA值出现错误")

}

//指纹比较
func PhaCompare(path1 string, path2 string) (int, error) {

	if fg1, err := GetImagePha(path1); err != nil {
		return -1, err
	} else {
		if fg2, err := GetImagePha(path2); err != nil {
			return -1, err
		} else {

			return CompareDiff(fg1, fg2), err
		}
	}

}
