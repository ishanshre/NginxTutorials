# Nginx 
1. Pull Nginx
    - docker pull nginx

2. Run the Container
    - docker run -it --rm -d -p 8000:80 --name website nginx:latest
        -it :allows us to have interactive bash shell
        --rm : cleanup the container if it exists 
        -d : backgroud process run as  a deamon
3. Nginx version
    - nginx -v
    - For docker 
      - docker exec <container_id> nginx -v

4. Nginx Connection Process Architecture

                    Master Process
                        1. Worker process
                        2. Worker Process 
                        3. worker processs
                        4. worker process

    - No. of current process
      - Docker 
        - docker top <container-name>
      - Inside shell
        - ps -C nginx -f

5. Nginx Config
   1. resides in /etc/nginx/nginx.conf
   2. No. of workers set to no. of cpu cores 