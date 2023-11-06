# Шпаргалка по git

Git - система контроля версия, используется для синхронизации изменений в проекте.

## Перечень основных команд:
+ git init - создание локального репозитория в текущей директории
+ git status - отображение статуса локального репозитория
+ git add 'modified files' - добавление файлов которые git будет отслеживать
+ git commit -m 'Your commit'- фиксация изменений с обязательным комментированием
+ git checkout *branch* - переключение на другую ветку
+ git remote add *origin* url - связать локальный репозиторий с удаленным
>> origin  стандартный псевдоним, с помощью которого можно обращаться к главному удалённому репозиторию (обычно такой репозиторий один). Это значительно упрощает работу.

+ git push *origin* *branch* - залить изменения в удаленный репозиторий на выбранную ветку
+ git pull *origin* *branch* - синхронизировать содержимое локального репозитория с удаленным с выбранной ветки
+ git branch -C *branch* - создание новой ветки
+ git clone - клонировать репозиторий на локальную машину

