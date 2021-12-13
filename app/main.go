package main

import (
	"GitDeployer/app/internal"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	fmt.Println("Flags: ")
	fmt.Println("--sm : spark migrate")
	fmt.Println("--dc: docker compose build && up")
	//git fetch --all
	//git reset --hard origin/master

	for {
		response := internal.GoBash("git", "fetch", "--all")
		for _, v := range strings.Split(response, "\n") {
			if strings.Contains(v, "origin/master") {
				fmt.Println("Got updates")

				response = internal.GoBash("git", "reset", "--hard", "origin/master")
				fmt.Println(response)
				fmt.Println("Pull procedure done.")

				if internal.ArrayContains(os.Args, "--spark") {
					response = internal.GoBash("php", "spark", "migrate")
					fmt.Println(response)
				}

				if internal.ArrayContains(os.Args, "--dc") {
					response = internal.GoBash("docker-compose", "build")
					fmt.Println(response)
					response = internal.GoBash("docker-compose", "up", "-d")
					fmt.Println(response)
				}

			}
		}

		time.Sleep(5 * time.Second)
	}
}
