docker rmi rafaelbrunoss/general_server_go:latest
docker build --target prod -t rafaelbrunoss/general_server_go .
# docker run -v /var/run/docker.sock:/var/run/docker.sock aquasec/trivy image rafaelbrunoss/general_server_go
docker run -v /var/run/docker.sock:/var/run/docker.sock anchore/grype rafaelbrunoss/general_server_go