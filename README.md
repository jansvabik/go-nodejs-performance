# Go vs. Node.js performance test

Hello! This is my project for testing Go vs. Node.js REST API performance. This includes some data manipulations which you can expect (CRUD). I tried to maintain at least similar file structure to make easier to find the same code in the second language for newcomers.

## Results

In my tests all participants were located in Czech Republic (client and server). Tests were done macOS/localhost and on Debian 10 Buster/Nginx 1.14.

### Localhost on macOS

#### GET /

| Lang    | Requests | Average time   |
|---------|----------|----------------|
| Node.js | 10000    | ~17.1979748 ms |
| Go      | 10000    | ~18.3700849 ms |


#### POST /

| Lang    | Requests | Average time   |
|---------|----------|----------------|
| Node.js | 1000     | ~18.5389170 ms |
| Go      | 1000     | ~16.2108910 ms |

## Installation

You should have already prepared server for deployment with some kind of web server like [Nginx](https://www.nginx.com/) or [Apache](https://httpd.apache.org/).
