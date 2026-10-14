# tournament-bracket-generetor

**A small website built in go**

Its a website that will generate a bracket for a number of players.
And then start a simulation to see in real time who won.

## The technologies used

- go (*and only the ***standard*** library*)
- HTML
- html/template

## Directory structure
```text
.
├── README.md
├── go.mod
├── main.go
├── handlers
│   └── handlers.go
├── generator
    ├── bracket.go
    ├── list.go
    └── simulation.go
├── generator
    └── datas.go
├── assets/css
    └── style.css
└── templates
    ├── index.html
    └── simulation.html
```

## Depedencies
You just need go version of 1.26.1 or higher

## How to run the website ?

Just type the command `go run .` in the terminal.
And then it display the link to!

## State
It's actually finished !
