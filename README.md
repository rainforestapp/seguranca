# Seguranca

(it's portugese for security)

This tool allows updating of a security group automatically from Rainforest QA's list of VM hosts. You can then authorize just that security group for ingress.


## Using our security group

In your security group, allow "All TCP" and "All UDP" for "sg-b0574ba7" if you're in US-EAST-1. If you're not, please contact us.

## Using yourself

Setup your ``~/.aws/config`` file. The easiest way is to install the AWS cli tools and run ``aws configure``.

Then run:

```bash
RF_SECURITY_GROUP_REGION="us-east-1" RF_SECURITY_GROUP="sg-b0574ba7" go run main.go
```