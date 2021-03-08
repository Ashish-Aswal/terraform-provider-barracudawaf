provider "barracudawaf" {
    address  = "52.25.122.104"
    username = "admin"
    port     = "8443"
    password = "i-068e115398f17b347"
}

resource "barracudawaf_trusted_server_certificate" "demo_trusted_server_cert_1" {
  name        = "DemoTrustedServerCert1"
  certificate = "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUdFekNDQS91Z0F3SUJBZ0lRZlZ0UkpyUjJ1\naEhiZEJZTHZGTU5wekFOQmdrcWhraUc5dzBCQVF3RkFEQ0IKaURFTE1Ba0dBMVVFQmhNQ1ZWTXhF\nekFSQmdOVkJBZ1RDazVsZHlCS1pYSnpaWGt4RkRBU0JnTlZCQWNUQzBwbApjbk5sZVNCRGFYUjVN\nUjR3SEFZRFZRUUtFeFZVYUdVZ1ZWTkZVbFJTVlZOVUlFNWxkSGR2Y21zeExqQXNCZ05WCkJBTVRK\nVlZUUlZKVWNuVnpkQ0JTVTBFZ1EyVnlkR2xtYVdOaGRHbHZiaUJCZFhSb2IzSnBkSGt3SGhjTk1U\nZ3gKTVRBeU1EQXdNREF3V2hjTk16QXhNak14TWpNMU9UVTVXakNCanpFTE1Ba0dBMVVFQmhNQ1Iw\nSXhHekFaQmdOVgpCQWdURWtkeVpXRjBaWElnVFdGdVkyaGxjM1JsY2pFUU1BNEdBMVVFQnhNSFUy\nRnNabTl5WkRFWU1CWUdBMVVFCkNoTVBVMlZqZEdsbmJ5Qk1hVzFwZEdWa01UY3dOUVlEVlFRREV5\nNVRaV04wYVdkdklGSlRRU0JFYjIxaGFXNGcKVm1Gc2FXUmhkR2x2YmlCVFpXTjFjbVVnVTJWeWRt\nVnlJRU5CTUlJQklqQU5CZ2txaGtpRzl3MEJBUUVGQUFPQwpBUThBTUlJQkNnS0NBUUVBMW5NejF0\nYzhJTkFBMGhkRnVOWStCNkkveDBIdU1qREpzR3o5OUovTEVwZ1BMVCtOClRRRU1nZzhYZjJJdTZi\naEllZnNXZzA2dDF6SWxrN2NIdjdsUVA2bE13MEFxNlRuLzJZSEtIeFl5UWRxQUpya2oKZW9jZ0h1\nUC9JSm84bFVSdmgzVUdrRUMwTXBNV0NSQUlJejdTM1ljUGIxMVJGR29LYWNWUEFYSnB6OU9UVEcw\nRQpvS01iZ242eG1ybnR4WjdGTjNpZm1nZzArMVl1V01RSkRnWmtXN3czM1BHZktHaW9WckNTbzF5\nZnU0aVlDQnNrCkhhc3doYTZ2c0M2ZWVwM0J3RUljNGdMdzZ1QkswdStRRHJUQlFCYndiNFZDU21U\nM3BEQ2cvcjh1b3lkYWpvdFkKdUszREdSZUVZKzF2VnYyRHkyQTB4SFMrNXAzYjRlVGx5Z3hmRlFJ\nREFRQUJvNElCYmpDQ0FXb3dId1lEVlIwagpCQmd3Rm9BVVUzbS9XcW9yU3M5VWdPSFltOENkOHJJ\nRFpzc3dIUVlEVlIwT0JCWUVGSTJNWHNSVXJZcmhkK21iCitac0Y0YmdCaldIaE1BNEdBMVVkRHdF\nQi93UUVBd0lCaGpBU0JnTlZIUk1CQWY4RUNEQUdBUUgvQWdFQU1CMEcKQTFVZEpRUVdNQlFHQ0Nz\nR0FRVUZCd01CQmdnckJnRUZCUWNEQWpBYkJnTlZIU0FFRkRBU01BWUdCRlVkSUFBdwpDQVlHWjRF\nTUFRSUJNRkFHQTFVZEh3UkpNRWN3UmFCRG9FR0dQMmgwZEhBNkx5OWpjbXd1ZFhObGNuUnlkWE4w\nCkxtTnZiUzlWVTBWU1ZISjFjM1JTVTBGRFpYSjBhV1pwWTJGMGFXOXVRWFYwYUc5eWFYUjVMbU55\nYkRCMkJnZ3IKQmdFRkJRY0JBUVJxTUdnd1B3WUlLd1lCQlFVSE1BS0dNMmgwZEhBNkx5OWpjblF1\nZFhObGNuUnlkWE4wTG1OdgpiUzlWVTBWU1ZISjFjM1JTVTBGQlpHUlVjblZ6ZEVOQkxtTnlkREFs\nQmdnckJnRUZCUWN3QVlZWmFIUjBjRG92CkwyOWpjM0F1ZFhObGNuUnlkWE4wTG1OdmJUQU5CZ2tx\naGtpRzl3MEJBUXdGQUFPQ0FnRUFNcjlodlE1SXcwL0gKdWtkTitKeDRHUUhjRXgyQWIvekRjTFJT\nbWpFem1sZFMrekdlYTZUdlZLcUpqVUFYYVBnUkVIelN5ckh4VlliSAo3ck0ya1liMk9WRy9ScjhQ\nb0xxMDkzNUp4Q28yRjU3a2FEbDZyNVJPVm0reWV6dS9Db2E5emNWM0hBTzRPTEdpCkgxOSsyNHJj\nUmtpMmFBclBzclcwNGpUa1o2azRaZ2xlMHJqOG5TZzZGMEFud25KT0tmMGhQSHpQRS91V0xNVXgK\nUlAwVDdkV2JxV2xvZDN6dTRmK2srVFk0Q0ZNNW9vUTBuQm56dmc2czFTUTM2eU9vZU5EVDUrK1NS\nMlJpT1NMdgp4dmNSdmlLRnhtWkVKQ2FPRURLTnlKT3VCNTZEUGkvWitmVkdqbU8rd2VhMDNLYk5J\nYWlHQ3BYWkxvVW1HdjM4CnNiWlhRbTJWMFRQMk9SUUdna0U0OVk5WTNJQmJwTlY5bFhqOXA1di8v\nY1dvYWFzbTU2ZWtCWWRicWJlNG95QUwKbDZsRmhkMnppK1dKTjQ0cERmd0dGL1k0UUE1QzVCSUcr\nM3Z6eGhGb1l0L2ptUFFUMkJWUGk3RnAyUkJndkdRcQo2akczNUxXak9oU2JKdU1MZS8wQ2pyYVp3\nVGlYV1RiMnFIU2loclplNjhaazZzK2dvL2x1bnJvdEViYUdtQWhZCkxjbXNKV1R5WG5XME9NR3Vm\nMXBHZytwUnlyYnhtUkUxYTZWcWU4WUFzT2Y0dm1TeXJjakM4YXpqVWVxa2srQjUKeU9HQlFNa0tX\nK0VTUE1GZ0t1T1h3SWxDeXBUUFJwZ1NhYnVZME1MVERYSkxSMjdsazhReUtHT0hRK1N3TWo0Swow\nMHUvSTVzVUtVRXJtZ1Fma3kzeHh6bElQSzFhRW44PQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0t\n"
}

