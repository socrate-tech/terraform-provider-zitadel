package idp_github

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zitadel/terraform-provider-zitadel/zitadel/v2/idp_utils"
)

func GetResource() *schema.Resource {
	return &schema.Resource{
		Description: "Resource representing a OIDC IDP on the instance.",
		Schema: map[string]*schema.Schema{
			idp_utils.NameVar: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the IDP",
			},
			idp_utils.ClientIDVar: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "client id generated by the identity provider",
			},
			idp_utils.ClientSecretVar: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "client secret generated by the identity provider",
				Sensitive:   true,
			},
			idp_utils.ScopesVar: {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required:    true,
				Description: "the scopes requested by ZITADEL during the request on the identity provider",
			},
			idp_utils.IsLinkingAllowedVar: {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "enable if users should be able to link an existing ZITADEL user with an external account",
			},
			idp_utils.IsCreationAllowedVar: {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "enable if users should be able to create a new account in ZITADEL when using an external account",
			},
			idp_utils.IsAutoCreationVar: {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "enable if a new account in ZITADEL should be created automatically when login with an external account",
			},
			idp_utils.IsAutoUpdateVar: {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "enable if a the ZITADEL account fields should be updated automatically on each login",
			},
		},
		ReadContext:   read,
		UpdateContext: update,
		CreateContext: create,
		DeleteContext: idp_utils.Delete,
		Importer:      &schema.ResourceImporter{StateContext: idp_utils.ImportIDPWithClientSecret},
	}
}
