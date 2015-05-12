package main

import (
	"os/exec"
)

func validateLcovEnv() error {

	_, err := exec.Command("lcov", "--version").Output()
	if err != nil {
		return err
	}

	_, err = exec.Command("gcov", "--version").Output()
	if err != nil {
		return err
	}

	return nil
}

func Lcov(dir string, info string) error {

	validateLcovEnv()
	_, err := exec.Command("lcov", "--capture", "--directory", dir, "--output-file", info).Output()
	return err
}
