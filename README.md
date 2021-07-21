# go-rest-test
http 测试服务；

```go
[GIN-debug] GET    /healthz                  --> go-rest-test/web.(*HttpClient).buildHandler.func1 (4 handlers)
[GIN-debug] GET    /ping                     --> go-rest-test/web.(*HttpClient).buildHandler.func2 (4 handlers)
[GIN-debug] GET    /header                   --> go-rest-test/web.(*HttpClient).buildHandler.func3 (4 handlers)
[GIN-debug] GET    /data                     --> go-rest-test/web.(*HttpClient).buildHandler.func4 (4 handlers)
[GIN-debug] POST   /data                     --> go-rest-test/web.(*HttpClient).buildHandler.func4 (4 handlers)
[GIN-debug] PUT    /data                     --> go-rest-test/web.(*HttpClient).buildHandler.func4 (4 handlers)
[GIN-debug] PATCH  /data                     --> go-rest-test/web.(*HttpClient).buildHandler.func4 (4 handlers)
[GIN-debug] HEAD   /data                     --> go-rest-test/web.(*HttpClient).buildHandler.func4 (4 handlers)
[GIN-debug] OPTIONS /data                     --> go-rest-test/web.(*HttpClient).buildHandler.func4 (4 handlers)
[GIN-debug] DELETE /data                     --> go-rest-test/web.(*HttpClient).buildHandler.func4 (4 handlers)
[GIN-debug] CONNECT /data                     --> go-rest-test/web.(*HttpClient).buildHandler.func4 (4 handlers)
[GIN-debug] TRACE  /data                     --> go-rest-test/web.(*HttpClient).buildHandler.func4 (4 handlers)
[GIN-debug] GET    /ip                       --> go-rest-test/web.(*HttpClient).buildHandler.func5 (4 handlers)
[GIN-debug] GET    /panic                    --> go-rest-test/web.(*HttpClient).buildHandler.func6 (4 handlers)
[GIN-debug] GET    /1s                       --> go-rest-test/web.(*HttpClient).buildHandler.func7 (4 handlers)
[GIN-debug] GET    /err/404                  --> go-rest-test/web.(*HttpClient).buildHandler.func8 (4 handlers)
[GIN-debug] GET    /err/403                  --> go-rest-test/web.(*HttpClient).buildHandler.func9 (4 handlers)
[GIN-debug] GET    /err/500                  --> go-rest-test/web.(*HttpClient).buildHandler.func10 (4 handlers)
[GIN-debug] GET    /err/502                  --> go-rest-test/web.(*HttpClient).buildHandler.func11 (4 handlers)
[GIN-debug] POST   /post                     --> go-rest-test/web.(*HttpClient).buildHandler.func12 (4 handlers)
[GIN-debug] POST   /form                     --> go-rest-test/web.(*HttpClient).buildHandler.func13 (4 handlers)
[GIN-debug] POST   /query                    --> go-rest-test/web.(*HttpClient).buildHandler.func14 (4 handlers)

```

[GIN-debug] Listening and serving HTTP on :80






