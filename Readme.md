# API service
[![pipeline status](https://git.digital.rt.ru/autofaq/b2c_micro/api/badges/main/pipeline.svg)](https://git.digital.rt.ru/autofaq/b2c_micro/api/-/commits/main)  
Access to b2c microservices through this service


## Config file
Create config file with services addresses. 
Services:
1. vdc - get camera information

Config example:
```yaml
servers:
  vdc: 127.0.0.1:8001
```
> If use `docker-compose` or `swarm` user service name instead of IP address 

## Usage
```
api [OPTIONS]
Options:  
-config string
    API server configuration file (default "apiserver.yml")
-host string
    Bind addres by API server (default "0.0.0.0")
-port string
    Bind port by API server (default "8000") 
```

## Docs
User `/docs` to access to Redoc page.

## Metrics
User Prometheus to collect metrics. Metrics URL is `/metrics`
**Metrics**
* *api_http_requests_total* - total amount of http requests. Type counter.
* *api_http_request_duration* - requests duration. Type Histogram
