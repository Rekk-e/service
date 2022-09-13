## smartway_service
# How to start

Rename ".env_example"

Enter the command in root
```yaml
docker-compose up --build
```
# Queries

Add worker
```yaml
/add_worker
```
Fields
```yaml
{
Id int
Name string
Surname string
Phone string
CompanyId int
Passport {
Type string
Number string
}
Department {
Name string
Phone string
}
}

```





