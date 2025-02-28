package main

import (
	"log"
	"os"

	"fmt"

	gomaasclient "github.com/canonical/gomaasclient/client"
)

func main() {
	client, err := gomaasclient.GetClient(
		os.Getenv("MAAS_URL"),
		os.Getenv("MAAS_API_KEY"),
		os.Getenv("MAAS_API_VERSION"),
	)
	if err != nil {
		log.Fatalf("Error getting client: %v", err)
	}
	// Get all machines
	machines, err := client.Machines.Get(nil)
	if err != nil {
		log.Fatalf("Error getting machines: %v", err)
	}

	// Print machine details
	fmt.Println("MAAS Machines:")
	fmt.Println("-------------")
	for _, machine := range machines {
		fmt.Printf("System ID: %s, Hostname: %s\n",
			machine.SystemID, machine.Hostname)
	}
}
