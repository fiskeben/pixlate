package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
)

func main() {
	writeto := flag.String("output", "", "path and name of file to write the output to")
	flag.Parse()

	filename := flag.Arg(0)
	size := flag.Arg(1)
	if filename == "" || size == "" {
		fmt.Println("Usage: pixlate <path> <size>")
		os.Exit(1)
	}

	s, err := strconv.Atoi(size)
	if err != nil {
		fmt.Printf("%s is not a number\n", size)
		os.Exit(1)
	}
	blocksize := uint32(s * s)

	outputFilename := *writeto
	if outputFilename == "" {
		outputFilename = makeOutputFilename(filename)
	}

	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("failed to open file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	outf, err := os.Create(outputFilename)
	if err != nil {
		fmt.Printf("failed to create a file for writing: %v\n", err)
		os.Exit(1)
	}
	defer outf.Close()

	img, imagetype, err := image.Decode(f)
	if err != nil {
		fmt.Printf("unable to decode image: %v\n", err)
		os.Exit(1)
	}

	bounds := img.Bounds()
	destination := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y += s {
		for x := bounds.Min.X; x < bounds.Max.X; x += s {
			r, g, b, a := uint32(0), uint32(0), uint32(0), uint32(0)
			for y1 := y; y1 < y+s; y1++ {
				for x1 := x; x1 < x+s; x1++ {
					c := img.At(x1, y1)
					r1, g1, b1, a1 := c.RGBA()
					r += uint32(r1) / blocksize
					g += uint32(g1) / blocksize
					b += uint32(b1) / blocksize
					a += uint32(a1) / blocksize
				}
			}
			avg := color.RGBA{
				R: uint8(r >> 8),
				G: uint8(g >> 8),
				B: uint8(b >> 8),
				A: uint8(a >> 8),
			}
			for y1 := y; y1 < y+s; y1++ {
				for x1 := x; x1 < x+s; x1++ {
					destination.Set(x1, y1, avg)
				}
			}
		}
	}

	if err = encode(outf, destination, imagetype); err != nil {
		fmt.Printf("failed to encode result: %v\n", err)
		os.Exit(1)
	}
}

func makeOutputFilename(filename string) string {
	outfile := path.Base(filename)
	parts := strings.Split(outfile, ".")
	if len(parts) == 1 {
		outfile = outfile + ".pixlated"
	} else {
		extension := fmt.Sprintf(".%s", parts[len(parts)-1])
		name := strings.Join(parts[:len(parts)-1], ".")
		outfile = fmt.Sprintf("%s.pixlated%s", name, extension)
	}
	dir := path.Dir(filename)
	return path.Join(dir, outfile)
}

func encode(w io.Writer, i image.Image, imagetype string) error {
	switch imagetype {
	case "jpg", "jpeg":
		return jpeg.Encode(w, i, nil)
	case "png":
		return png.Encode(w, i)
	case "gif":
		return gif.Encode(w, i, nil)
	}
	return fmt.Errorf("unknown image type %s", imagetype)
}
