# Comandos de ayuda para Terraform

## Subir, actualizar

`terraform plan -var-file=prod.tfvars -out build.tfplan`

`terraform apply "build.tfplan"`

## Destruir, eliminar

`terraform plan -var-file=prod.tfvars --destroy -out destroy.tfplan`

`terraform apply "destroy.tfplan"`

## Google Token

`gcloud auth print-identity-token`