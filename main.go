package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	docker, err := client.NewEnvClient()
	if err != nil {
		fmt.Println("Failed at connecting to docker")
		fmt.Println(err)
		return
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.LoadHTMLFiles("./template.html")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "template.html", gin.H{
			"containers": getContainers(docker),
		})
	})

	fmt.Println("Start Serving...")
	err = router.Run()
	if err != nil {
		fmt.Println("Failed at starting Web server")
		fmt.Println(err)
		return
	}
}

func getContainers(docker *client.Client) (result []map[string]string) {
	result = make([]map[string]string, 0)

	containers, err := docker.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		fmt.Println("Failed at getting container list")
		fmt.Println(err)
		return
	}

	for _, container := range containers {
		local := map[string]string{
			"Image":   container.Image,
			"ID":      container.ID[0:12],
			"Command": container.Command,
			"Created": formatUnixDate(container.Created),
			"Status":  container.Status,
			"Ports":   formatPortArray(container.Ports),
			"Names":   formatStringArray(container.Names),
		}

		result = append(result, local)
	}
	return
}

func formatUnixDate(unix int64) string {
	durationFromNow := time.Now().Sub(time.Unix(unix, 0)).Truncate(time.Second)
	return fmt.Sprintf("%s ago", durationFromNow.String())
}

func formatPortArray(array []types.Port) (result string) {
	for _, v := range array {
		result += fmt.Sprintf("%s:%d->%d/%s, ", v.IP, v.PrivatePort, v.PublicPort, v.Type)
	}
	return
}

func formatStringArray(array []string) (result string) {
	for _, v := range array {
		result += fmt.Sprintf("%s, ", v)
	}
	return
}
