# Smartway_service
## How to start

Rename ".env_example" to ".env"

Build and run service
```yaml
docker-compose up --build
```
# Queries

## Add worker
```yaml
/add_worker
```
Body
```yaml
{
  id int,
  Name string,
  Surname string,
  Phone string,
  CompanyId int,
  Passport {
    id int
    Type string
    Number string
  },
  Department {
    id int
    Name string
    Phone string
  }
}
```

## Delete worker
```yaml
/delete_worker/{id}
```

## Get worker
```yaml
/get_worker/{id}
```

## Get workers by company id
```yaml
/get_workers_by_company_id/{company_id}
```

## Get workers by department
```yaml
/get_workers_by_department/{name}
```

## Update worker
```yaml
/change_worker/{id}
```

Body
```yaml
{
  id int,
  Name string,
  Surname string,
  Phone string,
  CompanyId int,
  Passport {
    id int
    Type string
    Number string
  },
  Department {
    id int
    Name string
    Phone string
  }
}
```




