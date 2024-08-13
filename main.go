package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "reverse-color [file/folder...]",
		Short: "Reverse video colors",
		Args:  cobra.MinimumNArgs(1),
		Run:   reverseColors,
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func reverseColors(cmd *cobra.Command, args []string) {
	var wg sync.WaitGroup
	for _, arg := range args {
		wg.Add(1)
		go func(path string) {
			defer wg.Done()
			info, err := os.Stat(path)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}
			if info.IsDir() {
				processFolder(path)
			} else {
				processFile(path)
			}
		}(arg)
	}
	wg.Wait()
}

func processFolder(folderPath string) {
	outputFolder := folderPath + "_reversed"
	err := os.MkdirAll(outputFolder, os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating output folder: %v\n", err)
		return
	}

	err = filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && isVideoFile(path) {
			relPath, err := filepath.Rel(folderPath, path)
			if err != nil {
				fmt.Printf("Error getting relative path: %v\n", err)
				return nil
			}
			outputPath := filepath.Join(outputFolder, relPath)
			os.MkdirAll(filepath.Dir(outputPath), os.ModePerm)
			processFile(path, outputPath)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error processing folder: %v\n", err)
	}
}

func processFile(inputFile string, outputFile ...string) {
	var output string
	if len(outputFile) > 0 {
		output = outputFile[0]
	} else {
		output = getOutputFilename(inputFile)
	}

	ffmpegCmd := exec.Command("ffmpeg",
		"-y",
		"-i", inputFile,
		"-vf", "lutrgb='r=255*0.9-val*0.9+25:g=255*0.9-val*0.9+25:b=255*0.9-val*0.9+25',eq=brightness=-0.05:contrast=1.1",
		"-c:v", "h264_videotoolbox",
		"-b:v", "5M",
		"-c:a", "copy",
		output,
	)

	err := ffmpegCmd.Run()
	if err != nil {
		fmt.Printf("Error processing %s: %v\n", inputFile, err)
	} else {
		fmt.Printf("Processed: %s -> %s\n", inputFile, output)
	}
}

func getOutputFilename(inputFile string) string {
	dir := filepath.Dir(inputFile)
	filename := filepath.Base(inputFile)
	ext := filepath.Ext(filename)
	name := filename[:len(filename)-len(ext)]
	return filepath.Join(dir, name+"_reversed"+ext)
}

func isVideoFile(filename string) bool {
	ext := filepath.Ext(filename)
	switch ext {
	case ".mp4", ".avi", ".mov", ".mkv", ".flv", ".wmv":
		return true
	}
	return false
}
