# aws-account-alias [![Build Status](https://travis-ci.org/securityscorecard/aws-account-alias.svg?branch=master)](https://travis-ci.org/securityscorecard/aws-account-alias)

## Motivation

You need the AWS account alias for the credentials found by the SDK's default
credential provider.

## Our Motivation

The AWS authentication backend in Hashicorp's Vault requires the roles to be uniquely
named. If you have multiple AWS accounst you wish to authenticate against identical
roles, you have two options:

1. Mount the AWS authentication backend multiple times, and duplicate the roles with
changed AWS account IDs. This requires having a way to detect which account you're
running in, and changing the API endpoint you use based off that.
2. Mount the AWS authentication backend once, follow the [cross account access](https://www.vaultproject.io/docs/auth/aws.html#cross-account-access)
guide, and prefix the Vault roles with a string depending on the account. This tool
facilitates that.