resource "barracudawaf_trusted_ca_certificate" "demo_trusted_ca_cert_1" {
  name        = "DemoTrustedCACert1"
  certificate = "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUYzakNDQThhZ0F3SUJBZ0lRQWYxdE1QeWp5\nbEdvRzd4a0RqVURMVEFOQmdrcWhraUc5dzBCQVF3RkFEQ0IKaURFTE1Ba0dBMVVFQmhNQ1ZWTXhF\nekFSQmdOVkJBZ1RDazVsZHlCS1pYSnpaWGt4RkRBU0JnTlZCQWNUQzBwbApjbk5sZVNCRGFYUjVN\nUjR3SEFZRFZRUUtFeFZVYUdVZ1ZWTkZVbFJTVlZOVUlFNWxkSGR2Y21zeExqQXNCZ05WCkJBTVRK\nVlZUUlZKVWNuVnpkQ0JTVTBFZ1EyVnlkR2xtYVdOaGRHbHZiaUJCZFhSb2IzSnBkSGt3SGhjTk1U\nQXcKTWpBeE1EQXdNREF3V2hjTk16Z3dNVEU0TWpNMU9UVTVXakNCaURFTE1Ba0dBMVVFQmhNQ1ZW\nTXhFekFSQmdOVgpCQWdUQ2s1bGR5QktaWEp6WlhreEZEQVNCZ05WQkFjVEMwcGxjbk5sZVNCRGFY\nUjVNUjR3SEFZRFZRUUtFeFZVCmFHVWdWVk5GVWxSU1ZWTlVJRTVsZEhkdmNtc3hMakFzQmdOVkJB\nTVRKVlZUUlZKVWNuVnpkQ0JTVTBFZ1EyVnkKZEdsbWFXTmhkR2x2YmlCQmRYUm9iM0pwZEhrd2dn\nSWlNQTBHQ1NxR1NJYjNEUUVCQVFVQUE0SUNEd0F3Z2dJSwpBb0lDQVFDQUVtVVhOZzdEMndpejBL\neFhEWGJ0elNmVFRLMVFnMkhpcWlCTkNTMWtDZHpPaVovTVBhbnM5cy9CCjNQSFRzZFo3TnlnUksw\nZmFPY2E4T2htMFg2YTlmWjJqWTBLMmR2S3BPeXVSK09KdjBPd1dJSkFKUHVMb2RNa1kKdEpIVVlt\nVGJmNk1HOFlnWWFwQWlQTHorRS9DSEZIdjI1QitPMU9SUnhoRm5SZ2hSeTRZVVZEKzhNLzUrYkp6\nLwpGcDBZdlZHT05hYW5ac2h5WjlzaFpySFVtM2dEd0ZBNjZNenczTHllVFA2dkJaWTFIMWRhdC8v\nTytUMjNMTGIyClZOM0k1eEk2VGE1TWlyZGNtclMzSUQzS2Z5STBybjQ3YUdZQlJPY0JUa1pUbXpO\nZzk1UytVemVRYzBQek1zTlQKNzl1cS9uUk9hY2RyakdDVDNzVEhETi9oTXE3TWt6dFJlSlZuaSs0\nOVZ2NE0wR2tQR3cvekpTWnJNMjMzYmtmNgpjMFBsZmc2bFpyRXBmREtFWTFXSnhBM0JrMVF3R1JP\nczAzMDNwK3RkT213MVhOdEIxeExhcVVrTDM5aUFpZ21UCllvNjFaczhsaU0yRXVMRS9wRGtQMlFL\nZTZ4Sk1sWHp6YXdXcFhoYUR6TGhuNHVnVG5jeGJndE5NcysxYi85N2wKYzZ3ak95MEF2elZWZEFs\nSjJFbFlHbitTTnVaUmtnN3pKbjBjVFJlOHlleERKdEMvUVY5QXFVUkU5Sm5uVjRlZQpVQjlYVktn\nKy9YUmpMN0ZRWlFubVdFSXVReHBNdFBBbFIxbjZCQjZUMUNaR1NsQ0JzdDYrZUxmOFp4WGh5VmVF\nCkhnOWoxdWxpdXRaZlZTN3FYTVlvQ0FRbE9iZ09LNm55VEpjY0J6OE5Vdlh0N3krQ0R3SURBUUFC\nbzBJd1FEQWQKQmdOVkhRNEVGZ1FVVTNtL1dxb3JTczlVZ09IWW04Q2Q4cklEWnNzd0RnWURWUjBQ\nQVFIL0JBUURBZ0VHTUE4RwpBMVVkRXdFQi93UUZNQU1CQWY4d0RRWUpLb1pJaHZjTkFRRU1CUUFE\nZ2dJQkFGelVmQTNQOXdGOVFabGxESFBGClVwL0wrTStaQm44YjJrTVZuNTRDVlZlV0ZQRlNQQ2VI\nbENqdEh6b0JONkoyL0ZOUXdJU2J4bXRPdW93aFQ2S08KVldLUjgya1YyTHlJNDhTcUMvM3ZxT2xM\nVlNvR0lHMVZlQ2taN2w4d1hFc2tFVlgvSkpwdVhpb3I3Z3RObjMvMwpBVGlVRkpWREJ3bjdZS251\nSEtzU2pLQ2FYcWVZYWxsdGl6OEkrOGpSUmE4WUZXU1FFZzl6S0M3RjRpUk8vRmpzCjhQUkYvaUt6\nNnkrTzB0bEZZUVhCbDIrb2RuS1BpNHcycjc4TkJjNXhqZWFtYng5c3BuRml4ZGpRZzNJTThXY1IK\naVF5Y0UweHlOTis4MVhIZnFuSGQ0YmxzakR3U1hXWGF2VmNTdGtOci8rWGVUV1lSVWMrWnJ1d1h0\ndWh4a1l6ZQpTZjdkTlhHaUZTZVVITTloNHlhN2I2Tm5KU0ZkNXQwZEN5NW9HenVDcit5RFo0WFVt\nRkYwc2JtWmdJbi9mM2daClhIbEtZQzZTUUs1TU55b3N5Y2RpeUE1ZDl6WmJ5dUFsSlFHMDNSb0hu\nSGNBUDlEYzFldzkxUHE3UDh5RjFtOS8KcVMzZnVRTDM5WmVhdFRYYXcyZXdoMHFwS0o0amp2OWNK\nMnZoc0UvekIrNEFMdFJaaDh0U1FaWHE5RWZYN21SQgpWWHlOV1FLVjNXS2R3cm51V2loMGhLV2J0\nNURIREFmZjlZazJkRExXS01Hd3NBdmduRXpESE5iODQybTFSMGFCCkw2S0NxOU5qUkhERWpmOHRN\nN3F0ajN1MWNJaXVQaG5QUUNqWS9NaVF1MTJaSXZWUzVsakZINGd4USs2SUhkZkcKamp4RGFoMm5H\nTjU5UFJieFl2bktrS2o5Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K\n"
  depends_on  = [ barracudawaf_trusted_server_certificate.demo_trusted_server_cert_1 ]
}

