#!/bin/bash -ex

eksctl get cluster --name txterv8
aws cloudformation list-stacks | jq '.StackSummaries[]|select(.StackStatus!="DELETE_COMPLETE")'
eksctl get nodegroup --region us-east-2 --cluster txterv8
kubectl get po,svc,deploy
