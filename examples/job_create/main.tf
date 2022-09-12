terraform {
  required_providers {
    render = {
      version = "0.1.0"
      source  = "terraform.local/render/services"
    }
  }
}

provider "render" {
  api_key = var.render_api_key
}

resource "render_job" "job" {
  service_id = var.service_id
  start_command = var.start_command
}

resource "render_job" "job_plan" {
  service_id = var.service_id
  start_command = "echo 'Plan id is set to ${var.plan_id}'"
  plan_id = var.plan_id
}

output "job" {
  value = render_job.job
}

output "job_plan" {
  value = render_job.job_plan
}

