Задание - написать бенчмарк, который оценивает RPS запроса Postgresql.
Минимальные входные данные (конфиг): оцениваемый sql запрос,
dsn postgresql, время проведения измерения RPS в миллисекундах.
Нужно использовать конкурентность для максимальной нагрузки на бд.
Программа не должна содержать гонки (data race).

в папке test реализован кастомный бенчмарк и встроенный из пакета testing
конфиги передаются с помощью optios

запуск для тестирования работоспособности $go test -v
