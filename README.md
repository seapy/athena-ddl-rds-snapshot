# Athena DDL from RDS Snapshot

Create Athena table DDL from Amazon RDS Snapshot Export to S3 information json.

- [Exporting DB Snapshot Data to Amazon S3](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_ExportSnapshot.html)


## Install

Download binary from github release.

## Usage

```shell script
$ adrs -i export_tables_info.json \
    -a hoian -s s3://my-export-bucket/prefix/your-database-name/ \
    -t users,articles
```

- `-a`, `--athenaDB`
  - Required. Your athena database name.
- `-s`, `--s3Prefix`
  - Required. Exported snapshot s3 location. include database name. not include table name.
- `-i`, `--infoPath`
  - Required. Table exported information json file location. you can download from s3 exported result.
- `-t`, `--tables`
  - Optional. Table name if you want DDL only some tables. Seprated by comma(,)

DDL result is present standard output 
