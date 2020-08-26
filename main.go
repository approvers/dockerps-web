package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
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
	router.Use(static.ServeRoot("/", "./node_modules/@approvers/dockerps-web-frontend/public"))
	router.GET("/api.json", func(ctx *gin.Context) {
		containers, err := docker.ContainerList(context.Background(), types.ContainerListOptions{All: true})

		if err != nil {
			ctx.Status(500)
		} else {
			ctx.JSON(200, formatContainers(containers))
		}
	})

	fmt.Println("Start Serving...")
	err = router.Run()
	if err != nil {
		fmt.Println("Failed at starting Web server")
		fmt.Println(err)
		return
	}
}

func formatContainers(containers []types.Container) (result []map[string]string) {
	for _, container := range containers {
		local := map[string]string{
			"image":        container.Image,
			"id":           container.ID[0:12],
			"command":      container.Command,
			"created":      formatUnixDate(container.Created),
			"status":       container.Status,
			"ports":        formatPortArray(container.Ports),
			"name":         formatStringArray(container.Names),
			"status_title": strings.Split(container.Status, " ")[0],
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

func formatStringArray(array []string) string {
	return strings.Join(array, ", ")
}
