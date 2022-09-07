package render

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	renderRest "github.com/borkod/render-rest-go"
)

func resourceJob() *schema.Resource {
	return &schema.Resource{
		ReadContext:   resourceJobRead,
		CreateContext: resourceJobCreate,
		UpdateContext: resourceJobCreate, // TODO: implement update
		DeleteContext: resourceJobDelete, // TODO: implement delete
		Schema: map[string]*schema.Schema{

			"id": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
				Computed: true,
			},
			"service_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"start_command": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"plan_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
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

func resourceJobRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

func resourceJobCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*renderRest.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	nj := renderRest.NewJob{
		ServiceId:    d.Get("service_id").(string),
		PlanId:       d.Get("plan_id").(string),
		StartCommand: d.Get("start_command").(string),
	}

	job, err := client.CreateJob(nj)
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

func resourceJobDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}
