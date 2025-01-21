package spaces_test

import (
	"fmt"
	"testing"

	"github.com/digitalocean/terraform-provider-digitalocean/digitalocean/acceptance"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDigitalOceanSpacesKey_basic(t *testing.T) {
	expectedName := acceptance.RandomTestName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.TestAccPreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckDigitalOceanSpacesKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDigitalOceanSpacesKeyConfig(expectedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"digitalocean_spaces_key.key", "name", expectedName),
					resource.TestCheckResourceAttr(
						"digitalocean_spaces_key.key", "grant.0.bucket", "my-bucket"),
					resource.TestCheckResourceAttr(
						"digitalocean_spaces_key.key", "grant.0.permission", "read"),
				),
			},
		},
	})
}

func TestAccDigitalOceanSpacesKey_updateGrant(t *testing.T) {
	expectedName := acceptance.RandomTestName()
	expectedNewName := acceptance.RandomTestName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.TestAccPreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckDigitalOceanSpacesKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDigitalOceanSpacesKeyConfig(expectedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"digitalocean_spaces_key.key", "name", expectedName),
					resource.TestCheckResourceAttr(
						"digitalocean_spaces_key.key", "grant.0.bucket", "my-bucket"),
					resource.TestCheckResourceAttr(
						"digitalocean_spaces_key.key", "grant.0.permission", "read"),
				),
			},
			{
				Config: testAccDigitalOceanSpacesKeyConfigWithGrantUpdate(expectedNewName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"digitalocean_spaces_key.key", "name", expectedNewName),
					resource.TestCheckResourceAttr(
						"digitalocean_spaces_key.key", "grant.0.bucket", "my-bucket2"),
					resource.TestCheckResourceAttr(
						"digitalocean_spaces_key.key", "grant.0.permission", "read"),
				),
			},
		},
	})
}

func TestAccDigitalOceanSpacesKey_multipleGrants(t *testing.T) {
	expectedName := acceptance.RandomTestName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.TestAccPreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckDigitalOceanSpacesKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDigitalOceanSpacesKeyConfigMultipleGrants(expectedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"digitalocean_spaces_key.key", "name", expectedName),
					resource.TestCheckResourceAttr(
						"digitalocean_spaces_key.key", "grant.0.bucket", "my-bucket"),
					resource.TestCheckResourceAttr(
						"digitalocean_spaces_key.key", "grant.0.permission", "read"),
					resource.TestCheckResourceAttr(
						"digitalocean_spaces_key.key", "grant.1.bucket", "my-bucket2"),
					resource.TestCheckResourceAttr(
						"digitalocean_spaces_key.key", "grant.1.permission", "readwrite"),
				),
			},
			{
				Config: testAccDigitalOceanSpacesKeyConfig(expectedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"digitalocean_spaces_key.key", "name", expectedName),
					resource.TestCheckResourceAttr(
						"digitalocean_spaces_key.key", "grant.0.bucket", "my-bucket"),
					resource.TestCheckResourceAttr(
						"digitalocean_spaces_key.key", "grant.0.permission", "read"),
					resource.TestCheckNoResourceAttr(
						"digitalocean_spaces_key.key", "grant.1.bucket"),
					resource.TestCheckNoResourceAttr(
						"digitalocean_spaces_key.key", "grant.1.permission"),
				),
			},
		},
	})
}

func testAccDigitalOceanSpacesKeyConfig(name string) string {
	return fmt.Sprintf(`
resource "digitalocean_spaces_key" "key" {
  name = "%s"
  grant {
    bucket     = "my-bucket"
    permission = "read"
  }
}
`, name)
}

func testAccDigitalOceanSpacesKeyConfigWithGrantUpdate(name string) string {
	return fmt.Sprintf(`
resource "digitalocean_spaces_key" "key" {
  name = "%s"
  grant {
    bucket     = "my-bucket2"
    permission = "read"
  }
}
`, name)
}

func testAccDigitalOceanSpacesKeyConfigMultipleGrants(name string) string {
	return fmt.Sprintf(`
resource "digitalocean_spaces_key" "key" {
  name = "%s"
  grant {
    bucket     = "my-bucket"
    permission = "read"
  }
  grant {
    bucket     = "my-bucket2"
    permission = "readwrite"
  }
}
`, name)
}
