# S3对象存储网络文件浏览服务器
用goland编写，主要功能是可以连接S3对象存储服务器，并且监听端口将配置的存储服务器中桶的文件通过网页展示出来，并且S3服务器连接只在与本服务器，通过路由断绝，用户端看不到S3服务器后端链接。

效果就像用nginx搭建的本地文件服务器一样。
一个桶一个服务器runtime一个端口，互不干扰::A