# osenv
set env to os by .env

```shell
## windows

>dir

2023/01/07  10:44    <DIR>          .
2023/01/07  10:44    <DIR>          ..
2023/01/06  18:59               185 .env


>type .env
SECRET_KEY=xxx
ACCESS_KEY=xxx
NACOS_PORT=1234
NACOS_REGION_ID=xxx
NACOS_URL=xxx

## install osenv
>go install github.com/tangzixiang/osenv

## run osenv to set env which read from file .env
>osenv
 
>set NACOS_PORT
NACOS_PORT=1234
```
