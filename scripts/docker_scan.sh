docker rmi rafaelbrunotech/general_server_go:latest
docker build --target prod -t rafaelbrunotech/general_server_go .
# docker run -v /var/run/docker.sock:/var/run/docker.sock aquasec/trivy image rafaelbrunotech/general_server_go
docker run -v /var/run/docker.sock:/var/run/docker.sock anchore/grype rafaelbrunotech/general_server_go