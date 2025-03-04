package main

import (
	"log"

	"fmt"

	gomaasclient "github.com/canonical/gomaasclient/client"
)

const (
	MAAS_URL      = "..."    // "http://10.10.0.22:5240/MAAS/"
	MAAS_API_KEY  = "..."    //"xxxxxxxxxxxxxxxxx:yyyyyyyyyyyyyyyyyy:zzzzzzzzzzzzzzzzzz"
	MAAS_API_VERSION = "2.0"
)

func main() {
	client, err := gomaasclient.GetClient(
		MAAS_URL,
		MAAS_API_KEY,
		MAAS_API_VERSION,
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
