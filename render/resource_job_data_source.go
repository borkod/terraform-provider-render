package render

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	renderRest "github.com/borkod/render-rest-go"
)

func dataSourceJob() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceJobRead,
		Schema: map[string]*schema.Schema{

			"id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"service_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"start_command": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
				Computed: true,
			},
			"plan_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_at": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"started_at": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"finished_at": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceJobRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*renderRest.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	serviceId := d.Get("service_id").(string)
	jobId := d.Get("id").(string)

	job, err := client.GetJob(jobId, serviceId)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(job.Id)
	d.Set("service_id", job.ServiceId)
	d.Set("start_command", job.StartCommand)
	d.Set("plan_id", job.PlanId)
	d.Set("status", job.Status)
	d.Set("created_at", job.CreatedAt)
	d.Set("started_at", job.StartedAt)
	d.Set("finished_at", job.FinishedAt)

	return diags
}
