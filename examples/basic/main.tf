// Create resource
resource "random_uuid" "this" {}

// Output result
output "uuid" {
  value = random_uuid.this.result
}
