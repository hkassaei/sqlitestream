#!/bin/bash
# create a k8s namespace for the cars app
kubectl create namespace sqlitestream

# create a k8s secret with aws credentials 
AWS_PROFILE=default 
ACCESS_KEY_ID=$(aws configure get aws_access_key_id --profile $AWS_PROFILE)
SECRET_ACCESS_KEY=$(aws configure get aws_secret_access_key --profile $AWS_PROFILE)
kubectl create secret generic aws-creds -n sqlitestream --from-literal=ACCESS_KEY_ID=$ACCESS_KEY_ID --from-literal=SECRET_ACCESS_KEY=$SECRET_ACCESS_KEY

# create confimap containing litestream config
kubectl create configmap litestream -n sqlitestream --from-file=litestream.yml
