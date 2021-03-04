variable address {}
variable username {}
variable password {}
variable port {}

provider "bigip" {
  address  = var.hostname
  username = var.username
  password = var.password
  port     = var.port
}