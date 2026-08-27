package dotenv

import (
	"bufio"
	"fmt"
	"log"
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
	defer func() {
		if err := source.Close(); err != nil {
			log.Printf("Load source.Close() err: %v", err)
		}
	}()

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
		err = os.Setenv(k, v)
		if err != nil {
			return err
		}
	}

	return nil
}
