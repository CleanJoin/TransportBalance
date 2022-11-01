Rest api server - управление балансом пользователей:

Swagger:
http://localhost:8000/swagger/index.html

База данных postgres (доступ через .env)
Два пользователя созданы:
Andey пароль Andey
INSERT INTO avito.users
(username, "password", "money")
VALUES('Andrey', '8e756c9f2b15da6a63f84852fc39667617523133', 0.0);
INSERT INTO avito.users
(username, "password", "money")
VALUES('Anton', '8e756c9f2b15da6a63f84852fc39667617523134', 0.0);

Show the status of server (GET heathHandler)
serverGin.router.GET("/api/health", heathHandler())

Регистрация пользователя (password и username)	
serverGin.router.POST("/api/user", userHandler(serverGin.userStorage))

Получения баланса пользователя (Принимает id пользователя. Баланс всегда в рублях.) + (POST getMoneyUserHandler) 
serverGin.router.POST("/api/money", getMoneyUserHadler(serverGin.userStorage))

Зачисление средств,  (Принимает id пользователя и сколько средств зачислить.)+ (POST addMoneyHandler)
serverGin.router.POST("/api/add", addMoneyHandler(serverGin.transactionsStorage))

Списание средств, (Принимает id пользователя и сколько средств списать.)+  (POST reduceUserHandler)
serverGin.router.POST("/api/reduce", reduceMoneyHandler(serverGin.transactionsStorage))

Перевод средств от пользователя к пользователю(Принимает id пользователя с которого нужно списать средства, id пользователя которому должны зачислить средства, а также сумму.)+ (POST transferMoneyHandler)
serverGin.router.POST("/api/transfer", transferMoneyHandler(serverGin.transactionsStorage))
	
   
Зачисление денег от пользователя в резервный счет ( userid,serviceid,orderid,money)
serverGin.router.POST("/api/reserve", addMoneyToReserveHandler(serverGin.transactionsStorage))

Списание денег из резерва и записывание в бд данной транзации ( userid,serviceid,orderid,money)
serverGin.router.POST("/api/reduceReserve", reduceReserveHandler(serverGin.transactionsStorage))


Запуск приложения с помощью: docker-compose up ---build
Проверить запросы можно через swagger + имеются Post curl(curl.txt)


Доп. задание 1:
Создание  отчета для Бухгалтерии в формате csv (На вход: год-месяц) На выходе ссылка на CSV файл.
serverGin.router.POST("/gendoc", createGenDocHandler())


Доп. задание 2:(параметр старницы и сортировку по сумме и дате) (Сделано для перевода денег между пользователями)
serverGin.router.POST("/api/getmovemoney", getLastTransactionHadler(serverGin.transactionsStorage))
curl -d '{"userid": 1}' -H "Content-Type: application/json" -X POST http://localhost:8000/api/getmovemoney?page=1&filtermoney=asc&filtertime=desc

Доп. задание (Валюта) : (API)
Необходима возможность вывода баланса пользователя в отличной от рубля валюте.
curl -d '{"userid": 1}' -H "Content-Type: application/json" -X POST http://localhost:8000/api/money?currency=USD
Курсу валют беру отсюда https://freecurrencyapi.net/.


