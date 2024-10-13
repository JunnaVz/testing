```bash
sudo apt-add-repository ppa:qameta/allure
sudo apt-get update
sudo apt-get install allure
```

```bash
pikasoft@pikasoft:~/GolandProjects/testing$ gotestsum --junitfile test_results.xml ./tests/unit_tests/unit_services/mock
pikasoft@pikasoft:~/GolandProjects/testing$ mv test_results.xml allure-results/

pikasoft@pikasoft:~/GolandProjects/testing$ allure generate allure-results --clean -o allure-report
pikasoft@pikasoft:~/GolandProjects/testing$ allure serve allure-results
```


[//]: # (run tests)
```bash
pikasoft@pikasoft:~/GolandProjects/testing/tests/unit_tests/unit_repository/mock$ go test
pikasoft@pikasoft:~/GolandProjects/testing/tests/unit_tests/unit_repository/mock$ go test -shuffle=on
pikasoft@pikasoft:~/GolandProjects/testing/tests/unit_tests/unit_repository/mock$ go test -v
```


[//]: # (open Wireshark and find Loopback-lo to start capturing traffic)
[//]: # (then do commands with curl/postman to imitate e2e test: get, post, put, delete)
```bash
curl -X GET http://localhost:8080/worker-auth/signin -d '{"InputEmail":"default@admin.com","InputPassword":"admin123"}' -H "Content-Type: application/x-www-form-urlencoded"
```


[//]: # (install docker)
```bash
sudo apt-get update
sudo apt install docker-compose
sudo snap install docker
sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
sudo apt-get install docker-ce docker-ce-cli containerd.io
Environment="HTTP_PROXY=http://proxy.example.com:80/"
Environment="HTTPS_PROXY=https://proxy.example.com:443/"
docker compose up
docker-compose disable-v2
docker compose up
sudo lsof -i :5432 
sudo systemctl stop postgresql
```