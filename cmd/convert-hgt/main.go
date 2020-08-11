package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Hades32/height-map"

	"github.com/itzg/go-flagsfiller"
)

type Config struct {
	In           string `usage:"The .hgt file to convert"`
	Out          string `usage:"The image file to create (only supports png at the moment)"`
	DebugEnabled bool   `default:"true" usage:"Show debug information"`
}

func main() {
	config := parseArguments()

	heightData, err := os.Open(config.In)
	if err != nil {
		fmt.Println("Couldn't open input file", err)
		os.Exit(1)
	}
	imageFile, err := os.Create(config.Out)
	if err != nil {
		fmt.Println("Couldn't create output file", err)
		os.Exit(1)
	}
	err = heightmap.Convert(heightData, imageFile)
	if err != nil {
		fmt.Println("Couldn't convert file", err)
		os.Exit(1)
	}
}

func parseArguments() Config {
	var config Config
	filler := flagsfiller.New()
	err := filler.Fill(flag.CommandLine, &config)
	if err != nil {
		fmt.Println("Couldn't parse arguments", err)
		os.Exit(1)
	}
	flag.Parse()

	if config.In == "" || config.Out == "" {
		fmt.Println("A tool to convert SRTM height map data to grayscale pictures.")
		fmt.Println("Data can be retrieved from https://dds.cr.usgs.gov/srtm/version2_1/SRTM3/")
		flag.Usage()
		os.Exit(1)
	}
	return config
}
