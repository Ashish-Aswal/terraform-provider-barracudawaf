provider "barracudawaf" {
    address = "x.x.x.x"
    username = "xxxxx"
    port = "8443"
    password = "xxxxx"
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