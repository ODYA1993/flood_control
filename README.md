# Что нужно сделать

- Реализовать интерфейс с методом для проверки правил флуд-контроля. Если за последние N секунд вызовов метода Check будет больше K, значит, проверка на флуд-контроль не пройдена.

- Интерфейс FloodControl располагается в файле main.go.

- Флуд-контроль может быть запущен на нескольких экземплярах приложения одновременно, поэтому нужно предусмотреть общее хранилище данных. Допустимо использовать любое на ваше усмотрение.

# Необязательно, но было бы круто:

Добавить поддержку конфигурации итоговой реализации. Параметры — на ваше усмотрение.

# Запуск:
- в консоли "make docker_up"
- в консоли "make run"
  
Redis был выбран в качестве хранилища данных для реализации флуд-контроля по нескольким причинам:

Быстрая работа: Redis является высокопроизводительной внутриисточной базой данных, которая хранит все данные в памяти. Это позволяет ей обеспечивать очень быструю работу при чтении и записи данных.
Сортированные множества: Redis предоставляет возможность использовать сортированные множества (sorted sets), которые позволяют хранить данные в упорядоченном виде. Это позволяет реализовать эффективный алгоритм флуд-контроля, основанный на подсчете количества запросов в заданный промежуток времени.
Простота использования: Redis имеет простой и понятный API, который легко использовать в приложениях. Он также имеет хорошую поддержку в различных языках программирования, включая Go.
Масштабируемость: Redis может масштабироваться горизонтально, что позволяет распределять нагрузку между несколькими серверами. Это позволяет обеспечить высокую производительность и надежность приложения при большом количестве запросов.
