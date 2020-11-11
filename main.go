package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/eiannone/keyboard"
	"github.com/moby/moby/client"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Fatalf("Error creating docker client: %v\n", err)
	}

	noDanglingfilters := filters.NewArgs()
	noDanglingfilters.Add("dangling", "false")

	images, err := cli.ImageList(ctx, types.ImageListOptions{Filters: noDanglingfilters})
	if err != nil {
		log.Fatalf("Error in image list: %v\n", err)
	}

	if len(images) == 0 {
		fmt.Println("No available docker images to remove. Exiting...")
		os.Exit(0)
	}

	if err := keyboard.Open(); err != nil {
		log.Fatalf("Error in opening terminal in raw mode: %v\n", err)
	}

	defer closeKeyboard()

	var wg sync.WaitGroup
	for _, image := range images {
		for _, tag := range image.RepoTags {
			fmt.Printf("Delete: %s (y/n) ", tag)
			input, key, err := keyboard.GetKey()
			if err != nil {
				log.Fatalf("Error in scanning user input: %v\n", err)
			}

			if key == keyboard.KeyCtrlC {
				closeKeyboard()
				os.Exit(0)
			}

			if input == 'y' || input == 'Y' {
				fmt.Printf("\u2705")
				wg.Add(1)
				go removeDockerImage(cli, ctx, tag, &wg)
			} else {
				fmt.Printf("\u274C")
			}
			fmt.Println()
		}
	}

	// Check and remove dangling images.
	danglingFilters := filters.NewArgs()
	danglingFilters.Add("dangling", "true")

	images, err = cli.ImageList(ctx, types.ImageListOptions{Filters: danglingFilters})
	if err != nil {
		log.Fatalf("Error in listing dangling images: %v\n", err)
	}

	if len(images) > 0 {
		removeDanglingImages(images)
	}

	wg.Wait()
}

func removeDanglingImages(images []types.ImageSummary) {
	fmt.Print("Delete: dangling <none>:<none> images (y/n) ")
	input, key, err := keyboard.GetKey()
	if err != nil {
		log.Fatalf("Error in scanning user input: %v\n", err)
	}

	if key == keyboard.KeyCtrlC {
		closeKeyboard()
		os.Exit(0)
	}
	if input == 'y' || input == 'Y' {
		fmt.Printf("\u2705")
		for _, image := range images {
			cmd := exec.Command("docker", "rmi", "--force", fmt.Sprintf("%s", image.ID))
			if err := cmd.Run(); err != nil {
				closeKeyboard()
				log.Fatalln(err)
			}
		}
	} else {
		fmt.Printf("\u274C")
	}
	fmt.Println()
}

func removeDockerImage(cli *client.Client, ctx context.Context, tag string, wg *sync.WaitGroup) {
	// docker client SDK only has an {ImageRemove} method, and no method for untagging images.
	// We don't want to forcefully remove an image, if it has multiple tags.
	// Rather we would just like to untag the tag. If it's the only tag left, we will remove the image.
	// exec docker CLI due to lack of support from docker client SDK.
	cmd := exec.Command("docker", "rmi", "--force", fmt.Sprintf("%s", tag))
	if err := cmd.Run(); err != nil {
		closeKeyboard()
		log.Fatalln(err)
	}
	wg.Done()
}

func closeKeyboard() {
	if err := keyboard.Close(); err != nil {
		log.Fatalf("Error in closing terminal in raw mode: %v\n", err)
	}
}
