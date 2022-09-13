## smartway_service
# How to start

Rename ".env_example"

Enter the command in root
```yaml
docker-compose up --build
```
## Queries

# Add worker

Url
```yaml
/add_worker
```
Body
```yaml
{
  Id int,
  Name string,
  Surname string,
  Phone string,
  CompanyId int,
  Passport {
    Type string
    Number string
  },
  Department {
    Name string
    Phone string
  }
}
```

# Delete worker

Url
```yaml
/delete_worker/{id}





