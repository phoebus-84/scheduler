package temporal

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const url = "https://api-conformance.ebsi.eu/trusted-issuers-registry/v5/issuers"

func FetchIssuersActivity(ctx context.Context) (FetchIssuersActivityResponse, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return FetchIssuersActivityResponse{}, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return FetchIssuersActivityResponse{}, fmt.Errorf("failed to perform HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return FetchIssuersActivityResponse{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return FetchIssuersActivityResponse{}, fmt.Errorf("failed to read response body: %w", err)
	}

	var root Root
	if err = json.Unmarshal(body, &root); err != nil {
		return FetchIssuersActivityResponse{}, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	hrefs := make([]string, 0, len(root.Items))
	for _, item := range root.Items {
		hrefs = append(hrefs, item.Href)
	}

	return FetchIssuersActivityResponse{Issuers: hrefs}, nil
}
