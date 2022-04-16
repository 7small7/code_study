## Redis哨兵

```php
'redis' => [
    'client' => env('REDIS_CLIENT', 'predis'),
    'options' => [
        'cluster' => env('REDIS_CLUSTER', 'predis'),
        'prefix'  => env('REDIS_PREFIX', Str::slug(env('APP_NAME', 'laravel'), '_') . '_database_'),
    ],
    'default' => [
        //这3个都是sentinel节点的地址
        'tcp://192.168.2.102:8005',
        'tcp://192.168.2.102:8006',
        'tcp://192.168.2.102:8007',    
        'options' => [
            'replication' => 'sentinel',
            'service'     => env('REDIS_SENTINEL_SERVICE', 'mymaster'),    //sentinel
            'parameters'  => [
                //redis的密码,没有时写null
                'password' => env('REDIS_PASSWORD', null),    
                'database' => 0,
            ],
        ],
    ],
    'cache' => [
        'url'      => env('REDIS_URL'),
        'host'     => env('REDIS_HOST', '127.0.0.1'),
        'password' => env('REDIS_PASSWORD', null),
        'port'     => env('REDIS_PORT', 6379),
        'database' => env('REDIS_CACHE_DB', 1),
    ],

],
```