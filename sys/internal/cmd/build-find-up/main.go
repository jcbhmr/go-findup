package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

const Version = "8.0.0"

func License(version string) (license []byte, err error) {
	resp, err := http.Get("https://unpkg.com/find-up@" + Version + "/LICENSE")
	if err != nil {
		return nil, err
	}
	defer func() {
		err = errors.Join(err, resp.Body.Close())
	}()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status %q: %s %d", resp.Request.URL, resp.Status, resp.StatusCode)
	}
	license, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return license, nil
}

func main() {
	f, err := os.Create("./find-up.js")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = errors.Join(err, f.Close())
	}()

	license, err := License(Version)
	if err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Fprintf(f, "/*!\n%s\n*/\n", bytes.Trim(license, "\r\n"))
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("deno", "bundle", "npm:find-up@"+Version, "--minify")
	js, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.Write(js)
	if err != nil {
		log.Fatal(err)
	}
}
