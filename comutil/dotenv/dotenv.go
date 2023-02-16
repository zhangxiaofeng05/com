package dotenv

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Load the specified file and set the environment variable
//
// tip: override global variable
func Load(dotenvPath string) error {
	source, err := os.Open(dotenvPath)
	if err != nil {
		return err
	}
	defer source.Close()

	mp := make(map[string]string)
	scanner := bufio.NewScanner(source)
	for scanner.Scan() {
		line := scanner.Text()
		list := strings.Split(line, "=")
		if len(list) != 2 {
			return fmt.Errorf("line: %v invalid", line)
		}
		k, v := list[0], list[1]
		if _, ok := mp[k]; ok {
			return fmt.Errorf("key: %v repeat", k)
		}
		mp[k] = v
	}

	for k, v := range mp {
		os.Setenv(k, v)
	}

	return nil
}
