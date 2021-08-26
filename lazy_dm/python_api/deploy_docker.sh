#!/bin/bash

aws ecr-public get-login-password --region us-east-1 | docker login --username AWS --password-stdin public.ecr.aws
docker build -t 154723347372.dkr.ecr.us-east-2.amazonaws.com/lazy-dm-api:latest . -f Dockerfile
aws ecr get-login-password --region us-east-2 | docker login --username AWS --password-stdin 154723347372.dkr.ecr.us-east-2.amazonaws.com
docker push 154723347372.dkr.ecr.us-east-2.amazonaws.com/lazy-dm-api:latest
