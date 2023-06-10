# Nginx 
  - Nginx is a websever
    - Serves the web content
  - It is also a reverse proxy (interacts with the backends)
    - Load Balancing
      - It refers to efficiently distributing incoming network traffic accross the group of backend servers
      - It can also be called as a server pool
      - Acts as a traffic cop sitting infront of servers and routing client requests across all servers capable of fulfilling those request in a manner that maximize speed and capacity utilization and that ensures no server is overworked, which could reduce performance
      - Distributes client requests or network loads efficiently accross multiple servers
      - Ensures high availability and reliability by sending requests only to servers that are online
      - Provide the flexibility to add or subtract servers as demand dictates
    - Backend Routing (which backend to send the request to)
    - Caching (Cache the comman requests)
    - API Gateway

## NGINX layer 4 vs layer 7 proxying
  - Layer 4 ( TCP/IP stack) and 7 (Application Layer) refers to OSI model
    - We have access to in TCP/IP stack
      - Source IP, Source Port
      - Destination IP, Destination Port
      - Simple packet Inseption (SYN/TLS hello)
    - We have in layer 7 , http/gRPC etc
      - We have access to the content
      - i know where the client is going, which page they are visiting
      - Requre decryption
  - Nginx can operate in Layer 7 (eg. http) or layer 4 (tcp)
  - Layer 4 proxying is usefull when NGINX does not understand the protocol (MySql database protocol)
    - Nginx does not understand, parse or communicate with these protocol backends
  - Layer 7 proxying is usefull when nginx want to share backend connection, cache results, rewrite headers and add more to the content
  - Using stream context it becomes a layer 4 proxy
  - Using http context it becomes a layer 7 proxy

## TLS, terminations and passthrough
  - TLS -> Transport Layer Security
  - It is a way to establish end to end encryption between one another
  - Symmetric encryption is used for communication is used from communication
  - Asymmetric encryption is used intiailly to exchange symmetric key (diffie hellman)
  - Server need to authenticate themselves by supplying a certificate signed by certificate authority
  - Ngnix has TLS i.e. HTTPS and backend does not then nginx terminates the tls and decrypts and send unencrypted to backend
  - Nginx has TLS and Backend also has TLS then nginx terminates tls, decrypt it, optionally rewrite and then reencrypt the content to the backend
  - Nginx can look at the layer 7 data, rewrite the headers, cache but need to share the backend certificates or at least has its own
  - Nginx has no TLS and backend has tls then nginx proxies/streams is forwarded all the way to the backend just like a tunnel. In this type of condition, there is no caching and layer 4 check but more secure. Also nginx will not require backend certificates

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
      - get the network info : docker inspect <container name or id>

4. Nginx Connection Process Architecture
    - By default, nginx has 1 master process and 4 worker processes
    - OS kernel is responsible to send the connection to the worker processes
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