package signalfx

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

// This resource leverages common methods for read and delete from
// integration.go!

func integrationGCPResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the integration",
			},
			"enabled": &schema.Schema{
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Whether the integration is enabled or not",
			},
			"poll_rate": &schema.Schema{
				Type:         schema.TypeInt,
				Optional:     true,
				Description:  "GCP poll rate",
				ValidateFunc: validatePollRate,
			},
			"services": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "GCP enabled services",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"project_service_keys": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "GCP project service keys",
				Sensitive:   true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"project_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"project_key": {
							Type:      schema.TypeString,
							Required:  true,
							Sensitive: true,
						},
					},
				},
			},
			"synced": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Whether the resource in the provider and SignalFx are identical or not. Used internally for syncing.",
			},
			"last_updated": &schema.Schema{
				Type:        schema.TypeFloat,
				Computed:    true,
				Description: "Latest timestamp the resource was updated",
			},
		},

		Create: integrationGCPCreate,
		Read:   integrationRead,
		Update: integrationGCPUpdate,
		Delete: integrationDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func getGCPPayloadIntegration(d *schema.ResourceData) ([]byte, error) {
	payload := map[string]interface{}{
		"name":               d.Get("name").(string),
		"enabled":            d.Get("enabled").(bool),
		"type":               "GCP",
		"pollRate":           d.Get("poll_rate").(int),
		"services":           expandServices(d.Get("services").([]interface{})),
		"projectServiceKeys": expandProjectServiceKeys(d.Get("project_service_keys").([]interface{})),
	}
	return json.Marshal(payload)
}

func integrationGCPCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*signalfxConfig)
	payload, err := getGCPPayloadIntegration(d)
	if err != nil {
		return fmt.Errorf("Failed creating json payload: %s", err.Error())
	}
	url, err := buildURL(config.APIURL, INTEGRATION_API_PATH, map[string]string{})
	if err != nil {
		return fmt.Errorf("[DEBUG] SignalFx: Error constructing API URL: %s", err.Error())
	}

	return resourceCreate(url, config.AuthToken, payload, d)
}

func integrationGCPUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*signalfxConfig)
	payload, err := getGCPPayloadIntegration(d)
	if err != nil {
		return fmt.Errorf("Failed creating json payload: %s", err.Error())
	}
	path := fmt.Sprintf("%s/%s", INTEGRATION_API_PATH, d.Id())
	url, err := buildURL(config.APIURL, path, map[string]string{})
	if err != nil {
		return fmt.Errorf("[DEBUG] SignalFx: Error constructing API URL: %s", err.Error())
	}

	return resourceUpdate(url, config.AuthToken, payload, d)
}
