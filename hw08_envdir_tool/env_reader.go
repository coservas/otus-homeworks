package main

import (
	"errors"
	"io/ioutil"
	"strings"
)

type Environment map[string]EnvValue

type EnvValue struct {
	Value      string
	NeedRemove bool
}

func ReadDir(dir string) (Environment, error) {
	if len(dir) == 0 {
		return nil, errors.New("directory path is empty")
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	envs := make(Environment)

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if file.Size() == 0 {
			envs[file.Name()] = EnvValue{NeedRemove: true}
			continue
		}

		data, err := ioutil.ReadFile(dir + "/" + file.Name())
		if err != nil {
			return nil, err
		}

		strData := strings.Builder{}
		for _, datum := range data {
			if datum == '\n' {
				break
			}

			if datum == 0x00 {
				datum = '\n'
			}

			strData.WriteByte(datum)
		}

		value := strings.TrimRight(strData.String(), " \t\n")

		envs[file.Name()] = EnvValue{Value: value}
	}

	return envs, nil
}
