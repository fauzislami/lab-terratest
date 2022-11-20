output "instance_app-server1_public_dns" {
  value = aws_instance.webserver.public_dns
}

output "vpn_id" {
  value = aws_vpc.vpc1.id
}

output "subnet_id" {
  value = aws_subnet.private1.id

}

output "route_table_id" {
  value = aws_route_table.this-rt.id
}

output "gateway_id" {
  value = aws_internet_gateway.this-igw.id
}