# Go vs. Node.js performance test

Hello! This is my project for testing Go vs. Node.js REST API performance. This includes some data manipulations which you can expect (CRUD). I tried to maintain at least similar file structure to make easier to find the same code in the second language for newcomers.

## Results

In my tests all participants were located in Czech Republic (client and server). Tests were done macOS/localhost and on Debian 10 Buster/Nginx 1.14.

### Localhost on macOS

#### GET / (few records)

| # | Lang    | Requests | Average time   | Saved |
|---|---------|----------|----------------|-------|
| 1 | Node.js | 10000    | ~17.1979748 ms | 6.38% |
| 2 | Go      | 10000    | ~18.3700849 ms | 0%    |

#### GET / (20k records at once)

| # | Lang    | Requests | Average time    | Saved |
|---|---------|----------|-----------------|-------|
| 1 | Go      | 10000    | ~134.1045839 ms | 65.2% |
| 2 | Node.js | 10000    | ~385.3674946 ms | 0%    |


#### POST /

| # | Lang    | Requests | Average time   | Saved |
|---|---------|----------|----------------|-------|
| 1 | Go      | 10000    | ~16.5671425 ms | 3.65% |
| 2 | Node.js | 10000    | ~17.1951589 ms | 0%    |

#### DELETE /{url}/

| # | Lang    | Requests | Average time   | Saved |
|---|---------|----------|----------------|-------|
| 1 | Go      | 1000     | ~16.8426170 ms | 2.38% |
| 2 | Node.js | 1000     | ~17.2534320 ms | 0%    |

## Installation

You should have already prepared server for deployment with some kind of web server like [Nginx](https://www.nginx.com/) or [Apache](https://httpd.apache.org/).
