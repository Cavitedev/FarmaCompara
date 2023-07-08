# Operaciones Google Cloud (WSL2 con gcloud y terraform)

Crear el proyecto

`PROJECT_ID="farma-compara"`

`gcloud projects create ${PROJECT_ID}`

`gcloud config set project ${PROJECT_ID}`

## Asignación de cuenta de facturación

Esto lo he hecho desde la web de Google Cloud en la sección de Facturación (billing) y eligiendo la cuenta que tenía para educación.

## Añadir región de computación y servicios

`gcloud services enable compute.googleapis.com`

`gcloud config set compute/zone europe-west2-b`

## Añadir otros servicios

`gcloud services enable cloudfunctions.googleapis.com`

`gcloud services enable run.googleapis.com`

`gcloud services enable artifactregistry.googleapis.com`

`gcloud services enable cloudbuild.googleapis.com`

## Crear bucket para terraform

`BACKEND_BUCKET="farmacompara_terraform_dev"`

El mismo que en terraform/dev.tfbackend

```sh
gcloud storage buckets create gs://$BACKEND_BUCKET \
--location europe-west2
```

## Terraform

`cd terraform`

Ir a la carpeta de terraform

`terraform init  -backend-config=dev.tfbackend`

`terraform plan -var-file=prod.tfvars -out build.tfplan`

`terraform apply "build.tfplan"`



## Firebase

Crear un proyecto en https://firebase.google.com/ asociado al que se creó en Google Cloud.

ir a la carpeta de firebase

`cd firebase`

y hacer un deploy con la configuración allí establecida

`firebase deploy`

## Flutter

