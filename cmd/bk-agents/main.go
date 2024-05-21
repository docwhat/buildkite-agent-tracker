package main

import (
	"fmt"
	"log"
	"os"

	"github.com/buildkite/go-buildkite/v3/buildkite"
)

func main() {
	token := os.Getenv("BUILDKITE_API_TOKEN")
	org := os.Getenv("BUILDKITE_ORG")

	config, err := buildkite.NewTokenConfig(token, false)

	if err != nil {
		log.Fatalf("client config failed: %s", err)
	}

	client := buildkite.NewClient(config.Client())

	page := 1
	per_page := 100

	for more_agents := true; more_agents; {
		opts := &buildkite.AgentListOptions{
			ListOptions: buildkite.ListOptions{
				Page:    page,
				PerPage: per_page,
			},
		}

		agents, resp, err := client.Agents.List(org, opts)

		if err != nil {
			log.Fatalf("client agents failed: %s", err)
		}

		for _, agent := range agents {
			if agent.Job == nil {
				fmt.Printf("free %s\n", *agent.Name)
			} else {
				fmt.Printf("busy %s\n", *agent.Name)
			}
		}

		if resp.NextPage == 0 {
			more_agents = false
		} else {
			page = resp.NextPage
		}
	}
}

// EOF
