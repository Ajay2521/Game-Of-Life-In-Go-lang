## Game-Of-Life-In-Go-lang

The universe of the Game of Life is an infinite, two-dimensional orthogonal grid of square cells, each of which is in one of two possible states, live or dead, (or populated and unpopulated, respectively). Every cell interacts with its eight neighbours, which are the cells that are horizontally, vertically, or diagonally adjacent. At each step in time, the following transitions occur:

- Any live cell with fewer than two live neighbours dies, as if by underpopulation.
- Any live cell with two or three live neighbours lives on to the next generation.
- Any live cell with more than three live neighbours dies, as if by overpopulation.
- Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.

These rules, which compare the behavior of the automaton to real life, can be condensed into the following:

- Any live cell with two or three live neighbours survives.
- Any dead cell with three live neighbours becomes a live cell.
- All other live cells die in the next generation. Similarly, all other dead cells stay dead.

### To run the program

#### Run the GameOfLife.go life in your system by using the command

`go run GameOflife.go`

The Program ask for necessary inputs like:

- Size of the board to be created.
- Number of Iteration(Gneration) needed to be created

Program provide 2 difference choice of generating the board such as

- User Generated Board -> Where user give the row and column index and program set the values as 1's at that user given index on the board.
- Auto Generated Board -> Where the program generate the random row and column index and set the values as 1's at generated index in the board

### Sample Output

#### User Generated Board

<p align="center">

<img src="https://user-images.githubusercontent.com/60919132/148764553-eb3391a4-8935-449c-a585-158b625d6ba5.png" >

</p>

<p align="center">

<img src="https://user-images.githubusercontent.com/60919132/148764743-fb82ad33-5d8a-41d9-a53b-2b10c2f89d89.png" >

</p>

<p align="center">

<img src="https://user-images.githubusercontent.com/60919132/148764910-6061da05-698d-4ef4-aa62-2e5024891649.png" >

</p>


#### Auto Generated Board

<p align="center">

<img src="https://user-images.githubusercontent.com/60919132/148765498-efed6546-4cf1-4127-9d9a-2693618fe14a.png" >

</p>

<p align="center">

<img src="https://user-images.githubusercontent.com/60919132/148765568-f00863d4-09bb-4689-bcd3-aea365f03c20.png" >

</p>


#### For input of 50x50 Matrix and Iteration as 100

<p align="center">

<img src="https://user-images.githubusercontent.com/60919132/148765758-a863df93-b7ce-420c-b35a-ecb1cc22cf96.png" >

</p>

