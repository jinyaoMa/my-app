### todo

- [x] entity.Option.Value encryption
- [x] entity snowflake+aes change
- [x] jwt
- [x] file, path+name length limit?

### 文件

上传+本地扫描 储存 md5:sha512:size

发送 md5:sha512:size =================> md5+sha256+crc32+size
`128bit/4+256bit/4+32bit/4+64bit/4=32+64+8+16=123bytes`
-> 如果文件系统仓库没有文件
创建临时占位文件 md5:sha512:size.tmp
-> 如果有文件，直接创建文件记录

（临时占位文件和记录由http服务操作）

上传文件
-> 如果存在文件记录，abort
-> 如果不存在记录，上传完毕后创建记录

断点续传
-> 多用户同时上传相同的文件，服务端使用websocket推上传文件range，客户端按照range需求上传数据

（使用websocket实现上传下载）

Websocket只推送，拦截接收
Http上传限制线程，使用abort不等待

下载文件
-> 通过提交range头部下载数据
