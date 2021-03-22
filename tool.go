package main

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

func RunPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		logrus.Error(err)
	}
	return dir + "/"
}
