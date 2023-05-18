#!/bin/bash
aws lambda create-function --function-name logAnalyzer --runtime go1.x --role arn:aws:iam::number:role/lambda-ex --handler main --zip-file fileb://main.zip
