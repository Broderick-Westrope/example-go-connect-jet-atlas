schema "public" {}

table "users" {
  schema = schema.public
  column "id" {
    type = bigserial
    null = false
  }
  column "email" {
    type = text
    null = false
  }
  column "name" {
    type = text
    null = false
  }
  column "created_at" {
    type    = timestamp
    null    = false
    default = sql("now()")
  }
  column "updated_at" {
    type    = timestamp
    null    = false
    default = sql("now()")
  }

  primary_key {
    columns = [column.id]
  }

  index "users_email_idx" {
    columns = [column.email]
    unique  = true
  }
}