resource "barracudawaf_self_signed_certificate" "demo_self_signed_cert_1" {
    name                     = "DemoSelfSignedCert1"
    allow_private_key_export = "Yes"
    city                     = "Bangalore"
    common_name              = "waf.test.local"
    country_code             = "IN"
    key_size                 = "1024"
    key_type                 = "rsa"
    organization_name        = "Barracuda Networks"
    organizational_unit      = "Engineering"
    state                    = "Karnataka"
    depends_on               = [barracudawaf_trusted_ca_certificate.demo_trusted_ca_cert_1]
}

resource "barracudawaf_services" "demo_app_1" {
    name            = "DemoApp1"
    ip_address      = "172.31.89.13"
    port            = "80"
    type            = "HTTP"
    vsite           = "default"
    address_version = "IPv4"
    status          = "On"
    group           = "default"
    comments        = "Demo Service with Terraform"

    basic_security {
      mode = "Active"
    }

    depends_on = [ barracudawaf_self_signed_certificate.demo_self_signed_cert_1 ]
}

resource "barracudawaf_servers" "demo_server_1" {
    name            = "DemoServer1"
    identifier      = "IP Address"
    address_version = "IPv4"
    status          = "In Service"
    ip_address      = "104.43.130.86"
    port            = "80"
    comments        = "Creating the Demo Server"
    parent          = [ barracudawaf_services.demo_app_1.name ]
    
    out_of_band_health_checks {
      enable_oob_health_checks = "Yes"
      interval                 = "900"
    }

    depends_on      = [ barracudawaf_services.demo_app_1 ]
}

