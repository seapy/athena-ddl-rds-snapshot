我是光年实验室高级招聘经理。
我在github上访问了你的开源项目，你的代码超赞。你最近有没有在看工作机会，我们在招软件开发工程师，拉钩和BOSS等招聘网站也发布了相关岗位，有公司和职位的详细信息。
我们公司在杭州，业务主要做流量增长，是很多大型互联网公司的流量顾问。公司弹性工作制，福利齐全，发展潜力大，良好的办公环境和学习氛围。
公司官网是http://www.gnlab.com,公司地址是杭州市西湖区古墩路紫金广场B座，若你感兴趣，欢迎与我联系，
电话是0571-88839161，手机号：18668131388，微信号：echo 'bGhsaGxoMTEyNAo='|base64 -D ,静待佳音。如有打扰，还请见谅，祝生活愉快工作顺利。

# Athena DDL from RDS Snapshot

Create Athena table DDL from Amazon RDS Snapshot Export to S3 information json.

- [Exporting DB Snapshot Data to Amazon S3](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_ExportSnapshot.html)


## Install

Download binary from github release.

```shell script
# macOS
$ curl -LO https://github.com/seapy/athena-ddl-rds-snapshot/releases/download/v1.0/macos.tar
$ tar -xvf macos.tar
```

## Usage

Print DDL 

```shell script
$ adrs -i export_tables_info.json \
    -a hoian \
    -s s3://my-export-bucket/prefix/your-database-name/ \
    -t users,articles
```

Execute DDL to Athena
```shell script
$ adrs -i export_tables_info.json \
    -a hoian \
    -s s3://my-export-bucket/prefix/your-database-name/ \
    -t users,articles \
    -w primary \
    -e
```

- `-a`, `--athenaDB`
  - Required. Your athena database name.
- `-s`, `--s3Prefix`
  - Required. Exported snapshot s3 location. include database name. not include table name.
- `-i`, `--infoPath`
  - Required. Table exported information json file location. you can download from s3 exported result.
- `-t`, `--tables`
  - Optional. Table name if you want DDL only some tables. Seprated by comma(,)
- `-e`, `--execute`
  - Optional. Execute sql to Athena (Default false)
  - Required AWS Secret Environment(AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_REGION)
- `-w`, `--workgroup`
  - Optional.(Required if execute is true) Athena workgroup. 
