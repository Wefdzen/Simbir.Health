#main task

1. Account URL: http://localhost:8080/ui-swagger/index.html
2. Hospital URL: http://localhost:8081/ui-swagger/index.html
3. Timetable URL: http://localhost:8082/ui-swagger/index.html
4. Document URL: http://localhost:8083/ui-swagger/index.html

#additional information

1. у меня роли с маленькой буквы
2. accessToken and refreshToken хранятся в cookie. Live: accessToken 15min, refreshToken 30days.
3. в swagger в поле Authorize не надо добалять свой accessToken он просто для показателя замочков в endopoints. Все в cookie.
4. по заданию нет endpoint который бы возвращал appointments, поэтому непонятно как пользователь будет удалить по айди свою запись, => пока pgAdmin.
5. не вырубать docker контейнерe, иначе слетит env test и при следующем запуске в account-microservice будет снова созданы default users.
