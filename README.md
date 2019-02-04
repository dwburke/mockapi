# mockapi


## Sample Config

```
routes:
  "/test2":
    method: "GET"
    result: "result2"
    result_type: "text/html"
  "/get/{id}":
    method: "GET"
    result: "param id:{{.Params.id}}"
    result_type: "application/json"
  "/post/{id}":
    method: "POST"
    result: "{\"id\":{{.Params.id}},\"x\":{{.Params.x}}}"
    result_type: "application/json"
```

