---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "barracudawaf_servers Resource - terraform-provider-barracudawaf"
subcategory: ""
description: |-
  barracudawaf_servers manages Servers on the Barracuda Web Application Firewall.
---

# barracudawaf_servers (Resource)

`barracudawaf_servers` manages `Servers` on the Barracuda Web Application Firewall.

## Example Usage

```terraform
resource "barracudawaf_servers" "web_server_1" {
    name            = "DemoWebServer1"
    identifier      = "IP Address"
    address_version = "IPv4"
    ip_address      = "x.x.x.x"
    port            = "80"
    comments        = "Demo web server behind DemoApp1"
    parent          = [ "DemoApp1" ]
    depends_on      = [barracudawaf_services.application_1]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **parent** (List of String)

### Optional

- **address_version** (String) Version
- **advanced_configuration** (Block List) (see [below for nested schema](#nestedblock--advanced_configuration))
- **application_layer_health_checks** (Block List) (see [below for nested schema](#nestedblock--application_layer_health_checks))
- **comments** (String) Comments
- **connection_pooling** (Block List) (see [below for nested schema](#nestedblock--connection_pooling))
- **hostname** (String) Hostname
- **id** (String) The ID of this resource.
- **identifier** (String) Identifier
- **in_band_health_checks** (Block List) (see [below for nested schema](#nestedblock--in_band_health_checks))
- **ip_address** (String) Server IP
- **load_balancing** (Block List) (see [below for nested schema](#nestedblock--load_balancing))
- **name** (String) Server Name
- **out_of_band_health_checks** (Block List) (see [below for nested schema](#nestedblock--out_of_band_health_checks))
- **port** (String) Server Port
- **resolved_ips** (String)
- **ssl_policy** (Block List) (see [below for nested schema](#nestedblock--ssl_policy))
- **status** (String) Status

### Read-Only

- **redirect** (Block List) (see [below for nested schema](#nestedblock--redirect))

<a id="nestedblock--advanced_configuration"></a>
### Nested Schema for `advanced_configuration`

Optional:

- **client_impersonation** (String) Client Impersonation
- **max_connections** (String) Max Connections
- **max_establishing_connections** (String) Max Establishing Connections
- **max_keepalive_requests** (String) Max Keepalive Requests
- **max_requests** (String) Max Requests
- **max_spare_connections** (String) Max Spare Connections
- **source_ip_to_connect** (String) Source IP to Connect
- **timeout** (String) Timeout


<a id="nestedblock--application_layer_health_checks"></a>
### Nested Schema for `application_layer_health_checks`

Optional:

- **additional_headers** (String) Additional Headers
- **domain** (String) Domain
- **match_content_string** (String) Match content String
- **method** (String) Method
- **status_code** (String) Status Code
- **url** (String) URL


<a id="nestedblock--connection_pooling"></a>
### Nested Schema for `connection_pooling`

Optional:

- **enable_connection_pooling** (String) Enable Connection Pooling
- **keepalive_timeout** (String) Keepalive Timeout


<a id="nestedblock--in_band_health_checks"></a>
### Nested Schema for `in_band_health_checks`

Optional:

- **max_http_errors** (String) Max HTTP Errors
- **max_other_failure** (String) Max Other Failure
- **max_refused** (String) Max Refused
- **max_timeout_failure** (String) Max Timeout Failures


<a id="nestedblock--load_balancing"></a>
### Nested Schema for `load_balancing`

Optional:

- **backup_server** (String) Backup Appliance
- **weight** (String) WRR Weight


<a id="nestedblock--out_of_band_health_checks"></a>
### Nested Schema for `out_of_band_health_checks`

Optional:

- **enable_oob_health_checks** (String) Enable OOB Health Checks
- **interval** (String) Interval


<a id="nestedblock--ssl_policy"></a>
### Nested Schema for `ssl_policy`

Optional:

- **client_certificate** (String) Client Certificate
- **enable_https** (String) Server uses SSL
- **enable_sni** (String) Enable SNI
- **enable_ssl_3** (String) SSL 3.0 (Insecure)
- **enable_ssl_compatibility_mode** (String) Enable SSL Compatibility Mode
- **enable_tls_1** (String) TLS 1.0 (Insecure)
- **enable_tls_1_1** (String) TLS 1.1
- **enable_tls_1_2** (String) TLS 1.2
- **enable_tls_1_3** (String) TLS 1.3
- **validate_certificate** (String) Validate Server Certificate


<a id="nestedblock--redirect"></a>
### Nested Schema for `redirect`


