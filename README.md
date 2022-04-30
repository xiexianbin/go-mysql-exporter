# go-mysql-exporter

golang mysql-exporter demo, Learn to develop exporter as the goal.

```
$ go run main.go -h
Usage: mysql-exporter -addr 0.0.0.0:9306 -mysqlAddr 127.0.0.1:3306 -u root -p root -d mysql
  -addr string
        server url (default "0.0.0.0:9306")
  -d string
        mysql db (default "mysql")
  -help
        show help message
  -mysqlAddr string
        mysql url (default "127.0.0.1:3306")
  -p string
        mysql password (default "root")
  -u string
        mysql user (default "root")
```

## mysql monitor sql

Only part of the function is realized

```
mysql> show global status where variable_name like "%slow%";
+---------------------+-------+
| Variable_name       | Value |
+---------------------+-------+
| Slow_launch_threads | 0     |
| Slow_queries        | 0     |
+---------------------+-------+
2 rows in set (0.02 sec)
mysql> show global status where variable_name like "%queries%";
+---------------+-------+
| Variable_name | Value |
+---------------+-------+
| Queries       | 9513  |
| Slow_queries  | 0     |
+---------------+-------+
2 rows in set (0.01 sec)
mysql> show global status where variable_name like "com_%";
+-------------------------------------+-------+
| Variable_name                       | Value |
+-------------------------------------+-------+
...
| Com_delete                          | 4     |
| Com_delete_multi                    | 0     |
| Com_insert                          | 182   |
| Com_insert_select                   | 0     |
| Com_select                          | 2937  |
| Com_update                          | 10    |
| Com_update_multi                    | 0     |
...
mysql> show global status where variable_name like "thread%";
+-------------------+-------+
| Variable_name     | Value |
+-------------------+-------+
| Threads_cached    | 5     |
| Threads_connected | 1     |
| Threads_created   | 6     |
| Threads_running   | 2     |
+-------------------+-------+
4 rows in set (0.00 sec)
mysql> show global variables like "%connect%";
+-----------------------------------------------+----------------------+
| Variable_name                                 | Value                |
+-----------------------------------------------+----------------------+
| character_set_connection                      | utf8mb4              |
| collation_connection                          | utf8mb4_0900_ai_ci   |
| connect_timeout                               | 10                   |
| connection_memory_chunk_size                  | 8912                 |
| connection_memory_limit                       | 18446744073709551615 |
| disconnect_on_expired_password                | ON                   |
| global_connection_memory_limit                | 18446744073709551615 |
| global_connection_memory_tracking             | OFF                  |
| init_connect                                  |                      |
| max_connect_errors                            | 100                  |
| max_connections                               | 151                  |
| max_user_connections                          | 0                    |
| mysqlx_connect_timeout                        | 30                   |
| mysqlx_max_connections                        | 100                  |
| performance_schema_session_connect_attrs_size | 512                  |
+-----------------------------------------------+----------------------+
15 rows in set (0.08 sec)
mysql>  show global variables where variable_name = "max_connections";
+-----------------+-------+
| Variable_name   | Value |
+-----------------+-------+
| max_connections | 151   |
+-----------------+-------+
1 row in set (0.00 sec)
mysql>  show global status where variable_name like "Bytes%";
+----------------+---------+
| Variable_name  | Value   |
+----------------+---------+
| Bytes_received | 1087895 |
| Bytes_sent     | 4168119 |
+----------------+---------+
2 rows in set (0.00 sec)
mysql>  show global status where variable_name like "open_%";
+--------------------------+-------+
| Variable_name            | Value |
+--------------------------+-------+
| Open_files               | 2     |
| Open_streams             | 0     |
| Open_table_definitions   | 74    |
| Open_tables              | 499   |
| Opened_files             | 2     |
| Opened_table_definitions | 844   |
| Opened_tables            | 982   |
+--------------------------+-------+
7 rows in set (0.00 sec)
```
