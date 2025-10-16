package androidnuclei

import (
	"bufio"
	"os"
	"strings"
	"sync"

	"github.com/projectdiscovery/nuclei/v2/pkg/runner"
)

// RunNucleiSingle scans a single target with a template
func RunNucleiSingle(target, template string) (string, error) {
	options := runner.Options{
		Targets:   []string{target},
		Templates: []string{template},
	}

	r, err := runner.New(&options)
	if err != nil {
		return "", err
	}
	defer r.Close()

	if err := r.Run(); err != nil {
		return "", err
	}
	return "Scan completed for " + target, nil
}

// RunNucleiFile runs nuclei on all targets from a file (parallel)
func RunNucleiFile(targetsFile, template string, maxConcurrency int) (string, error) {
	targets, err := readLines(targetsFile)
	if err != nil {
		return "", err
	}

	sem := make(chan struct{}, maxConcurrency)
	var wg sync.WaitGroup

	for _, t := range targets {
		t = strings.TrimSpace(t)
		if t == "" {
			continue
		}
		wg.Add(1)
		go func(target string) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			options := runner.Options{
				Targets:   []string{target},
				Templates: []string{template},
			}

			r, err := runner.New(&options)
			if err == nil {
				defer r.Close()
				_ = r.Run()
			}
		}(t)
	}
	wg.Wait()

	return "File scan completed", nil
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
