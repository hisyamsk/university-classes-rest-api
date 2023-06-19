# **University Classes REST API**

Base URL: `127.0.0.1:8000/api`

Endpoints:

- `/students`
- `/students/:id`
- `/students/:id/classes`
- `/classes`
- `/classes/:id/students`
- `/enrolled-classes`

Set the value of `X-API-Key` header to `SECRET` when accessing the API

### **How to run**

- Using docker-compose

```
docker compose up
```

- Run manually (requires docker)

  1. Clone this repository

     ```
     git clone https://github.com/hisyamsk/university-classes-rest-api.git
     ```

  2. Create custom docker volume and Run PostgreSQL image

     ```
     docker volume create YOUR_VOLUME
     ```

     ```
       docker run -d \
         --name db \
         -v YOUR_VOLUME:/var/lib/postgresql/data \
         -v "$(pwd)/init.sql:/docker-entrypoint-initdb.d/init.sql" \
         -e POSTGRES_PASSWORD=foobarbaz \
         -p 5432:5432 \
         --restart unless-stopped \
         postgres:15.1-alpine
     ```

  3. Run the app image

     ```
       docker run -d \
         --name univ-api \
         -e DB_USERNAME=postgres \
         -e DB_PASSWORD=foobarbaz \
         -e DB_HOST=db \
         -e DB_PORT=5432 \
         -e APP_ADDRESS=:8000 \
         -e API_KEY_SECRET=SECRET \
         --link=db \
         -p 8000:8000 \
         --restart unless-stopped \
         hisyamsk/university-classes-api

     ```
