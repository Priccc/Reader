# Reader

>准备

1. go

2. node

3. nginx (不是必须，看需要)

>使用步骤

#### 服务端

1. 安装go

2. 新建go项目目录

    ----golang

    --------bin

    --------pkg

    --------src

3. 配置go环境

4. 在src目录下
`git clone git@github.com:Priccc/Reader.git`

5. 进入项目目录 `go run main.go`

6. 浏览器输入 `localhost:7778/index`

#### web端

1. `cd web`

2. `npm install`

3. `npm run dev` 进行开发

4. 浏览器输入 `localhost:7779`

5. `npm run bdc` 将项目打包并拷贝到服务端

#### nginx配置

```server {
  listen       7777;
  server_name  localhost;
  location / {
    proxy_pass   http://127.0.0.1:7779;
  }
  location /api {
    proxy_pass   http://127.0.0.1:7778;
  }
}
