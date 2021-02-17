# terraform-provider-barracudawaf #

Manage Barracuda WAF config resources

## Requirements ##
* Terraform v0.14.x
* Go 1.15 (to build the provider plugin)

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

Checkout the **main.tf** for more under **examples**



## Develop The Provider ##

### Dependencies for building from source ###

If you need to build from source, you should have a working Go environment setup.  If not check out the Go [getting started](http://golang.org/doc/install) guide.

This project uses [Go Modules](https://github.com/golang/go/wiki/Modules) for dependency management.  To fetch all dependencies run `go get` inside this repository.


### Build ###

```go
git clone https://github.com/Ashish-Aswal/terraform-provider-barracudawaf.git

cd $HOME/terraform-provider-barracudawaf/

go build -o terraform-provider-barracudawaf
```


### Install ###

```sh
mkdir $HOME/.terraform.d/plugins/registry.terraform.io/hashicorp/barracudawaf/0.0.1/darwin_amd64/ 
mv terraform-provider-barracudawaf $HOME/.terraform.d/plugins/registry.terraform.io/hashicorp/barracudawaf/0.0.1/darwin_amd64/terraform-provider-barracudawaf
```

This will place the binary under `$HOME/.terraform.d/plugins/registry.terraform.io/hashicorp/barracudawaf/0.0.1/darwin_amd64/`.  After installing you will need to run `terraform init` in any project using the plugin.


## This project is auto-generated, incase of issues please reachout at email: aaswal@barracuda.com ##
