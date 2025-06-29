package cluster

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

type ProgressLine struct {
	Status         string `json:"status"`
	ID             string `json:"id,omitempty"`
	Progress       string `json:"progress,omitempty"`
	ProgressDetail struct {
		Current int `json:"current"`
		Total   int `json:"total"`
	} `json:"progressDetail,omitempty"`
}

var progressMap = make(map[string]string)

func renderProgress() {
	// move cursor up and clear lines (ANSI)
	// fmt.Print("\033[H\033[2J") // optional: clear screen
	for id, line := range progressMap {
		fmt.Printf("%s: %s\n", id, line)
	}
}

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull latest Docker images",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		if err != nil {
			panic(err)
		}
		defer cli.Close()

		imageNames := []string{
			"ghcr.io/ente-io/server",
			"ghcr.io/ente-io/web",
			"alpine/socat",
			"postgres:15",
			"minio/minio",
		}

		for _, imageName := range imageNames {
			fmt.Printf("Pulling image: %s\n", imageName)

			out, err := cli.ImagePull(ctx, imageName, image.PullOptions{})
			if err != nil {
				fmt.Printf("Failed to pull image %s: %v\n", imageName, err)
				continue
			}
			defer out.Close()

			progressMap = make(map[string]string) // reset for each image

			scanner := bufio.NewScanner(out)
			for scanner.Scan() {
				var line ProgressLine
				if err := json.Unmarshal(scanner.Bytes(), &line); err != nil {
					continue
				}

				if line.ID != "" {
					text := line.Status
					if line.Progress != "" {
						text += " " + line.Progress
					}
					progressMap[line.ID] = text
				} else if strings.HasPrefix(line.Status, "Digest:") || strings.HasPrefix(line.Status, "Status:") {
					fmt.Println(line.Status)
				}

				renderProgress()
			}

			if err := scanner.Err(); err != nil {
				fmt.Printf("Error reading output for image %s: %v\n", imageName, err)
			}

			fmt.Println("Image pulled successfully")
		}
	},
}
