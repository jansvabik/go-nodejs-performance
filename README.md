# Go vs. Node.js performance test

Hello! This is my project for testing Go vs. Node.js REST API performance. This includes some data manipulations which you can expect (CRUD). I tried to maintain at least similar file structure to make easier to find the same code in the second language for newcomers.

## Results

In my tests all participants were located in Czech Republic (client and server). Tests were done macOS/localhost and on Debian 10 Buster/Nginx 1.14.

### Localhost on macOS

#### GET /

| # | Lang    | Requests | Average time   |
|---|---------|----------|----------------|
| 1 | Node.js | 10000    | ~17.1979748 ms |
| 2 | Go      | 10000    | ~18.3700849 ms |


#### POST /

| # | Lang    | Requests | Average time   |
|---|---------|----------|----------------|
| 1 | Go      | 10000    | ~16.8654061 ms |
| 2 | Node.js | 10000    | ~18.6725748 ms |

#### DELETE /{url}/

| # | Lang    | Requests | Average time   |
|---|---------|----------|----------------|
| 1 | Go      | 1000     | ~16.8426170 ms |
| 2 | Node.js | 1000     | ~17.2534320 ms |

## Installation

You should have already prepared server for deployment with some kind of web server like [Nginx](https://www.nginx.com/) or [Apache](https://httpd.apache.org/).
