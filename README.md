## Requirement

- [Finch](https://github.com/runfinch/finch/releases) or docker

```
@ ~/src/bingo # cat ~/.aws/config
[default]
region = ap-northeast-1
output = json
@ ~/src/bingo # cat ~/.aws/credentials
[default]
aws_access_key_id = dummy
aws_secret_access_key = dummy
```

## Usage

When vm is not running  
`make vm-start`

Start containers  
`make setup`

Just once.  
`make create-tables`

`make run`