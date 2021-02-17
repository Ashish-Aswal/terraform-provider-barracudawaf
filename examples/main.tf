provider "barracudawaf" {
    ip = "52.25.122.104"
    username = "admin"
    admin_port = "8000"
    password = "i-068e115398f17b347"
}

resource "barracudawaf_services" "DemoService1" {
    name = "DemoService1"
    ip_address = "172.31.89.13"
    port = "80"
    type = "HTTP"
    vsite = "default"
    address_version = "IPv4"
    status = "On"
    group = "default"
    comments = "Demo Service with Terraform"
}

resource "barracudawaf_servers" "TestServer1" {
    name = "TestServer1"
    identifier= "IP Address"
    address_version = "IPv4"
    status = "In Service"
    ip_address = "107.191.119.130"
    port = "80"
    comments = "Creating the Demo Server"
    parent = [ "DemoService1" ]
    depends_on = [barracudawaf_services.DemoService1]
}

resource "barracudawaf_self_signed_certificate" "DemoSelfSignedCert1" {
    name  = "DemoSelfSignedCert1"
    allow_private_key_export = "Yes"
    city   = "Bangalore"
    common_name = "waf.test.local"
    country_code = "IN"
    key_size = "1024"
    key_type = "rsa"
    organization_name = "Barracuda Networks"
    organizational_unit = "Engineering"
    state = "Karnataka"
    depends_on = [barracudawaf_servers.TestServer1]
}

resource "barracudawaf_services" "DemoService2" {
    name = "DemoService2"
    ip_address = "172.31.49.71"
    port = "90"
    type = "HTTP"
    vsite = "default"
    address_version = "IPv4"
    status = "On"
    group = "default"
    comments = "Demo Service with Terraform"
    depends_on = [barracudawaf_self_signed_certificate.DemoSelfSignedCert1]
}

resource "barracudawaf_servers" "TestServer2" {
    name = "TestServer2"
    identifier= "IP Address"
    address_version = "IPv4"
    status = "In Service"
    ip_address = "107.191.119.130"
    port = "80"
    comments = "Creating the Demo Server"
    parent = [ "DemoService2" ]
    depends_on = [barracudawaf_services.DemoService2]
}

resource "barracudawaf_security_policies" "DemoPolicy1" {
    name = "DemoPolicy1"
    based_on = "Create New"
    depends_on = [barracudawaf_servers.TestServer2]
}

resource "barracudawaf_content_rules" "DemoRuleGroup1" {
    name = "DemoRuleGroup1"
    url_match = "/testing.html"
    host_match = "www.example.com"
    web_firewall_policy = "DemoPolicy1"
    parent = [ "DemoService2" ]
    depends_on = [barracudawaf_security_policies.DemoPolicy1]
}
 
resource "barracudawaf_content_rule_servers" "DemoRgServer1" {
    name = "DemoRgServer1"
    identifier = "Hostname"
    hostname = "imdb.com"
    parent = [ "DemoService2", "DemoRuleGroup1" ]
    depends_on = [barracudawaf_content_rules.DemoRuleGroup1]
}