# Overview #

A [Terraform](terraform.io) provider for Barracuda Web Application Firewall.

&nbsp;
## Requirements ##
-	[Terraform](https://www.terraform.io/downloads.html) v0.14.x
-	[Go](https://golang.org/doc/install) 1.15 (to build the provider plugin)

&nbsp;
## Usage ##

**Use provider**
```hcl
provider "barracudawaf" {
    address  = "x.x.x.x"
    username = "xxxxxxx"
    port     = "8443"
    password = "xxxxxxx"
}
```
**Create Service**
```hcl
resource "barracudawaf_services" "DemoService1" {
    name = "DemoService1"
    ip_address = "x.x.x.x"
    port = "80"
    type = "HTTP"
    vsite = "default"
    address_version = "IPv4"
    status = "On"
    group = "default"
    comments = "Demo Service with Terraform"
}
```
**Create Servers**
```hcl
resource "barracudawaf_servers" "TestServer1" {
    name = "TestServer1"
    identifier= "IP Address"
    address_version = "IPv4"
    status = "In Service"
    ip_address = "x.x.x.x"
    port = "80"
    comments = "Creating the Demo Server"
    parent = [ "DemoService1" ]
    depends_on = [barracudawaf_services.DemoService1]
}
```
**Create Self Signed Certificates**
```hcl
resource "barracudawaf_self_signed_certificate" "DemoSelfSignedCert1" {
    name  = "DemoSelfSignedCert1"
    allow_private_key_export = "Yes"
    city   = "xxxx"
    common_name = "xxxxx"
    country_code = "xx"
    key_size = "1024"
    key_type = "rsa"
    organization_name = "xxxxxxx"
    organizational_unit = "xxxxxx"
    state = "xxxxxxx"
    depends_on = [barracudawaf_servers.TestServer1]
}
```
**Create Security Policies**
```hcl
resource "barracudawaf_security_policies" "DemoPolicy1" {
    name = "DemoPolicy1"
    based_on = "Create New"
}
```
**Create Rule Groups**
```hcl
resource "barracudawaf_content_rules" "DemoRuleGroup1" {
    name = "DemoRuleGroup1"
    url_match = "/xxxx.xxx"
    host_match = "xxxxxx"
    web_firewall_policy = "DemoPolicy1"
    parent = [ "DemoService1" ]
    depends_on = [barracudawaf_security_policies.DemoPolicy1]
}
```
**Create Rule Group Servers**
```hcl
resource "barracudawaf_content_rule_servers" "DemoRgServer1" {
    name = "DemoRgServer1"
    identifier = "Hostname"
    hostname = "xxxxxx"
    parent = [ "DemoService1", "DemoRuleGroup1" ]
    depends_on = [barracudawaf_content_rules.DemoRuleGroup1]
}
```

&nbsp;&nbsp;
## Building The Provider ##

### Dependencies for building from source ###
If you need to build from source, you should have a working Go environment setup.  If not check out the Go [getting started](http://golang.org/doc/install) guide.

This project uses [Go Modules](https://github.com/golang/go/wiki/Modules) for dependency management.  To fetch all dependencies run `go get` inside this repository.

&nbsp;&nbsp;
### Build ###

Clone repository to: $GOPATH/src/github.com/Ashish-Aswal/terraform-provider-barracudawaf
```shell
$ mkdir -p $GOPATH/src/github.com/Ashish-Aswal; cd $GOPATH/src/github.com/Ashish-Aswal
$ git clone https://github.com/Ashish-Aswal/terraform-provider-barracudawaf.git
```

Enter the provider directory and build the provider
```shell
cd $GOPATH/src/github.com/Ashish-Aswal/terraform-provider-barracudawaf
make build
```

&nbsp;&nbsp;
### Install ###

```shell
$ cd $GOPATH/src/github.com/Ashish-Aswal/terraform-provider-barracudawaf
$ make install

```

&nbsp;&nbsp;
# Using the Provider

If you're building the provider, follow the instructions to install it as a plugin. After placing it into your plugins directory, run terraform init to initialize it.

&nbsp;&nbsp;
# Developing the Provider

If you wish to work on the provider, you'll first need Go installed on your machine (version 1.15 is required). You'll also need to correctly setup a GOPATH, as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run make build. This will create a binary with name `terraform-provider-barracudawaf` in `$GOPATH/src/github.com/Ashish-Aswal/terraform-provider-barracudawaf` directory.

```shell
$ make build
...
$ $GOPATH/src/github.com/Ashish-Aswal/terraform-provider-barracudawaf
...

```

&nbsp;
# Using the binary instead of building it from source #

Download the binary added under [releases](https://github.com/Ashish-Aswal/terraform-provider-barracudawaf/releases), and follow below :


```shell
$ git clone https://github.com/Ashish-Aswal/terraform-provider-barracudawaf.git

```

Copy the downloded binary into `terraform-provider-barracudawaf` directory created with abvoe git clone command.
```shell
cd terraform-provider-barracudawaf/
make plugin
```
