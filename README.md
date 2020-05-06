# What is it?
Simple Google Cloud Function that shows up file structure when using the Go 1.13 runtime. \
Useful if you need to read config file(s) that come along with your function.

# Deployment
`gcloud functions deploy find --entry-point Find --region europe-west1 --runtime go113 --trigger-http --timeout 300 --memory 256MB --allow-unauthenticated`

# Findings
Static file(s) left under that dir, with source code sent by gcloud command: `/workspace/serverless_function_source_code/`

# Output sample
See that gist: https://gist.github.com/bcollard/103fe9fbbbfee098e0a06a4e3d844e0a
