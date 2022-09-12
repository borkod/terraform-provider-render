variable "render_api_key" {
  type = string
  description = "Render API Key"
  default = ""
}

variable "service_id" {
  type    = string
  description = "Value of the service's id attribute"
  default = ""
}

variable "start_command" {
  type    = string
  description = "Value of the job's start_command attribute"
  default = ""
}

variable "plan_id" {
  type    = string
  description = "Value of the job's plan_id attribute"
  default = ""
}