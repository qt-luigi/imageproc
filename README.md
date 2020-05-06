# imageproc

imageproc processes an image file to step by step decreasing brightness and upside down.

## Installation

When you have installed the Go, Please execute following `go get` command:

```sh
go get -u github.com/qt-luigi/imageproc
```

## Usage

```sh
$ imageproc
imageproc processes an image file to step by step decreasing brightness and upside down.

Usage:

	imageproc -f <pngfile> -upsidedown
	imageproc -f <pngfile> -dbs <steps>
	imageproc -f <pngfile> -dbs <steps> -upsidedown

Arguments are:

	pngfile		base .png file
	steps		decreasing brightness steps
	-upsidedown	upside down pngfile
```

## License

MIT

## Author

Ryuji Iwata

## Note

This tool is mainly using by myself. :-)