resource "barracudawaf_services" "demo_app_2" {
    name            = "DemoApp2"
    ip_address      = "172.31.89.13"
    port            = "443"
    type            = "HTTPS"
    vsite           = "default"
    address_version = "IPv4"
    status          = "On"
    group           = "default"
    comments        = "Demo Service with Terraform"
    certificate     = barracudawaf_self_signed_certificate.demo_self_signed_cert_1.name

    basic_security {
      mode = "Active"
    }

    depends_on = [ barracudawaf_servers.demo_server_1 ]
}

resource "barracudawaf_servers" "demo_server_2" {
    name            = "TestServer2"
    identifier      = "IP Address"
    address_version = "IPv4"
    status          = "In Service"
    ip_address      = "104.43.130.86"
    port            = "80"
    comments        = "Creating the Demo Server"
    parent          = [ barracudawaf_services.demo_app_2.name ]

    out_of_band_health_checks {
      enable_oob_health_checks = "Yes"
      interval                 = "900"
    }

    depends_on = [ barracudawaf_services.demo_app_2 ]
}

resource "barracudawaf_security_policies" "demo_security_policy_1" {
    name       = "DemoPolicy1"
    based_on   = "Create New"
    
    depends_on = [ barracudawaf_servers.demo_server_2 ]
}

resource "barracudawaf_content_rules" "demo_rule_group_1" {
    name                = "DemoRuleGroup1"
    url_match           = "/testing.html"
    host_match          = "www.example.com"
    web_firewall_policy = "DemoPolicy1"
    mode                = "Active"
    parent              = [ barracudawaf_services.demo_app_1.name ]
    
    depends_on          = [ barracudawaf_security_policies.demo_security_policy_1 ]
}
 
resource "barracudawaf_content_rule_servers" "demo_rule_group_server_1" {
    name        = "DemoRuleGroupServer1"
    identifier  = "Hostname"
    hostname    = "barracuda.com"
    parent      = [ barracudawaf_services.demo_app_1.name, barracudawaf_content_rules.demo_rule_group_1.name ]
    

    application_layer_health_checks {
        method               = "POST"
        match_content_string = "index"
        domain               = "example.com"
    }

    depends_on = [ barracudawaf_content_rules.demo_rule_group_1 ]
}

resource "barracudawaf_url_acls" "demo_url_acl_1" {
    name         = "DemoUrlAcl1"
    redirect_url = "http://www.example.com/index.html"
    action       = "Allow and Log"
    parent       = [ barracudawaf_services.demo_app_1.name ]
    
    depends_on   = [ barracudawaf_content_rule_servers.demo_rule_group_server_1 ]
}