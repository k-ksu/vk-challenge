package techotestapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (c *TechTestAPIClient) UpdateContainersStatus() error {
	log.Println("Updating containers status...")

	conts, err := c.ContainerExplorer.ListAllContainersStatus()
	if err != nil {
		return fmt.Errorf("ContainerExplorer.ListAllContainersStatus: %w", err)
	}

	b, err := json.Marshal(conts)
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err)
	}

	_, err = http.Post(c.Addr+"/update_containers_status", "application/json", bytes.NewBuffer(b))
	if err != nil {
		return fmt.Errorf("http.Post: %w", err)
	}

	log.Println("Successfully updated containers status.")

	return nil
}
