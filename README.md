## Magazyn sklepu rowerowego

### Architektura projektu

- **controllers** - widoki projektu
- **database** - połączenie z bazą danych, funkcje do komunikacji z nią
- **models** - modele bazy danych
- **routes** - ścieżki URL aplikacji
- **helpers** - funkcje pomocnicze
- **middleware** - funkcje do uwierzytalniania uzytkownika


### Uruchamianie projektu

Aby uruchomić projekt należy mieć na komputerze aplikacje Docker.
Jeśli już ją mamy należy w terminalu w głównym katalogu naszego projektu wpisać `docker compose up --build`
Po zbudowaniu sie aplikacji będzie ona dostępna pod adresem http://localhost:8000

### Baza danych

Po uruchomieniu projektu mamy możliwość dostępu do bazy danych w taki sposób:
- `docker exec -it <nazwa-kontenera> bash`, u mnie `docker exec -it warehouse-backend-frontend-db-1 bash`
- `psql -U group15 -d group15`
- teraz możemy wykonywać komendy SQL na naszych tabelach, żeby wylistować wszystkie tabele należy napisać `\dt`


### Importy

Żeby zaimportować funkcje z jednego modułu do drugiego wystarczy w importach wpisać `"warehouse/<nazwa-modułu>"`

### Metody do modeli

Jeśli jest metoda, która nawiązuje do modelu np. `CreateUser` lub `GetUserPassword` dajemy ją w pliku tego modelu

### Zresetowanie bazy danych aby odtworzyć ją od nowa

Jeśli zmieniłeś coś w bazie danych (migrations.go) to musisz wykonać tą komendę: docker_utils/clear.sh
