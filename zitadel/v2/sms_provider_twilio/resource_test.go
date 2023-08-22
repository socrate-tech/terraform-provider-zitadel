package sms_provider_twilio_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/zitadel/zitadel-go/v2/pkg/client/zitadel/admin"

	"github.com/zitadel/terraform-provider-zitadel/zitadel/v2/helper/test_utils"
	"github.com/zitadel/terraform-provider-zitadel/zitadel/v2/sms_provider_twilio"
)

func TestAccSMSProviderTwilio(t *testing.T) {
	frame := test_utils.NewInstanceTestFrame(t, "zitadel_sms_provider_twilio")
	resourceExample, exampleAttributes := test_utils.ReadExample(t, test_utils.Resources, frame.ResourceType)
	exampleProperty := test_utils.AttributeValue(t, sms_provider_twilio.SenderNumberVar, exampleAttributes).AsString()
	exampleSecret := test_utils.AttributeValue(t, sms_provider_twilio.TokenVar, exampleAttributes).AsString()
	test_utils.RunLifecyleTest(
		t,
		frame.BaseTestFrame,
		nil,
		test_utils.ReplaceAll(resourceExample, exampleProperty, exampleSecret),
		exampleProperty, "987654321",
		exampleSecret, "updatedSecret",
		false,
		checkRemoteProperty(*frame),
		test_utils.ZITADEL_GENERATED_ID_REGEX,
		test_utils.CheckNothing,
		nil, nil, "", sms_provider_twilio.TokenVar,
	)
}

func checkRemoteProperty(frame test_utils.InstanceTestFrame) func(string) resource.TestCheckFunc {
	return func(expect string) resource.TestCheckFunc {
		return func(state *terraform.State) error {
			resp, err := frame.GetSMSProvider(frame, &admin.GetSMSProviderRequest{Id: frame.State(state).ID})
			if err != nil {
				return fmt.Errorf("getting sms provider failed: %w", err)
			}
			actual := resp.GetConfig().GetTwilio().GetSenderNumber()
			if actual != expect {
				return fmt.Errorf("expected %s, but got %s", expect, actual)
			}
			return nil
		}
	}
}