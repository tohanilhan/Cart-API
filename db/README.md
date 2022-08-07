
# POSTGRESQL

You have to make sure that you are in the **db directory at the project**.

`cd db`

**init-db.sql** file is used to initialize the database.


**Create Postgresql Image**

`docker build -t postgresql-img:latest .`

**Create Postgresql Container**

`docker run -d --name postgresql-cnt  -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=kbueRm8nuEj9DBVK -e POSTGRES_DB=postgres -v {YOUR-DESIRED-PATH}:/var/lib/postgresql/data -p 5432:5432 postgresql-img:latest`

