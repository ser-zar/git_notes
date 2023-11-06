# Шпаргалка по git

Git - система контроля версия, используется для синхронизации изменений в проекте.

## Перечень основных команд

+ git init - создание локального репозитория в текущей директории
+ git status - отображение статуса локального репозитория
+ git add 'modified files' - добавление файлов которые git будет отслеживать
  > Для добавления всех файлов опция --all или .
+ git commit -m 'Your commit'- фиксация изменений с обязательным комментированием
  > Чтобы открыть редактор для комита, убраьт опцию -m
+ git checkout *branch* - переключение на другую ветку
  > Опция -b для создание ветки, если таковая не существует
+ git remote add *origin* url - связать локальный репозиторий с удаленным
  > origin  стандартный псевдоним, с помощью которого можно обращаться к главному удалённому репозиторию (обычно такой репозиторий один). Это значительно упрощает работу.
+ git remote -v - убедиться, что репозитории связаны
+ git push *origin* *branch* - залить изменения на удаленный репозиторий на выбранную ветку
  > origin  стандартный псевдоним, с помощью которого можно обращаться к главному удалённому репозиторию (обычно такой репозиторий один). Это значительно упрощает работу.

+ git push *origin* *branch* - залить изменения в удаленный репозиторий на выбранную ветку
+ git pull *origin* *branch* - синхронизировать содержимое локального репозитория с удаленным с выбранной ветки
+ git branch -C *branch* - создание новой ветки
  > Для перемещения ветки (изменения ее имени) -M
+ git clone *URL*- клонировать репозиторий на локальную машину
+ git log - просмотр логов изменений
