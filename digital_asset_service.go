package sec

import (
	"context"
	"encoding/json"
	"fmt"
)

const digitalAssetBaseURL = "https://api.sec.or.th/DigitalAsset"

// DigitalAssetIntermediary represents a digital asset intermediary profile.
type DigitalAssetIntermediary struct {
	UniqueID      string `json:"unique_id"`
	NameTH        string `json:"name_th"`
	NameEN        string `json:"name_en"`
	LicCode       string `json:"lic_code"`
	LicActionCode string `json:"lic_action_code"`
	LicEffDate    string `json:"lic_efft_date"`
	LicActDate    string `json:"lic_act_date"`
	LicExpDate    string `json:"lic_exp_date"`
}

// DigitalAssetIntermediaryRequest represents the request body for searching digital asset intermediaries.
type DigitalAssetIntermediaryRequest struct {
	IntermediaryName string `json:"IntermediaryName"`
}

// GetDigitalAssetIntermediaries returns digital asset intermediary profiles.
// Pass an empty IntermediaryName to list all intermediaries.
func (c *Client) GetDigitalAssetIntermediaries(ctx context.Context, req DigitalAssetIntermediaryRequest) ([]DigitalAssetIntermediary, error) {
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	url := digitalAssetBaseURL + "/profile/intermediary"
	data, err := c.postAbsolute(ctx, url, payload)
	if err != nil {
		return nil, fmt.Errorf("get digital asset intermediaries: %w", err)
	}

	var items []DigitalAssetIntermediary
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("unmarshal digital asset intermediaries: %w", err)
	}
	return items, nil
}
