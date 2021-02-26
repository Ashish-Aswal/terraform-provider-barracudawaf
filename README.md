# terraform-provider-barracudawaf #

A [Terraform](terraform.io) provider for Barracuda Web Application Firewall.

## Requirements ##
-	[Terraform](https://www.terraform.io/downloads.html) v0.14.x
-	[Go](https://golang.org/doc/install) 1.15 (to build the provider plugin)

<br/><br/>

## Usage ##

**Use provider**
```hcl
provider "barracudawaf" {
    ip = "x.x.x.x"
    username = "admin"
    admin_port = "8000"
    password = "xxxxxxx"
}
```
<br/><br/>
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
<br/><br/>
**Servers**

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

&nbsp;&nbsp;
## Building The Provider ##

### Dependencies for building from source ###
&nbsp;&nbsp;

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
$ make insatll

```

This will place the binary under `$HOME/.terraform.d/plugins/registry.terraform.io/hashicorp/barracudawaf/0.1.0/darwin_amd64/`.  After installing you will need to run `terraform init` in any project using the plugin.

&nbsp;&nbsp;
# Using the Provider

If you're building the provider, follow the instructions to install it as a plugin. After placing it into your plugins directory, run terraform init to initialize it.

&nbsp;&nbsp;
# Developing the Provider

If you wish to work on the provider, you'll first need Go installed on your machine (version 1.15 is required). You'll also need to correctly setup a GOPATH, as well as adding $GOPATH/bin to your $PATH.

To compile the provider, run make build. This will create a binary terraform-provider-barracudawaf in $GOPATH/src/github.com/Ashish-Aswal/terraform-provider-barracudawaf directory.

```shell
$ make build
...
$ $GOPATH/src/github.com/Ashish-Aswal/terraform-provider-barracudawaf
...

```

&nbsp;
# Use binary direclty instead building the provider from source #

Download the binary added under [releases](https://github.com/Ashish-Aswal/terraform-provider-barracudawaf/releases), and follow below :


```shell
$ git clone https://github.com/Ashish-Aswal/terraform-provider-barracudawaf.git

```

Copy the downloded binary into terraform-provider-barracudawaf directory created with aboe git clone command.
```shell
cd terraform-provider-barracudawaf/
make plugin
```