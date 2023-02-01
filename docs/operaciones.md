# Operaciones Google Cloud (WSL2 con gcloud y terraform)

Crear el proyecto

`PROJECT_ID="farma-compara"`

`gcloud projects create ${PROJECT_ID}`

`gcloud config set project ${PROJECT_ID}`

## Asignación de cuenta de facturación

Esto lo he hecho desde la web de Google Cloud en la sección de Facturación (billing) y eligiendo la cuenta que tenía para educación.

`gcloud services enable compute.googleapis.com`

`gcloud config set compute/zone europe-west2-b`

