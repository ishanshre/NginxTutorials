# Nginx Connection Processing Ar
  - No. of current process
      - Docker 
        - docker top <container-name>
      - Inside shell
        - ps -C nginx -f
  - Worker
    - Nginx deployes event based model and os based request and connections
    - Max 512 connections by default