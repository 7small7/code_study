[TOC]

## MySQL主从复制

```php
'mysql' => [
    'driver'         => 'mysql',
    'read'           => [
        // 第一组slave节点配置
        [
            'host' => env('DB_HOST', '127.0.0.1'),
            'port' => env('DB_READ_PORT', '3305'),
        ],
        // 第二组slave节点配置
        [
            'host' => env('DB_HOST', '127.0.0.1'),
            'port' => env('DB_READ_PORT', '3305'),
        ],
        // 第三组slave节点配置
        [
            'host' => env('DB_HOST', '127.0.0.1'),
            'port' => env('DB_READ_PORT', '3305'),
        ],
    ],
    'write'          => [
        // 一组master节点配置
        [
            'host' => env('DB_HOST', '127.0.0.1'),
            'port' => env('DB_WRITE_PORT', '3304'),
        ]
    ],
    'url'            => env('DATABASE_URL'),
    'database'       => env('DB_DATABASE', 'forge'),
    'username'       => env('DB_USERNAME', 'forge'),
    'password'       => env('DB_PASSWORD', ''),
    'unix_socket'    => env('DB_SOCKET', ''),
    'charset'        => 'utf8mb4',
    'collation'      => 'utf8mb4_unicode_ci',
    'prefix'         => '',
    'prefix_indexes' => true,
    'strict'         => true,
    'engine'         => null,
    'options'        => extension_loaded('pdo_mysql') ? array_filter([
        PDO::MYSQL_ATTR_SSL_CA => env('MYSQL_ATTR_SSL_CA'),
    ]) : [],
],
```