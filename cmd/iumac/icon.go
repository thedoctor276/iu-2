package main

import (
	"errors"
	"fmt"
	"image"
	"image/png"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/maxence-charriere/iu-log"
	"github.com/nfnt/resize"
)

func GenerateIcon(resourcePath string, icon string) error {
	iconPath := path.Join(resourcePath, icon)
	if _, err := os.Stat(iconPath); os.IsNotExist(err) {
		iulog.Warnf("icon %v does not exists")
		return nil
	}

	iconset := path.Join(resourcePath, iconToIconsetName(icon))
	if err := os.MkdirAll(iconset, os.ModeDir|0755); err != nil {
		return err
	}

	file, err := os.Open(iconPath)
	if err != nil {
		return err
	}
	defer file.Close()

	iconFilename := path.Join(iconset)

	img, err := png.Decode(file)
	if err != nil {
		return err
	}

	if err := createIcon(img, iconFilename, 512, 512, 2); err != nil {
		return err
	}

	if err := createIcon(img, iconFilename, 512, 512, 1); err != nil {
		return err
	}

	if err := createIcon(img, iconFilename, 256, 256, 2); err != nil {
		return err
	}

	if err := createIcon(img, iconFilename, 256, 256, 1); err != nil {
		return err
	}

	if err := createIcon(img, iconFilename, 128, 128, 2); err != nil {
		return err
	}

	if err := createIcon(img, iconFilename, 128, 128, 1); err != nil {
		return err
	}

	if err := createIcon(img, iconFilename, 32, 32, 2); err != nil {
		return err
	}

	if err := createIcon(img, iconFilename, 32, 32, 1); err != nil {
		return err
	}

	if err := createIcon(img, iconFilename, 16, 16, 2); err != nil {
		return err
	}

	if err := createIcon(img, iconFilename, 16, 16, 1); err != nil {
		return err
	}

	if err := RunCmd("iconutil", "-c", "icns", iconFilename); err != nil {
		return nil
	}

	return os.RemoveAll(iconFilename)
}

func createIcon(img image.Image, filename string, width uint, height uint, mult uint) error {
	if mult == 0 {
		return errors.New("mult can't be 0")
	}

	resized := resize.Resize(width*mult, height*mult, img, resize.Lanczos3)

	if mult == 1 {
		filename = path.Join(filename, fmt.Sprintf("icon_%vx%v.png", width, height))
	} else {
		filename = path.Join(filename, fmt.Sprintf("icon_%vx%v@%vx.png", width, height, mult))
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	return png.Encode(f, resized)
}

func iconId(icon string) string {
	iconBase := filepath.Base(icon)
	return strings.TrimSuffix(iconBase, filepath.Ext(iconBase))
}

func iconToIconsetName(icon string) string {
	return iconId(icon) + ".iconset"
}
