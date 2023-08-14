package utils

import (
	"bytes"
	"fmt"
	"os/exec"
)

func ExtractFirstFrameAsImage(inputBuffer *bytes.Buffer) (*bytes.Buffer, error) {
	outputBuffer := &bytes.Buffer{}

	cmd := exec.Command("ffmpeg", "-i", "pipe:0", "-vframes", "1", "-f", "image2", "pipe:1")
	cmd.Stdin = inputBuffer
	cmd.Stdout = outputBuffer
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("error running FFmpeg: %v\n%s", err, stderr.String())
	}

	return outputBuffer, nil
}
