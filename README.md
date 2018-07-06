# Pixlate

Pixelates photos.
Written in Go.

The program will take an image file and a block size as input
and write a new image file where blocks of pixels are averaged.

For example, passing a block size of 10 will create a new image
where the color of pixels in 10x10 blocks is averaged,
creating a pixelating effect.

Works with jpeg, gif, and png.

## Building

`go build .`

## Usage

`pixlate <infile> <blocksize> [-output <outfile>]`

## Limitations

Animated gifs will only have the first frame pixelated
and the rest of the frames removed.
