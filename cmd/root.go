package cmd

import (
	"bufio"
	"fmt"
	"image"
	"os"

	"github.com/spf13/cobra"
)

const MaxMessageSize = 1000

var rootCmd = &cobra.Command{
	Use: "steganography.exe",
	Long: `Well done darling, if you read that you're good!
This tool will serve you throughout your journey!

Good luck, I love you 😍`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func OpenImageFromPath(filename string) (image.Image, error) {
	inFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer inFile.Close()
	reader := bufio.NewReader(inFile)
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}
	return img, nil
}
