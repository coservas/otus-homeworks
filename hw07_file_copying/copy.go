package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	srcFile, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	fStat, err := srcFile.Stat()
	if err != nil {
		return err
	}

	if !fStat.Mode().IsRegular() {
		return ErrUnsupportedFile
	}

	if offset > fStat.Size() {
		return ErrOffsetExceedsFileSize
	}

	if _, err = srcFile.Seek(offset, io.SeekStart); err != nil {
		return err
	}

	if limit == 0 {
		limit = fStat.Size()
	}

	dstFile, _ := os.Create(toPath)
	defer dstFile.Close()

	bar := pb.Full.Start64(limit)

	if _, err := io.CopyN(io.Writer(dstFile), bar.NewProxyReader(io.Reader(srcFile)), limit); err != nil {
		return err
	}

	bar.Finish()

	return nil
}
