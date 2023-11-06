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

## Как гид хранит изменения

Git хеширует информацию о коммите с помощью алгоритма SHA-1 и получает для каждого коммита свой уникальный индификатор — хеш.
Хеш — основной идентификатор коммита и позволяет узнать его автора, дату и содержимое закоммиченных файлов.
Все хеши, а также таблицу соответствий `хеш → информация` о коммите Git хранит в папке .git

В числе прочих файлов в папке .git есть служебный файл HEAD. Он указывает на самый свежий коммит. Вместо хеша последнего коммита можно написать слово HEAD — Git вас поймёт.

## Работа с логами

Можно вызвать не только полный лог, но и сокращённый — `git log --oneline`.
В сокращённом логе выводятся сокращённые хеши — их можно использовать точно так же, как и полные.

## Расшифровывание git status

Есть несколько состояний в которых может находится файл:

+ untracked - файл не отслеживаеться
+ tracked - файла отслеживается
+ staged - в это состояние файла поподает после `git add`
+ modified - файл был изменен

### Жизненый цикл состояний файла

```mermaid
graph LR;
  untracked -- "git add" --> staged;
  staged    -- "git commit"     --> tracked/comitted;
  staged    -- "changing"     --> modified;
  tracked/comitted  --  "changing"  --> modified;
  modified -- "git add" --> staged;
%%graph TB;

```
