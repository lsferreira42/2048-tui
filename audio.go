//go:build audio
// +build audio

package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var (
	moveSound   string
	endingSound string
)

func initAudio() error {
	var err error
	moveSound, err = extractSound("move_sound.mp3")
	if err != nil {
		return fmt.Errorf("error extracting move sound: %v", err)
	}

	endingSound, err = extractSound("ending.mp3")
	if err != nil {
		return fmt.Errorf("error extracting ending sound: %v", err)
	}

	return nil
}

func extractSound(filename string) (string, error) {
	data, err := soundFiles.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("error reading embedded sound file %s: %v", filename, err)
	}

	tempFile, err := os.CreateTemp("", "sound-*.mp3")
	if err != nil {
		return "", fmt.Errorf("error creating temp file: %v", err)
	}

	if _, err := tempFile.Write(data); err != nil {
		return "", fmt.Errorf("error writing to temp file: %v", err)
	}

	if err := tempFile.Close(); err != nil {
		return "", fmt.Errorf("error closing temp file: %v", err)
	}

	return tempFile.Name(), nil
}

func playMoveSound() {
	playSound(moveSound)
}

func playEndingSound() {
	playSound(endingSound)
}

func playSound(file string) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("afplay", file)
	case "linux":
		cmd = exec.Command("aplay", file)
	case "windows":
		cmd = exec.Command("powershell", "-c", "(New-Object Media.SoundPlayer '"+file+"').PlaySync()")
	default:
		fmt.Println("Unsupported operating system for audio playback")
		return
	}

	err := cmd.Start()
	if err != nil {
		fmt.Printf("Error playing sound: %v\n", err)
	}
}
