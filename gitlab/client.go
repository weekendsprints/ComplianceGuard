package gitlab

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// VerifyConnection makes an API call to GitLab's /application/settings endpoint
// and returns the settings as a JSON map
func VerifyConnection(gitlabURL, token string, verbose bool) (map[string]interface{}, error) {
	// Construct the API endpoint
	apiURL := gitlabURL + "/api/v4/application/settings"

	if verbose {
		fmt.Printf("  [VERBOSE] Calling GitLab API: %s\n", apiURL)
	}

	// Create HTTP request
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Add authorization header
	req.Header.Add("PRIVATE-TOKEN", token)

	// Make the API call
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to call GitLab API: %w", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("GitLab API returned status %d: %s", resp.StatusCode, string(body))
	}

	// Parse JSON response
	var settings map[string]interface{}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(body, &settings); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	return settings, nil
}
