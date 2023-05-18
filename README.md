# VPC Flow Log Ananlyzer

This is an aws lambda function that ingests flow logs from CloudWatch and
alerts a concerned email with any ip addresses that are trying to acccess 
ports that are block in a security group for a set of interfaces.

