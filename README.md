# DOIT BASE
## reference 

https://github.com/flipped-aurora/gin-vue-admin


## tech stack

- backend: go
    - use go-fiber for http server
- frontend: nextjs (build static)
- database: postgres
    - use gorm for orm but don't auto migrate
    - use goose to migrate
    - write function, procedure, event schedule in migrate
- cache: memcache (with interface can use for redis)
- async task: 
    - cronjob (save status in db)
    - in memmory message queue (with interface can use for kafka)
- docker compose for dev env

