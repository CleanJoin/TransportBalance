Rest api server - управление балансом пользователей:

<h3>Swagger:</h3>
http://localhost:8000/swagger/index.html

<h4>База данных postgres (доступ через .env)</h4>
POSTGRES_DB=avito <br>
POSTGRES_USER=avito  <br>
POSTGRES_PASSWORD=avito123 <br>
POSTGRES_HOST=db <br>
SERVER_PORT=8080 <br>
DB_HOST=db <br>
DB_PORT=5432 <br>

<h3>Два пользователя созданы:</h4>

Andey пароль Andey <br>
INSERT INTO avito.users <br>
(username, "password", "money") <br>
VALUES('Andrey', '8e756c9f2b15da6a63f84852fc39667617523133', 0.0);

INSERT INTO avito.users<br>
(username, "password", "money")<br>
VALUES('Anton', '8e756c9f2b15da6a63f84852fc39667617523134', 0.0);<br>

<h3>Show the status of server (GET heathHandler)</h3>
 serverGin.router.GET("/api/health", heathHandler())

<h3>Регистрация пользователя (password и username)</h3>
 serverGin.router.POST("/api/user", userHandler(serverGin.userStorage))

<h3>Получения баланса пользователя (Принимает id пользователя. Баланс всегда в рублях.) + (POST getMoneyUserHandler) </h3>
 serverGin.router.POST("/api/money", getMoneyUserHadler(serverGin.userStorage))

<h3>Зачисление средств,  (Принимает id пользователя и сколько средств зачислить.)+ (POST addMoneyHandler)</h3>
 serverGin.router.POST("/api/add", addMoneyHandler(serverGin.transactionsStorage))

<h3>Списание средств, (Принимает id пользователя и сколько средств списать.)+  (POST reduceUserHandler)</h3>
 serverGin.router.POST("/api/reduce", reduceMoneyHandler(serverGin.transactionsStorage))

<h3>Перевод средств от пользователя к пользователю(Принимает id пользователя с которого нужно списать средства, id пользователя которому должны зачислить средства, а также сумму.)+ (POST transferMoneyHandler)</h3>
 serverGin.router.POST("/api/transfer", transferMoneyHandler(serverGin.transactionsStorage))
	
   
<h3>Зачисление денег от пользователя в резервный счет ( userid,serviceid,orderid,money)</h3>
 serverGin.router.POST("/api/reserve", addMoneyToReserveHandler(serverGin.transactionsStorage))

<h3>Списание денег из резерва и записывание в бд данной транзации ( userid,serviceid,orderid,money)</h3>
 serverGin.router.POST("/api/reduceReserve", reduceReserveHandler(serverGin.transactionsStorage))


<h4>Запуск приложения с помощью: docker-compose up ---build <br>
 Проверить запросы можно через swagger + имеются Post curl(curl.txt) <br></h4>


<h3>Доп. задание 1:</h3>
 Создание  отчета для Бухгалтерии в формате csv (На вход: год-месяц) На выходе ссылка на CSV файл. <br>
 serverGin.router.POST("/api/docs", createGenDocHandler(serverGin.transactionsStorage))

 По ссылке которая выдается в json (createGenDocHandler) Можно сделать get и сохранить файл <br>
 serverGin.router.GET("/api/download/:filename", func(ctx *gin.Context)  <br>
	fileName := ctx.Param("filename") <br>
	ctx.FileAttachment("/usr/src/server/download/"+fileName, fileName) <br>


<h3>Доп. задание 2:(параметр старницы и сортировку по сумме и дате) (Сделано для перевода денег между пользователями)</h3>
serverGin.router.POST("/api/getmovemoney", getLastTransactionHadler(serverGin.transactionsStorage)) <br>
curl -d '{"userid": 1}' -H "Content-Type: application/json" -X POST http://localhost:8000/api/getmovemoney?page=1&filtermoney=asc&filtertime=desc

<h3>Доп. задание (Валюта) : (API)</h3>
Необходима возможность вывода баланса пользователя в отличной от рубля валюте. <br>
curl -d '{"userid": 1}' -H "Content-Type: application/json" -X POST http://localhost:8000/api/money?currency=USD <br>
Курсу валют беру отсюда https://freecurrencyapi.net/.


