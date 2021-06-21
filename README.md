Розобратся с ошибкой:
- Почему git clone пишет в stderror


Запуск для отладки
GO111MODULE=off go run main.go create -n kkm_prototype -p /home/alk/alente/kkm_1/

Сборка
Не работает разобратся почему
GO111MODULE=off go build

Сдедующий функционал:
- Формирование docker-compose для develop и master
- celery добовлять по ключу
- сделать нормальную документацию


Пример команды для разворачивания проекта 
GO111MODULE=off go run main.go create -n kkm_prototype -p /home/alk/alente/kkm_1/ domain t.ru
GO111MODULE=off go run main.go create -n kkm_prototype -p /home/alk/alente/kkm_1/ --ng true