# 3 Node js app with nginx in one docker network

## Commands
    - Create a docker file for node js
    - docker command to create image (simple): docker build . -t <image name>
    - Now running a container of that image:- docker run -p 8080:8080 -d <image name>
    - Another container with hostname: docker run -p 8001:8080 --hostname testnode -d nodeapp
    - After nginx.config
      - spin up 3 node app docker container
        - docker run --hostname nodeapp1 --name nodeapp1 -d nodeapp
        - docker run --hostname nodeapp2 --name nodeapp2 -d nodeapp
        - docker run --hostname nodeapp3 --name nodeapp3 -d nodeapp
      - spin up docker container
        - docker run --name nginx --hostname ng1 -p 8080:8080 -v NGINX.CONF_path:/etc/share/nginx nginx
        - docker run --name nginx --hostname ng1 -p 8080:8080 -v /user/app/nginx.conf:/etc/share/nginx nginx
      - Create a network using docker
        - docker network create <network name>
        - docker network create backendnodenet
      - Connect the nodeapp backend
        - docker network connect backendnodenet nodeapp1
        - docker network connect backendnodenet nodeapp2
        - docker network connect backendnodenet nodeapp3