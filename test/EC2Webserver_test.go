package test

import (
	"fmt"
	"testing"
	"time"

	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestTerraformEC2WebserverExample(t *testing.T) {

	// the values to pass into the module
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{

		// the path where the module is located
		TerraformDir: "../ec2-webserver",

		// variables to pass to the module using -var options
		Vars: map[string]interface{}{
			"region": "us-east-2",
		},
	})

	// run a Terraform Init and Apply with the terraform options
	terraform.InitAndApply(t, terraformOptions)

	// run a Terraform Destroy at the end of the test
	defer terraform.Destroy(t, terraformOptions)

	publicDNS := terraform.Output(t, terraformOptions, "instance_app-server1_public_dns")

	url := fmt.Sprint(publicDNS)

	http_helper.HttpGetWithRetry(t, "http://"+url, nil, 200, "This is a terraform module for EC2", 15, 10*time.Second)
	terraform.OutputRequired(t, terraformOptions, "vpn_id")
	terraform.OutputRequired(t, terraformOptions, "subnet_id")
	terraform.OutputRequired(t, terraformOptions, "route_table_id")
	terraform.OutputRequired(t, terraformOptions, "gateway_id")

}

// ref: https://www.youtube.com/watch?v=GLhtnOdSIh0ssss
