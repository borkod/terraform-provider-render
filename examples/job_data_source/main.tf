terraform {
  required_providers {
    render = {
      version = "0.1.0"
      source  = "terraform.local/render/services"
    }
  }
}

module "get_job_info" {
  source = "./job"

  service_id = var.service_id
  job_id = var.job_id
}

output "get_job_info" {
  value = module.get_job_info.job
}
