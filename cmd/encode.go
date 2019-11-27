package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/auyer/steganography.v2"
)

var encodeCmd = &cobra.Command{
	Use:   "encode [input-file-path] [output-file-path] [message]",
	Short: "Hide secret message inside PNG",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		err := encodeMessage(args[0], args[1], args[2])
		if err != nil {
			fmt.Printf("Error while calling encode:\n%v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(encodeCmd)
}

func encodeMessage(inputImagePath, outputImagePath, message string) error {
	if len(message) > MaxMessageSize {
		return errors.New("the message is too long")
	}

	img, err := OpenImageFromPath(inputImagePath)
	if err != nil {
		return err
	}

	// Encode secret message inside new PNG
	w := new(bytes.Buffer)
	err = steganography.Encode(w, img, []byte(message))
	if err != nil {
		return err
	}

	// Create new file and write image inside
	outFile, err := os.Create(outputImagePath)
	if err != nil {
		return err
	}
	_, err = w.WriteTo(outFile)
	if err != nil {
		return err
	}
	_ = outFile.Close()

	return nil
}
