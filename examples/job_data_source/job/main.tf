terraform {
  required_providers {
    render = {
      version = "0.1.0"
      source  = "terraform.local/render/services"
    }
  }
}

data "render_job" "job" {
  service_id = var.service_id
  id = var.job_id
}

output "job" {
  value = data.render_job.job
}

