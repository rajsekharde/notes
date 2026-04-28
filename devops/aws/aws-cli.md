# AWS CLI

Check version
```bash
aws --version
```

Access key csv stored in Downloads

Configure CLI
```bash
aws configure

Enter access key id, access key, region, output format
```

Check EC2 instances
```bash
aws ec2 describe-instances
# Displays data about all instances in json format
# Displayed in less by default. Navigate using arrow keys or space/b. Exit by typing q.

aws ec2 describe-instances --output table
# Output displayed in table format

aws ec2 describe-instances --output table > output.txt
# Store output to a file
```
