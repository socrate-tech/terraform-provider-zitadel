package org_idp_github

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/zitadel/terraform-provider-zitadel/zitadel/v2/idp_utils"
	"github.com/zitadel/terraform-provider-zitadel/zitadel/v2/org_idp_utils"
)

func GetDatasource() *schema.Resource {
	return &schema.Resource{
		Description: "Datasource representing a GitHub IdP of the organization.",
		Schema: map[string]*schema.Schema{
			idp_utils.IdpIDVar:             idp_utils.IdPIDDataSourceField,
			org_idp_utils.OrgIDVar:         org_idp_utils.OrgIDDatasourceField,
			idp_utils.NameVar:              idp_utils.NameDataSourceField,
			idp_utils.ClientIDVar:          idp_utils.ClientIDDataSourceField,
			idp_utils.ClientSecretVar:      idp_utils.ClientSecretDataSourceField,
			idp_utils.ScopesVar:            idp_utils.ScopesDataSourceField,
			idp_utils.IsLinkingAllowedVar:  idp_utils.IsLinkingAllowedDataSourceField,
			idp_utils.IsCreationAllowedVar: idp_utils.IsCreationAllowedDataSourceField,
			idp_utils.IsAutoCreationVar:    idp_utils.IsAutoCreationDataSourceField,
			idp_utils.IsAutoUpdateVar:      idp_utils.IsAutoUpdateDataSourceField,
		},
		ReadContext: read,
		Importer:    &schema.ResourceImporter{StateContext: org_idp_utils.ImportIDPWithOrgAndSecret(idp_utils.ClientSecretVar)},
	}
}