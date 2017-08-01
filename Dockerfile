FROM iron/base
WORKDIR /app
# copy binary into image
COPY tic-tac-toe /app/
ENTRYPOINT ["./tic-tac-toe"]

EXPOSE 8080