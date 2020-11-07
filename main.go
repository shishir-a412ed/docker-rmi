package main

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/moby/moby/client"
	"github.com/pkg/term"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Fatalf("Error creating docker client: %v\n", err)
	}

	images, err := cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		log.Fatalf("Error in image list: %v\n", err)
	}

	for _, image := range images {
		for _, tag := range image.RepoTags {
			fmt.Printf("Delete: %s (y/n) ", tag)
			input, err := scanUserInput()
			if err != nil {
				log.Fatalln(err)
			}

			if input == "y" {
				fmt.Printf("\u2705")
			} else {
				fmt.Printf("\u274C")
			}
			fmt.Println()
		}
	}

}

// Scan user input (y/n) without newline '\n' character.
func scanUserInput() (string, error) {
	t, _ := term.Open("/dev/tty")
	term.RawMode(t)

	defer t.Close()
	defer t.Restore()

	b := make([]byte, 1)
	_, err := t.Read(b)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
