# terraform-provider-render
The Terraform provider for managing Render resources.

## Build provider
Implementation of the provider is in the `/render` directory.

Run the following command to build and install the provider

```shell
$ make install
```

## Test sample configuration

First, build and install the provider.

```shell
$ make install
```

Then, navigate to one of the examples in the `examples` directory. 

You will have to provide variables as specified in the `variables.tf` files.  This can be done using a `.tfvars` [variable definitions file](https://www.terraform.io/language/values/variables#variable-definitions-tfvars-files) file.

Run the following command to initialize the workspace:

```shell
$ terraform init
```

In order to apply the examples, you must provide your [Render API key](https://render.com/docs/api#creating-an-api-key) using `RENDER_API_KEY` environment variable:

```shell
RENDER_API_KEY=<your api key> terraform apply
```