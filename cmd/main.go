package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/qt-luigi/imageproc"
)

const (
	usageMsg = `imageproc processes an image file to step by step decreasing brightness and upside down.

Usage:

	imageproc -f <pngfile> -upsidedown
	imageproc -f <pngfile> -dbs <steps>
	imageproc -f <pngfile> -dbs <steps> -upsidedown

Arguments are:

	pngfile		base .png file
	steps		decreasing brightness steps
	-upsidedown	upside down pngfile
`
)

var (
	pngfile    string
	steps      int
	upsidedown bool
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usageMsg)
	}
	flag.StringVar(&pngfile, "f", "", "base .png file")
	flag.IntVar(&steps, "dbs", 0, "decreasing brightness steps")
	flag.BoolVar(&upsidedown, "upsidedown", false, "upside down pngfile")
}

func main() {
	flag.Parse()

	if pngfile == "" {
		flag.Usage()
		os.Exit(1)
	}

	if err := validateFile(pngfile, ".png"); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if steps == 0 {
		if upsidedown {
			if _, err := rotate180(pngfile); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		} else {
			flag.Usage()
			os.Exit(1)
		}
	} else if steps >= 1 && steps <= 100 {
		files, err := decreaseBrightness(pngfile, steps)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if upsidedown {
			dir, _ := filepath.Split(pngfile)
			if _, err = imageproc.Rotate180(dir, files); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		}
	} else {
		fmt.Fprintln(os.Stderr, "steps is out of range [1-100]")
		os.Exit(1)
	}
}

func validateFile(file, ext string) error {
	if fi, err := os.Stat(file); err != nil {
		return err
	} else if fi.IsDir() {
		return fmt.Errorf("%s is not a file", file)
	} else if strings.ToLower(filepath.Ext(file)) != ext {
		return fmt.Errorf("%s is not a %s file", file, ext)
	}
	return nil
}
