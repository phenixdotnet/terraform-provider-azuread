package azuread

import (
	"log"
	"os"
	"strings"
)

// This file contains feature flags for functionality which will prove more challenging to implement en-mass
var requireResourcesToBeImported = strings.EqualFold(os.Getenv("ARM_PROVIDER_STRICT"), "true")

/*UserFeatures represents the features a user can configure for use in this provider */
type UserFeatures struct {
	MSGraphAPI MSGraphAPIFeatures
}

/*MSGraphAPIFeatures represents the options for MS Graph */
type MSGraphAPIFeatures struct {
	IsEnabledForExistingResources bool
	UseBetaEndPoint               bool
}

func expandFeatures(input []interface{}) UserFeatures {
	// these are the defaults if omitted from the config
	features := UserFeatures{
		MSGraphAPI: MSGraphAPIFeatures{
			IsEnabledForExistingResources: false,
			UseBetaEndPoint:               false,
		},
	}

	log.Printf("[DEBUG] features length: %s", len(input))
	if len(input) == 0 || input[0] == nil {
		return features
	}

	val := input[0].(map[string]interface{})

	if raw, ok := val["msgraphapi"]; ok {
		items := raw.([]interface{})
		if len(items) > 0 {
			msGraphAPIRaw := items[0].(map[string]interface{})
			if v, ok := msGraphAPIRaw["is_enabled_for_existing_resources"]; ok {
				log.Printf("[DEBUG] setting IsEnabledForExistingResources to %s", v.(bool))
				features.MSGraphAPI.IsEnabledForExistingResources = v.(bool)
			}
			if v, ok := msGraphAPIRaw["use_beta_endpoint"]; ok {
				log.Printf("[DEBUG] setting UseBetaEndPoint to %s", v.(bool))
				features.MSGraphAPI.UseBetaEndPoint = v.(bool)
			}
		}
	}

	return features
}
