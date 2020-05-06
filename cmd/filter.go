package main

import (
	"path/filepath"

	"github.com/qt-luigi/imageproc"
)

func rotate180(file string) ([]string, error) {
	d, fn := filepath.Split(file)
	fs := []string{fn}
	fns, err := imageproc.Rotate180(d, fs)
	if err != nil {
		return nil, err
	}
	return fns, nil
}

func decreaseBrightness(file string, steps int) ([]string, error) {
	d, fn := filepath.Split(file)
	fns := make([]string, steps)
	for i := 0; i < steps; i++ {
		fns[i] = fn
	}
	fns, err := imageproc.DecreaseBrightness(d, fns)
	if err != nil {
		return nil, err
	}
	return fns, nil
}
