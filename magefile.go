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

type BUILD mg.Namespace

func Clean() {
	os.RemoveAll("build")
}

func Build() {
	b := BUILD{}
	mg.Deps(b.GIO, b.Fyne)
}

func (b BUILD) GIO() error {
	sources := []string{
		"./gio/counter",
		"./gio/temperature-converter",
		"./gio/timer",
	}

	goexe := mg.GoCmd()
	for _, source := range sources {
		if err := b.build(goexe, source); err != nil {
			return err
		}
	}
	return nil
}

func (b BUILD) Fyne() error {
	sources := []string{
		"./fyne/counter",
		"./fyne/temperature-converter",
		"./fyne/timer",
	}

	goexe := mg.GoCmd()
	for _, source := range sources {
		if err := b.build(goexe, source); err != nil {
			return err
		}
	}
	return nil
}

func (BUILD) build(goexe, source string) error {
	dir := filepath.Join("build", filepath.Dir(source))
	file := filepath.Join("build", source)
	if modified, err := target.Dir(file, source); err != nil {
		return fmt.Errorf("check target modified error: %w", err)
	} else if !modified {
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
