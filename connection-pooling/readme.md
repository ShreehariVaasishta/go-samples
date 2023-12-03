# Connection Pooling Sample Implementation

Implement a simple connection pooling from scratch.

## Usage

### Alter maximum connection locally 
1. mysql

2. Current max connections

    ```bash
    mysql> show variables like 'max_connections';
    +-----------------+-------+
    | Variable_name   | Value |
    +-----------------+-------+
    | max_connections | 151   |
    +-----------------+-------+
    1 row in set (0.02 sec)
    ```

3. Lower it for testing purpose

    ```bash
    mysql> SET GLOBAL max_connections = 5;
    ```

    ```bash
    mysql> show variables like 'max_connections';
    +-----------------+-------+
    | Variable_name   | Value |
    +-----------------+-------+
    | max_connections | 5     |
    +-----------------+-------+
    1 row in set (0.02 sec)
    ```

### Connection Pooling

#### Without Connection Pooling
1. Comment the `withConnPool()` in the `main` function and run the program.

    ```bash
    go run main.go
    ```

2. Run the test script

    ```bash
    sh test.sh
    ```

3. You will see the errors in the consolte:

    ```bash
    connection-pooling/connection-pooling on  main [?] via 🐹 v1.20.3 on ☁  (eu-west-1) on ☁  shreehari@aiplanet.com(europe-west4) 
    ❯ go run main.go

    !!!Running without connection pooling.
    [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

    [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
    - using env:   export GIN_MODE=release
    - using code:  gin.SetMode(gin.ReleaseMode)

    [GIN-debug] GET    /query                    --> main.withoutConnPool.func1 (3 handlers)
    [GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
    Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
    [GIN-debug] Listening and serving HTTP on :9082
    [GIN] 2023/12/03 - 18:51:13 | 200 |          1m0s |             ::1 | GET      "/query"
    [GIN] 2023/12/03 - 18:51:13 | 200 |          1m0s |             ::1 | GET      "/query"
    [GIN] 2023/12/03 - 18:51:13 | 200 |          1m0s |             ::1 | GET      "/query"
    Error 1040: Too many connections
    Error 1040: Too many connections
    [GIN] 2023/12/03 - 18:51:14 | 500 |          1m1s |             ::1 | GET      "/query"
    [GIN] 2023/12/03 - 18:51:14 | 500 |          1m1s |             ::1 | GET      "/query"
    Error 1040: Too many connections
    Error 1040: Too many connections
    ```

#### With Connection Pooling

1. Comment the `withoutConnPool()` and uncomment `withConnPool()`  and run the program.

2. Run the test script

    ```bash
    sh test.sh
    ```

3. Output

    ```bash
    ❯ go run main.go

    !!!Running with connection pooling.
    [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

    [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
    - using env:   export GIN_MODE=release
    - using code:  gin.SetMode(gin.ReleaseMode)

    [GIN-debug] GET    /query                    --> main.withConnPool.func1 (3 handlers)
    [GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
    Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
    [GIN-debug] Listening and serving HTTP on :9082
    [GIN] 2023/12/03 - 18:46:12 | 200 |          1m0s |             ::1 | GET      "/query"
    [GIN] 2023/12/03 - 18:46:12 | 200 |          1m0s |             ::1 | GET      "/query"
    [GIN] 2023/12/03 - 18:46:12 | 200 |          1m0s |             ::1 | GET      "/query"
    [GIN] 2023/12/03 - 18:46:12 | 200 |          1m0s |             ::1 | GET      "/query"
    [GIN] 2023/12/03 - 18:46:12 | 200 |          1m0s |             ::1 | GET      "/query"
    [GIN] 2023/12/03 - 18:47:12 | 200 |          2m0s |             ::1 | GET      "/query"
    [GIN] 2023/12/03 - 18:47:12 | 200 |          2m0s |             ::1 | GET      "/query"
    [GIN] 2023/12/03 - 18:47:12 | 200 |          2m0s |             ::1 | GET      "/query"
    [GIN] 2023/12/03 - 18:47:12 | 200 |          2m0s |             ::1 | GET      "/query"
    [GIN] 2023/12/03 - 18:47:12 | 200 |          2m0s |             ::1 | GET      "/query"

    ```