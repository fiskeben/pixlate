# Pixlate

Pixelates photos.
Written in Go.

The program will take an image file and a block size as input
and write a new image file where blocks of pixels are averaged.

For example, passing a block size of 10 will create a new image
where the color of pixels in 10x10 blocks is averaged,
creating a pixelating effect.

## Building

`go build .`

## Usage

`pixlate <infile> <blocksize> [-output <outfile>]`

## Limitations

Currently only supports jpeg image format.
