locals {
    app_id = "iplookup-dev-${random_id.unique_suffix.hex}"
}

variable "iam_policy_arn" {
    description = "IAM Policy to be attached to role"
    type        = list(string)

    default = [ 
        "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
    ]
}