FROM alpine
ADD login-service /login-service
ENTRYPOINT [ "/login-service" ]
