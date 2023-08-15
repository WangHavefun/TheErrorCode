package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func ExtractFirstFrameAsImage(inputBuffer *bytes.Buffer) (*bytes.Buffer, error) {
	outputBuffer := &bytes.Buffer{}

	tmpFile, err := ioutil.TempFile("", "input.mp4")
	if err != nil {
		return nil, fmt.Errorf("error creating temp file: %v", err)
	}
	defer func() {
		tmpFile.Close()
		os.Remove(tmpFile.Name()) // 删除临时文件
	}()

	// 将输入字节流写入临时文件
	_, err = tmpFile.Write(inputBuffer.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error writing to temp file: %v", err)
	}

	cmd := exec.Command("ffmpeg", "-i", tmpFile.Name(), "-vframes", "1", "-f", "image2", "pipe:1")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = outputBuffer

	err = cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("error running FFmpeg: %v\n%s", err, stderr.String())
	}

	return outputBuffer, nil
}
