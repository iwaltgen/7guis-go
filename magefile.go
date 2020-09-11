// +build mage

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/magefile/mage/target"
)

func Clean() {
	os.RemoveAll("build")
}

func Build() {
	mg.Deps(BuildGio, BuildFyne, BuildLorca)
}

func BuildGio() error {
	sources := []string{
		"./gio/counter",
		"./gio/temperature-converter",
	}

	goexe := mg.GoCmd()
	for _, source := range sources {
		if err := build(goexe, source); err != nil {
			return err
		}
	}
	return nil
}

func BuildFyne() error {
	sources := []string{
		"./fyne/counter",
		"./fyne/temperature-converter",
	}

	goexe := mg.GoCmd()
	for _, source := range sources {
		if err := build(goexe, source); err != nil {
			return err
		}
	}
	return nil
}

func BuildLorca() error {
	sources := []string{
		"./lorca/counter",
	}

	goexe := mg.GoCmd()
	for _, source := range sources {
		if err := build(goexe, source); err != nil {
			return err
		}
	}
	return nil
}

func build(goexe, source string) error {
	dir := filepath.Join("build", filepath.Dir(source))
	file := filepath.Join("build", source)
	if midified, err := target.Dir(file, source); err != nil {
		return fmt.Errorf("check target modified error: %w", err)
	} else if !midified {
		return nil
	}

	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("mkdir error: %w", err)
	}

	if err := sh.RunV(goexe, "generate", source); err != nil {
		return err
	}

	ldflags := "-s -w"
	if runtime.GOOS == "windows" {
		ldflags = ldflags + " -H windowsgui"
	}
	return sh.RunV(goexe, "build", "-trimpath", "-ldflags", ldflags, "-o", file, source)
}

func UPX() {
	_ = sh.RunV("sh", "-c", "upx build/**/*")
}
