package util

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

type Manifest struct {
	Entry         string
	Height, Width int
}

type ExtPackage struct {
	EntryHtml string
	Manifest  *Manifest
}

func ReadFile(f *zip.File) ([]byte, error) {
	rc, err := f.Open()
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	content, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func ReadManifest(m []byte) (*Manifest, error) {
	var manifest Manifest
	dec := json.NewDecoder(strings.NewReader(string(m)))
	if err := dec.Decode(&manifest); err == io.EOF {
	} else if err != nil {
		return nil, err
	}
	return &manifest, nil
}

func FindFileInArchive(fn string, files []*zip.File) (index int) {
	index = -1
	for i, f := range files {
		if f.Name == fn {
			index = i
		}
	}
	return
}

func OpenPackage(fn string) (*ExtPackage, error) {
	r, err := zip.OpenReader(fn)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	manifestIndex := FindFileInArchive("manifest.json", r.File)
	if manifestIndex < 0 {
		return nil, fmt.Errorf("no manifest file found in archive")
	}
	f, err := ReadFile(r.File[manifestIndex])
	if err != nil {
		return nil, err
	}
	manifest, err := ReadManifest(f)
	if err != nil {
		return nil, err
	}
	if manifest.Entry == "" {
		return nil, fmt.Errorf("entry file specified but appears to be empty: %s", manifest.Entry)
	}
	entryIndex := FindFileInArchive(manifest.Entry, r.File)
	if entryIndex < 0 {
		return nil, fmt.Errorf("entry file not found: '%s'", manifest.Entry)
	}
	f, err = ReadFile(r.File[entryIndex])
	if err != nil {
		return nil, err
	}
	extPackage := &ExtPackage{string(f), manifest}
	return extPackage, nil
}
