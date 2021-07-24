package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/magiconair/properties/assert"
)

// An example of how to test the Terraform module in examples/terraform-aws-s3-example using Terratest.
func TestTerraformAwsS3FlugelBucket(t *testing.T) {
	t.Parallel()

	expectedName := "s3-flugel-lab"
	expectedObject1 := "test1.txt"
	expectedObject2 := "test2.txt"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../",
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	// output := terraform.Output(t, terraformOptions, "hello_world")
	name := terraform.Output(t, terraformOptions, "name")
	region := terraform.Output(t, terraformOptions, "region")
	object1 := terraform.Output(t, terraformOptions, "object1")
	object2 := terraform.Output(t, terraformOptions, "object2")

	aws.AssertS3BucketExists(t, region, name)
	aws.GetS3ObjectContents(t, region, name, object1)
	aws.GetS3ObjectContents(t, region, name, object2)

	// assert.Equal(t, "Hello, World!", output)
	assert.Equal(t, expectedName, name)
	assert.Equal(t, expectedObject1, object1)
	assert.Equal(t, expectedObject2, object2)
}
