The best chat for school

1. run command
 `migrate -path ./migrations -database 'sqlite3://./database/db.sqlite' up` 

2. How to come migrations back 
 `migrate -path ./database/sqlite3/migrations -database 'sqlite3://./database/sqlite3/data/db.sqlite?' down` 

3. How to fix database 
`migrate -path .//migrations -database 'sqlite3://./database/db.sqlite?' force 1 `

4. add migr files
`migrate create -ext sql -dir migrations -seq init `
