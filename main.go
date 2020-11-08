package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"

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

	if len(images) == 0 {
		fmt.Println("No available docker images to remove. Exiting...")
		os.Exit(0)
	}

	var wg sync.WaitGroup
	for _, image := range images {
		for _, tag := range image.RepoTags {
			fmt.Printf("Delete: %s (y/n) ", tag)
			input, err := scanUserInput()
			if err != nil {
				log.Fatalf("Error in scanning user input: %v\n", err)
			}

			input = strings.ToLower(input)
			if input == "y" {
				fmt.Printf("\u2705")
				wg.Add(1)
				go removeDockerImage(cli, ctx, tag, &wg)
			} else {
				fmt.Printf("\u274C")
			}
			fmt.Println()
		}
	}
	wg.Wait()
}

func removeDockerImage(cli *client.Client, ctx context.Context, tag string, wg *sync.WaitGroup) {
	// docker client SDK only has an {ImageRemove} method, and no method for untagging images.
	// We don't want to forcefully remove an image, if it has multiple tags.
	// Rather we would just like to untag the tag. If it's the only tag left, we will remove the image.
	// exec docker CLI due to lack of support from docker client SDK.
	cmd := exec.Command("docker", "rmi", "--force", fmt.Sprintf("%s", tag))
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
	wg.Done()
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
