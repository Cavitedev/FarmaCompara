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

## Web

Se crea una aplicación con [Svelte Kit](https://kit.svelte.dev/) en la carpeta `web` con el comando:

`npm create svelte@latest farma-compara`

Seleccionando:

1. "Skeleton project"
2. "Yes, using TypeScript syntax"
3. "Yes" (ESLint)
4. "Yes" (Prettier)
5. "Playwright" (Yes)
6. "Vitest" (Yes)

Y luego se ejecuta `npm install` en la carpeta creada.

El contenido del .gitignore se mueve al .gitignore raíz al compartir proyecto git. Para compartir labels y organizar todo en un mismo proyecto. Normalmente se crearían 2 proyectos separados, pero al tratarse del TFG individual combinaré todo en un solo proyecto para que se vea claro el desarrollo y el historial de cambios.

## Firebase

Crear un proyecto en https://firebase.google.com/ asociado al que se creó en Google Cloud y habilitar análiticas.

Añadir app web e inicializarla.
