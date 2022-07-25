---
page_title: "Provider: ZITADEL"
description: |-
    Provider to manage resources on ZITADEL

---

# ZITADEL Provider

This provider allows the management of different resources on a ZITADEL instance.

Through this provider it is possible to interact with the [ZITADEL API](https://docs.zitadel.com/docs/apis/introduction) to configure different aspects of the instance.

## Configuring the provider

```terraform
terraform {
  required_providers {
    zitadel = {
      source  = "zitadel/zitadel"
      version = "1.0.0-alpha.4"
    }
  }
}

provider zitadel {
  domain = "localhost"
  insecure = "true"
  port = "8080"
  project = "170832731415117995"
  token   = "local-token"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `domain` (String) Domain used to connect to the ZITADEL instance
- `project` (String) ID of the ZITADEL project on your ZITADEL instance, for the audience-scope
- `token` (String) Path to the file containing credentials to connect to ZITADEL

### Optional

- `insecure` (Boolean) Use insecure connection
- `port` (String) Used port if not the default ports 80 or 443 are configured

## Limitations

The token which is used to connect to ZITADEL is currently limited to JWT-token for serviceaccounts, so you have to create a service account first, to download the key from this.