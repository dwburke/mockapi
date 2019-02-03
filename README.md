# mockapi


## Sample Config

```
routes:
  "/test2":
    method: "GET"
    result: "result2"
  "/get/{id}":
    method: "GET"
    result: "param id:{{.Params.id}}"
```

