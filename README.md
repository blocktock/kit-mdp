# go kit演示项目结合kit-cli快速生成项目服务

### Create a new service
    kit new service --help
    kit new service hello
    kit n s hello # using aliases
    
    or
    
    kit new service hello --module github.com/{group name}/hello
    kit n s hello -m github.com/{group name}/hello # using aliases
    
### Generate the service
    kit g s hello
    kit g s hello --dmw # to create the default middleware
    kit g s hello -t grpc # specify the transport (default is http)
    kit g s hello -w --gorilla # 本demo使用这个方式生成 -w 生成一些默认的服务中间件 --gorilla使用gorilla/mux来替换默认的http handler
    
### Generate new middleware
    - kit g m hi -s hello
    - kit g m hi -s hello -e # if you want to add endpoint middleware
    The only thing left to do is add your middleware logic and wire the middleware with your service/endpoint.

### Enable docker integration
    - kit g d
    This will add the individual service docker files and one docker-compose.yml file that will allow you to start your services. To start your services just run
    
    - docker-compose up
    After you run docker-compose up your services will start up and any change you make to your code will automatically rebuild and restart your service (only the service that is changed)