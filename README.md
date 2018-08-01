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
    git clone https://github.com/tantona/delete-s3-bucket.git
    cd delete-s3-bucket
    go install
    ```

1. Usage

    ```sh
    usage: delete-bucket --name=NAME [<flags>]

    Flags:
        --help       Show context-sensitive help (also try --help-long and --help-man).
    -n, --name=NAME  name of s3 bucket
    -v, --verbose    enable verbose logging
    ```

1. Example

    ```sh
    delete-bucket --name foobar
    ```
