
# Hometask 08

improve and refactor game:
* add file CHANGE.md for see all list of changes
* refactor scripts/test.mk
* improve - add package config:
  * refactor flags
  * refactor creating []player.Player and field.Field
  * make http variant of game default

# Hometask 07

improve and refactor game:
* game improves:
  * add posibility to start new game in html version of game,
  * add configuring before start new html variant game,
  * add winSeq variable to struct Field, package field for limit win sequence length on big fields,
* game refactors:
  * refactor go modules files
  * make new flags system (not finished, should be changed):
    * a - ask all settings (except html varian)
    * s - field size
    * w - win sequence length
    * p - first player side (X or O)
    * v - game version (cli or html)
  * move interfaces to separate package common
  * move some errors to package common

# Hometask 06

make basic http game with winner check:
* refact metod AssignCell() of Field package field, it do not need to use pointer to assign new value, because it does not change capacity of slice
* rename package httpServer to game, than:
* remove struct tictactoeHandler
* add struct Game with fields: iter, playerIdx, winner, draw, players, field
* add errors ErrWinnerExist, ErrNewWinner, ErrNoWinner
* add method NextMove() of Game which content basic game logic
* add methid GetWinnerString of Game it return formated winner (or draw) string or empty string and error ErrNoWinner
* improve method ToHtml() of Field package field:
  * add redirect to location /{col}/{row} on click to particular cell
  * move template to separate string variable htmlTemplate
  * add extra field after printing "field" based on input variadic arguments
* refactor go.mod
* add package httpGame:
  * add interface IGame, based on struct Game from package game
  * add struct HttpGame, content one field game with type IGame
  * add starting http server in func NewHttpGame()
  * add method ServeHTTP() of HttpGame for extract clicked cell, send it to IGame and return html page from it
* use package game and httpGame instead of httpServer
* refactor cmd/tic-tac-toe/main.go - use package game for cli
* sync logic of cli and http game variants


# Hometask 05

## Original hometask:

* create file cmd/orig\_hometask\_05/main.go
* create struct Fibonacci with private field cache
* add constructor NewFibonacci(),
  which init struct and add to cache 2 first numbers (need for calculate any next number)
* add method Get() of Fibonacci,
  which calculate number by it sequence position via formula
  Fib(i) = Fib(i-1) + Fib(i-2) and put all calculated numbers into cache.
  If number already in cache return value from cache
* add cli request for enter needed number in sequence
* add measure of time (in us) of calculating particular number

## Advanced hometask (add initial support of http server):

* refact package player: rename instant this to p
* refact package field: move errors to package shared variables
* refact Field, make some methods private
* update unit tests for package field
* add method ToHtml() of field.Field for generate html page of field state
* add package httpServer, it content:
  * interfaces IPlayer, IField
  * struct tictactoeHandler with method ServeHTTP() and constructor NewHttpGame()
* add func startBrowser(), copied from tour of golang
* implement starting http server:
  * add flag -h for use web interface instead of cli
  * add running http server on port 8080
  * add starting browser for load page with game

# Hometask 04

* cmd/orig\_hometask\_04/main.go:
  * implement sorting algorithm
* Add new features and refactor package field:
  * add field cells to struct Field
  * add initialization of field cells in constructor New()
  * add methods IsCellFree(), AssignCell(), GetCellValue(), IsFieldFull(), IsColumnWinner(), IsRowWinner(), IsDiagStreightWinner(), IsDiagReverseWinner(), IsCellWinner() to struct Field
  * update unit tests
* Refactor and update cmd/tic-tac-toe/main.go:
  * rename instance of Field from "field" to "instField"
  * add second player ('X' or 'O', depends on first player)
  * add using all initialized player in every steps
  * improve entering coordinates of cell if error happened
  * add assigning cells
  * add check winners
  * add checking draw
  * add printing Field after every input

# Hometask 03

* remove overingeneering, keep only code need acording [GB.pdf](https://github.com/pavelmvl/go-tic-tac-toe/files/8041203/GB.pdf),
  except 2 checks marked as done after lesson 4

# Hometask 02

* add module 'internal/field'
* reformat output in func (this Player)String()
* implement 'internal/field' in 'cmd/tic-tac-toe/main.go'
* update go.mod

# Hometask 01

* add main.go:
  * add struct Player for store info about future players
  * add func Init() for parse input flags or ask user for fill
  * empty 'Mark' variable

* refact app:
  * rename Init() to InitPlayer()
  * move InitPlayer() to package internal/pkg/initPlayer
  * include package initPlayer via go.mod
