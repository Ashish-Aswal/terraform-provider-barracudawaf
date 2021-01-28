provider "barracudawaf" {
    ip = "52.25.122.104"
    username = "admin"
    admin_port = "8000"
    password = "i-068e115398f17b347"
}

resource "barracudawaf_service" "DemoService1" {
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

resource "barracudawaf_server" "TestServer1" {
    name = "TestServer1"
    identifier= "IP Address"
    address_version = "IPv4"
    status = "In Service"
    ip_address = "107.191.119.130"
    service_name = "DemoService1"
    port = "80"
    comments = "Creating the Demo Server"
    depends_on = [barracudawaf_service.DemoService1]
}

resource "barracudawaf_certificate" "DemoSelfSignedCert1" {
    name  = "DemoSelfSignedCert1"
    allow_private_key_export = "Yes"
    city   = "Bangalore"
    common_name = "waf.test.local"
    country_code = "IN"
    curve_type = "secp256r1"
    key_size = "1024"
    key_type = "rsa"
    organization_name = "Barracuda Networks"
    organization_unit = "Engineering"
    state = "Karnataka"
    depends_on = [barracudawaf_server.TestServer1]
}

resource "barracudawaf_service" "DemoService2" {
    name = "DemoService2"
    ip_address = "172.31.49.71"
    port = "443"
    type = "HTTPS"
    vsite = "default"
    address_version = "IPv4"
    status = "On"
    group = "default"
    certificate = "DemoSelfSignedCert1"
    comments = "Demo Service with Terraform"
    depends_on = [barracudawaf_certificate.DemoSelfSignedCert1]
}

resource "barracudawaf_server" "TestServer2" {
    name = "TestServer2"
    identifier= "IP Address"
    address_version = "IPv4"
    status = "In Service"
    ip_address = "107.191.119.130"
    service_name = "DemoService2"
    port = "80"
    comments = "Creating the Demo Server"
    depends_on = [barracudawaf_service.DemoService2]
}

resource "barracudawaf_security_policy" "DemoPolicy1" {
    name = "DemoPolicy1"
    based_on = "Create New"
    depends_on = [barracudawaf_server.TestServer2]
}

resource "barracudawaf_rule_group" "DemoRuleGroup1" {
    name = "DemoRuleGroup1"
    service_name = "DemoService2"
    url_match = "/testing.html"
    host_match = "www.example.com"
    web_firewall_policy = "DemoPolicy1"
    depends_on = [barracudawaf_security_policy.DemoPolicy1]
}
 
resource "barracudawaf_rule_group_server" "DemoRgServer1" {
    name = "DemoRgServer1"
    service_name = "DemoService2"
    rule_group_name = "DemoRuleGroup1"
    identifier = "Hostname"
    hostname = "imdb.com"
    depends_on = [barracudawaf_rule_group.DemoRuleGroup1]
}