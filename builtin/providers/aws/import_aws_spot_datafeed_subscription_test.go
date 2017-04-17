package aws

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccAWSSpotDatafeedSubscription_importBasic(t *testing.T) {
	resourceName := "aws_spot_datafeed_subscription.baz"
	ri := acctest.RandIntRange(1, 50000)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSSpotDatafeedSubscriptionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAWSSpotDatafeedSubscription_import(ri),
			},

			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccAWSSpotDatafeedSubscription_import(randInt int) string {
	return fmt.Sprintf(`
resource "aws_s3_bucket" "baz" {
	bucket = "tf-spot-datafeed-disappears-%d"
}

resource "aws_spot_datafeed_subscription" "baz" {
	bucket = "${aws_s3_bucket.baz.bucket}"
}
`, randInt)
}
