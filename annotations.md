I started the test learning about the S3 terraform resource:

https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket


And then I went to:

https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket_object 

https://stackoverflow.com/questions/55833976/terraform-resource-s3-upload-file-is-not-updated

https://www.terraform.io/docs/language/functions/timestamp.html

https://stackoverflow.com/questions/53764505/terraform-creating-multiple-buckets

https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/s3_bucket_objects


So I created the base of the code:

```
provider "aws" {
    profile = "default"
    region = "us-east-1"
}

variable "s3_object_name" {
    type = list
    default = ["test1.txt", "test2.txt"]
}

resource "aws_s3_bucket" "bucket" {
    bucket = "s3-flugel-lab"
    acl = "private"
}

resource "aws_s3_bucket_object" "object" {
    bucket = "s3-flugel-lab"
    count = "${length(var.s3_object_name)}"
    key    = "${var.s3_object_name[count.index]}"
    acl    = "private"
    content = "${timestamp()}"
    depends_on = [aws_s3_bucket.bucket]
}

# output "hello_world" {
#   value = "Hello, World!"
# }

output "name" {
  value = "${aws_s3_bucket.bucket.bucket}"
}

output "region" {
  value = "${aws_s3_bucket.bucket.region}"
}

output "object1" {
  value = "${aws_s3_bucket_object.object[0].key}"
}

output "object2" {
  value = "${aws_s3_bucket_object.object[1].key}"
}
```

---

I never use the Terratest, then I started with:

https://terratest.gruntwork.io/docs/getting-started/quick-start/

https://stackoverflow.com/questions/60377803/terratest-use-mock-aws-services

https://blog.octo.com/en/test-your-infrastructure-code-with-terratest/

https://github.com/gruntwork-io/terratest/blob/master/modules/aws/s3.go

https://www.terraform.io/docs/language/values/outputs.html

https://stackoverflow.com/questions/49641484/how-to-define-execution-order-without-trigger-option-in-terraform

https://www.terraform.io/docs/language/meta-arguments/depends_on.html


---

About the Github Actions, I started with:

https://learn.hashicorp.com/tutorials/terraform/github-actions

https://medium.com/@petriautero/automate-terraform-testing-with-github-actions-and-terratest-78d74331fdf8

