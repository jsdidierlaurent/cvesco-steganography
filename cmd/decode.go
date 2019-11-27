package cmd

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/auyer/steganography.v2"

	"github.com/spf13/cobra"
)

var decodeCmd = &cobra.Command{
	Use:   "decode [input-file-path]",
	Short: "Read secret message hide inside PNG/JPEG",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		msg, err := decodeMessage(args[0])
		if err != nil {
			fmt.Printf("Error while calling decode:\n%v\n", err)
			os.Exit(1)
		}

		fmt.Println(msg)
	},
}

func init() {
	rootCmd.AddCommand(decodeCmd)
}

func decodeMessage(inputImagePath string) (string, error) {
	img, err := OpenImageFromPath(inputImagePath)
	if err != nil {
		return "", err
	}

	sizeOfMessage := steganography.GetMessageSizeFromImage(img)
	// probably wrong image
	if sizeOfMessage > MaxMessageSize {
		return "", errors.New("their is no secret message hide in this file")
	}

	// Decode secret message inside new PNG
	msg := steganography.Decode(sizeOfMessage, img)

	return string(msg), nil
}
