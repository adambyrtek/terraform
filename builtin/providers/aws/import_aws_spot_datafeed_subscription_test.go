package aws

import (
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccAWSSpotDatafeedSubscription_importBasic(t *testing.T) {
	resourceName := "aws_spot_datafeed_subscription.foo"
	ri := acctest.RandIntRange(1, 50000)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAWSSpotDatafeedSubscriptionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAWSSpotDatafeedSubscription_basic(ri),
			},

			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
