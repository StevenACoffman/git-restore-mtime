package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	gitLs, e := popen("git", "ls-files")
	check(e)
	files := map[string]bool{}
	for gitLs.Scan() {
		files[gitLs.Text()] = true
	}
	gitLog, e := popen(
		"git", "log", "-m",
		"--name-only", "--relative", "--pretty=format:%ct", ".",
	)
	check(e)
	for len(files) > 0 {
		gitLog.Scan()
		sec, e := strconv.ParseInt(gitLog.Text(), 10, 64)
		check(e)
		unix := time.Unix(sec, 0)
		for gitLog.Scan() {
			name := gitLog.Text()
			if name == "" {
				break
			}
			if ! files[name] {
				continue
			}
			os.Chtimes(name, unix, unix)
			delete(files, name)
		}
	}
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func popen(name string, arg ...string) (*bufio.Scanner, error) {
	cmd := exec.Command(name, arg...)
	pipe, e := cmd.StdoutPipe()
	if e != nil {
		return nil, e
	}
	return bufio.NewScanner(pipe), cmd.Start()
}

