 测试根据Open API规范文件生成文档和代码
==

### 依赖：     
maven   
java13

###使用说明：
可以通过com.generator_test.test.GeneratorTest类的main方法来进行测试。     
通过修改configurator.setGeneratorName的参数来决定输出的格式，具体枚举值见"附录1"     
通过修改configurator.setInputSpec参数来修改输入的yaml文件目录

####附录1：Generator枚举：

CLIENT generators:
- ada
- android
- apex
- bash
- c
- clojure
- cpp-qt-client
- cpp-restsdk
- cpp-tiny (beta)
- cpp-tizen
- cpp-ue4 (beta)
- crystal (beta)
- csharp
- csharp-netcore
- dart
- dart-dio
- dart-dio-next (experimental)
- dart-jaguar
- eiffel
- elixir
- elm
- erlang-client
- erlang-proper
- go
- groovy
- haskell-http-client
- java
- java-micronaut-client (beta)
- javascript
- javascript-apollo (beta)
- javascript-closure-angular
- javascript-flowtyped
- jaxrs-cxf-client
- jmeter
- k6 (beta)
- kotlin
- lua (beta)
- nim (beta)
- objc
- ocaml
- perl
- php
- php-dt (beta)
- powershell (beta)
- python (experimental)
- python-legacy
- r
- ruby
- rust
- scala-akka
- scala-gatling
- scala-sttp (beta)
- scalaz
- swift5
- typescript (experimental)
- typescript-angular
- typescript-aurelia
- typescript-axios
- typescript-fetch
- typescript-inversify
- typescript-jquery
- typescript-nestjs (experimental)
- typescript-node
- typescript-redux-query
- typescript-rxjs


SERVER generators:
- ada-server
- aspnetcore
- cpp-pistache-server
- cpp-qt-qhttpengine-server
- cpp-restbed-server
- csharp-nancyfx
- erlang-server
- fsharp-functions (beta)
- fsharp-giraffe-server (beta)
- go-echo-server (beta)
- go-gin-server
- go-server
- graphql-nodejs-express-server
- haskell
- java-inflector
- java-msf4j
- java-pkmst
- java-play-framework
- java-undertow-server
- java-vertx-web (beta)
- jaxrs-cxf
- jaxrs-cxf-cdi
- jaxrs-cxf-extended
- jaxrs-jersey
- jaxrs-resteasy
- jaxrs-resteasy-eap
- jaxrs-spec
- kotlin-server
- kotlin-spring
- kotlin-vertx (beta)
- nodejs-express-server (beta)
- php-laravel
- php-lumen
- php-mezzio-ph
- php-slim4
- php-symfony
- python-aiohttp
- python-blueplanet
- python-fastapi (beta)
- python-flask
- ruby-on-rails
- ruby-sinatra
- rust-server
- scala-akka-http-server (beta)
- scala-finch
- scala-lagom-server
- scala-play-server
- scalatra
- spring


DOCUMENTATION generators:
- asciidoc
- cwiki
- dynamic-html
- html
- html2
- markdown (beta)
- openapi
- openapi-yaml
- plantuml (beta)


SCHEMA generators:
- avro-schema (beta)
- graphql-schema
- ktorm-schema (beta)
- mysql-schema
- protobuf-schema (beta)
- wsdl-schema (beta)


CONFIG generators:
- apache2
