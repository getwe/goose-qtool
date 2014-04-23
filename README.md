#goose-qtool

查询goose的命令行工具.

把输入命令行转化为二进制buffer发给server,等待server返回结果后直接打印退出.使用方法:
    
    goose-qtool -i 127.0.0.1 -p 8808 -c '{"query":"test"}'
    
主要在于配合goose的demo使用,比如

[goose-demo](https://github.com/getwe/goose-demo)
