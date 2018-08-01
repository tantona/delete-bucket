# delete-bucket

```
remove_bucket failed: s3://foobar-12341234 An error occurred (BucketNotEmpty) when calling the DeleteBucket operation: The bucket you tried to delete is not empty. You must delete all versions in the bucket.
```

Arrrggghh!! 😡 😡 😡 😡

Do you hate using the console to empty and then delete versioned S3 Buckets?  Never fear! This little CLI will do all the heavy lifting for you!


## Getting started 

1. [Install and configure awscli](https://docs.aws.amazon.com/cli/latest/userguide/installing.html)

1. Install

    ```sh
    go get github.com/tantona/delete-bucket
    ```

1. Delete a bucket!

    ```sh
    delete-bucket --name foobar
    ```

## Usage

    ```sh
    usage: delete-bucket [<flags>] <bucket name>

    Flags:
        --help                Show context-sensitive help (also try --help-long and --help-man).
    -v, --verbose             enable verbose logging
    -p, --profile="default"   aws profile name
    -r, --region="us-east-1"  aws region

    Args:
    <bucket name>  name of s3 bucket
    ```

