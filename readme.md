# Read Me

It is a scaffold project to help quickly setup a golang based web application. The web framework being used is gin. Gorm is used for persistence. You can set some parameters, and then a new runnable project can be generated. The project receives Rest API call to do the CRUD operation. It also guides how to setup database.

Two goals to achieve:
  1. easy to setup
  2. relatively good practice

# How to setup

## Prerequisites:

docker and docker-compose should be ready to use.
to install docker: https://docs.docker.com/engine/install/
to install docker-compose: https://docs.docker.com/compose/install/

## Steps to setup:
  1. create folder structure based on your package name. For instance, if your package is going to be github.com/yourname/vblog, create folders github.com/yourname/vblog, and put this project into it.
  2. edit template.sh for the 'export' part.

 
    export project_name=vblog
    export package_name=github.com/yourname
    export entity_capital=Blog
    export entity=blog
    export entity_plural=blogs
    export entity_capital_plural=Blogs
 

  3. sudo chmod +x template.sh && ./template.sh
  4. go to scripts folder and run 'sudo chmod +x create_db.sh && sudo ./create_db.sh'
  5. Connect to postgres with any clients. I use DBeaver for most of time, sometimes pgAdmin. user/password: postgres/postgres, the db name can be found in docker-compose.yaml. Open db.sql and run the content, for example,
     Run script:


    CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
    CREATE TABLE <YOUR_PROJECT_NAME>
    (
          id UUID DEFAULT uuid_generate_v4(),
          created_at timestamp with time zone,
          updated_at timestamp with time zone,
          name text,
          CONSTRAINT blog_pkey PRIMARY KEY (id)
    )

  6. run 'go run main.go '
    then you will see output something like this (it is blogs because in template.sh, it uses 'blog' in template.sh):

           [GIN-debug] GET    /test                     --> github.com/vcyreg/vblog/routers.InitRouter.func1 (3 handlers)
    	   [GIN-debug] GET    /blogs                    --> github.com/vcyreg/vblog/routers.GetBlogs (3 handlers)
    	   [GIN-debug] GET    /blogs/:id                --> github.com/vcyreg/vblog/routers.GetBlog (3 handlers)
    	   [GIN-debug] POST   /blogs                    --> github.com/vcyreg/vblog/routers.AddBlog (3 handlers)
    	   [GIN-debug] PUT    /blogs/:id                --> github.com/vcyreg/vblog/routers.EditBlog (3 handlers)
    	   [GIN-debug] DELETE /blogs/:id                --> github.com/vcyreg/vblog/routers.DeleteBlog (3 handlers)

   Congratulations! You have a runnable project to start with!

   To create an entity, you can issue a post request to http://127.0.0.1:8000/blogs . The request body should be {"name": "any string"}
