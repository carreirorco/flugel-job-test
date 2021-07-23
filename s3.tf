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