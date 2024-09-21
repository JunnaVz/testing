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