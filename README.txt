Запуск:

#<ссылка> -i <период проверки в секундах>

# Ввод через pipe
cat links.txt | go run .

# Ручной ввод
go run .
http://ru.drivemusic.me/ -i 10
https://rus.hitmotop.com/ -i 15
https://pinkamuz.pro/ -i 10
https://mp3bob.ru/ -i 15

# Вывод в STDOUT
