package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"testing"
)

var sites = []string{
	"instagram",
	"facebook",
	"google",
}

func TempFile() *os.File {
	tmp, err := os.CreateTemp("", "hosts_test")
	if err != nil {
		log.Fatalf("Error creating a temporary file: %v \n", err)
	}
	return tmp
}

func FileScanner(path string) (*bufio.Scanner, *os.File) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	fs := bufio.NewScanner(f)
	fs.Split(bufio.ScanLines)
	return fs, f
}

func TestAdd(t *testing.T) {
	tmp := TempFile()
	defer os.Remove(tmp.Name())

	for _, s := range sites {
		h := &HostsFile{
			path:     tmp.Name(),
			startPos: 0,
			endPos:   0,
			content:  []string{},
			site:     "",
		}
		add(h, []string{"add", s})
		fs, f := FileScanner(h.path)
		defer f.Close()
		found := false
		for fs.Scan() {
			if strings.Contains(fs.Text(), s) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Did not find %s after running add() operation in hosts file.", h.site)
		}
	}
}

func TestRemove(t *testing.T) {
	tmp := TempFile()
	defer os.Remove(tmp.Name())

	for _, s := range sites {
		h := &HostsFile{
			path:     tmp.Name(),
			startPos: 0,
			endPos:   0,
			content:  []string{},
			site:     "",
		}
		add(h, []string{"add", s})
		remove(h, []string{"remove", s})

		fs, f := FileScanner(h.path)
		defer f.Close()
		found := false
		for fs.Scan() {
			if strings.Contains(fs.Text(), s) {
				found = true
				break
			}
		}
		if found {
			t.Errorf("Found %s after running remove() operation in hosts file.", h.site)
		}
		f.Close()
	}
}

func TestDisable(t *testing.T) {
	tmp := TempFile()
	defer os.Remove(tmp.Name())

	for _, s := range sites {
		h := &HostsFile{
			path:     tmp.Name(),
			startPos: 0,
			endPos:   0,
			content:  []string{},
			site:     "",
		}
		add(h, []string{"add", s})
		disable(h, []string{"disable", s})
		fs, f := FileScanner(h.path)
		defer f.Close()
		disabled := false
		for fs.Scan() {
			if strings.Contains(fs.Text(), s) {
				trimmed := strings.TrimSpace(s)
				if trimmed[0] == '#' {
				}
				disabled = true
				break
			}
		}
		if !disabled {
			t.Errorf("Did not disable properly %s after running disable() operation in hosts file.", h.site)
		}
		f.Close()
	}
}

func TestEnable(t *testing.T) {
	tmp := TempFile()
	defer os.Remove(tmp.Name())

	for _, s := range sites {
		h := &HostsFile{
			path:     tmp.Name(),
			startPos: 0,
			endPos:   0,
			content:  []string{},
			site:     "",
		}
		add(h, []string{"add", s})
		disable(h, []string{"disable", s})
		enable(h, []string{"enable", s})
		fs, f := FileScanner(h.path)
		defer f.Close()
		enabled := false
		for fs.Scan() {
			if strings.Contains(fs.Text(), s) {
				trimmed := strings.TrimSpace(s)
				if !(trimmed[0] == '#') {
				}
				enabled = true
				break
			}
		}
		if !enabled {
			t.Errorf("Did not enable properly %s after running disable() operation in hosts file.", h.site)
		}
		f.Close()
	}
}
