# Migrating Dockerised project to new EC2 instance

## On old instance:

Dump Postgres data into a file:
```bash
cd <project-folder>
docker exec <db-container-name> pg_dump -U <user> -d <db> > notes_db.sql
```
## On local device:

Copy dump file from old instance:
```bash
chmod 400 <key-pair-file> # Change permissions for ssh key file

scp -i <key-pair-file> ubuntu@<Old instance IP>:/home/ubuntu/<project-folder>/notes_db.sql .
```

Copy dump file to new instance:
```bash
scp -i <key-pair-file> ./notes_db.sql ubuntu@<new-ec2-IP>:/home/ubuntu/
```

## On new instance:

Build & run dockerised project.
Don't run db migrations. postgres should be empty

Restore data from dump file:
```bash
docker exec -i <db-container-name> psql -U <user> -d <db> < /home/ubuntu/notes_db.sq
```
