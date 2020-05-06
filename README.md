# find-go-cloud-function
simple repo hosting a Google Cloud Function that looks for a static file part of the project

# deployment
`gcloud functions deploy find --entry-point Find --region europe-west1 --runtime go113 --trigger-http --timeout 300 --memory 256MB --allow-unauthenticated`

# Findings
Static file(s) resides in `/workspace/serverless_function_source_code/`
