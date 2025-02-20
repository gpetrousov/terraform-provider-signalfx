package signalfx

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	sfx "github.com/signalfx/signalfx-go"
)

const newHeatmapChartConfig = `
resource "signalfx_heatmap_chart" "mychartHX" {
  name = "Fart Heatmap"
  description = "Farts"
	program_text = "data('cpu.total.idle').publish(label='CPU Idle')"

	disable_sampling = true
	hide_timestamp = true
	sort_by = "-foo"
	group_by = ["a", "b"]

	color_range {
		min_value = 1
		max_value = 100
		color = "magenta"
	}
}
`

const updatedHeatmapChartConfig = `
resource "signalfx_heatmap_chart" "mychartHX" {
  name = "Fart Heatmap NEW"
  description = "Farts NEW"
	program_text = "data('cpu.total.idle').publish(label='CPU Idle')"

	disable_sampling = true
	hide_timestamp = true
	sort_by = "-foo"
	group_by = ["a", "b"]

	color_range {
		min_value = 1
		max_value = 100
		color = "magenta"
	}
}
`

func TestAccCreateUpdateHeatmapChart(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccHeatmapChartDestroy,
		Steps: []resource.TestStep{
			// Create It
			{
				Config: newHeatmapChartConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHeatmapChartResourceExists,
					resource.TestCheckResourceAttr("signalfx_heatmap_chart.mychartHX", "name", "Fart Heatmap"),
					resource.TestCheckResourceAttr("signalfx_heatmap_chart.mychartHX", "description", "Farts"),
					resource.TestCheckResourceAttr("signalfx_heatmap_chart.mychartHX", "program_text", "data('cpu.total.idle').publish(label='CPU Idle')"),
					resource.TestCheckResourceAttr("signalfx_heatmap_chart.mychartHX", "disable_sampling", "true"),
					resource.TestCheckResourceAttr("signalfx_heatmap_chart.mychartHX", "hide_timestamp", "true"),
					resource.TestCheckResourceAttr("signalfx_heatmap_chart.mychartHX", "sort_by", "-foo"),
					resource.TestCheckResourceAttr("signalfx_heatmap_chart.mychartHX", "color_range.#", "1"),
					resource.TestCheckResourceAttr("signalfx_heatmap_chart.mychartHX", "color_range.452638366.color", "magenta"),
					resource.TestCheckResourceAttr("signalfx_heatmap_chart.mychartHX", "color_range.452638366.max_value", "100"),
					resource.TestCheckResourceAttr("signalfx_heatmap_chart.mychartHX", "color_range.452638366.min_value", "1"),
					resource.TestCheckResourceAttr("signalfx_heatmap_chart.mychartHX", "group_by.#", "2"),
					resource.TestCheckResourceAttr("signalfx_heatmap_chart.mychartHX", "group_by.0", "a"),
					resource.TestCheckResourceAttr("signalfx_heatmap_chart.mychartHX", "group_by.1", "b"),
				),
			},
			{
				ResourceName:      "signalfx_heatmap_chart.mychartHX",
				ImportState:       true,
				ImportStateIdFunc: testAccStateIdFunc("signalfx_heatmap_chart.mychartHX"),
				ImportStateVerify: true,
			},
			// Update Everything
			{
				Config: updatedHeatmapChartConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHeatmapChartResourceExists,
					resource.TestCheckResourceAttr("signalfx_heatmap_chart.mychartHX", "name", "Fart Heatmap NEW"),
					resource.TestCheckResourceAttr("signalfx_heatmap_chart.mychartHX", "description", "Farts NEW"),
				),
			},
		},
	})
}

func testAccCheckHeatmapChartResourceExists(s *terraform.State) error {
	client, _ := sfx.NewClient(os.Getenv("SFX_AUTH_TOKEN"))

	for _, rs := range s.RootModule().Resources {
		switch rs.Type {
		case "signalfx_heatmap_chart":
			chart, err := client.GetChart(rs.Primary.ID)
			if chart.Id != rs.Primary.ID || err != nil {
				return fmt.Errorf("Error finding chart %s: %s", rs.Primary.ID, err)
			}
		default:
			return fmt.Errorf("Unexpected resource of type: %s", rs.Type)
		}
	}
	return nil
}

func testAccHeatmapChartDestroy(s *terraform.State) error {
	client, _ := sfx.NewClient(os.Getenv("SFX_AUTH_TOKEN"))
	for _, rs := range s.RootModule().Resources {
		switch rs.Type {
		case "signalfx_heatmap_chart":
			chart, _ := client.GetChart(rs.Primary.ID)
			if chart != nil {
				return fmt.Errorf("Found deleted chart %s", rs.Primary.ID)
			}
		default:
			return fmt.Errorf("Unexpected resource of type: %s", rs.Type)
		}
	}

	return nil
}

func TestValidateHeatmapChartColors(t *testing.T) {
	_, err := validateHeatmapChartColor("blue", "color")
	assert.Equal(t, 0, len(err))
}

func TestValidateHeatmapChartColorsFail(t *testing.T) {
	_, err := validateHeatmapChartColor("whatever", "color")
	assert.Equal(t, 1, len(err))
}
