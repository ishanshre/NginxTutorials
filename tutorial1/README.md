# Custom Nginx create
    - Create two files Dockerfile and docker-compose.yml
    - Docker file consist of container infomations as well as actions and also to create images
    - docker-compose.yml is file where we customize nginx

# Building Image
    - Copy the html/index.html to /usr/share/nginx/html/ in Dockerfile
    - Use command 
      - To build
        - docker build -t webserver .
          - Then webserver image is created
      - To run
        - docker run -it --rm -d -p 8080:80 --name web webserver 