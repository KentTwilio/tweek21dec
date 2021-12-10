#!/bin/bash -ex

eksctl create nodegroup --cluster txterv8 --region us-east-2 --name txterv8-nodegroup --node-type t3.micro --nodes 1 --nodes-min 1 --nodes-max 1
