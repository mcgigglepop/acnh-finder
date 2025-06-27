resource "aws_dynamodb_table" "user_profiles" {
  name           = "UserProfiles"
  billing_mode   = "PAY_PER_REQUEST"
  hash_key       = "user_id"

  attribute {
    name = "user_id"
    type = "S"
  }

  tags = {
    Name = "UserProfiles"
  }
}

resource "aws_dynamodb_table" "fish" {
  name           = "Fish"
  billing_mode   = "PAY_PER_REQUEST"
  hash_key       = "fish_id"

  attribute {
    name = "fish_id"
    type = "S"
  }

  tags = {
    Name = "Fish"
  }
}

resource "aws_dynamodb_table" "user_fish" {
  name           = "UserFish"
  billing_mode   = "PAY_PER_REQUEST"
  hash_key       = "user_id"
  range_key      = "fish_id"

  attribute {
    name = "user_id"
    type = "S"
  }

  attribute {
    name = "fish_id"
    type = "S"
  }

  tags = {
    Name = "UserFish"
  }
}
