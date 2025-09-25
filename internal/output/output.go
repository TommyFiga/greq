package output

import (
	"fmt"
	"os"
)

func WriteResponseContentToFile(respCont string, filename string) error {
	if filename == "" {
		fmt.Println(respCont)
		return nil
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create output file %s: %w", filename, err)
	}

	defer file.Close()

	_, err = file.WriteString(respCont)
	if err != nil {
		return fmt.Errorf("failed to write to file %s: %w", filename, err)
	}

	fmt.Printf("Output written to %s\n", filename)
	
	return nil
}