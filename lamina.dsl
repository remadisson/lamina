zone "test" {
  cidr = "10.0.0.0/16"
  vlan = 10
  description = "Test-Zone"
}

zone "test2" {
  cidr = "10.1.0.0/16"
  vlan = 200
  description = "Test-Zone"
  parent = "test"
}

device "maxtower" {
  ip = "10.0.0.42"
  mac = "AA:BB:CC:DD:EE:FF"
  zone = "test"
}