# Lem-in - Digital Ant Farm

Lem-in is a project that simulates a digital version of an ant farm. The objective of this program is to find the quickest path for a given number of ants to navigate through a colony composed of rooms and tunnels. The program reads input from a file, which describes the layout of the ant farm and specifies the starting and ending points.

## How It Works

- An ant farm is constructed with interconnected rooms and tunnels.
- The ants are initially placed in the room labeled "##start".
- The goal is to guide the ants to the room labeled "##end" using the shortest possible path.
- It's important to note that the shortest path may not always be the simplest.
- The program handles various scenarios, such as colonies with complex room structures, loops, invalid data formats, missing start/end rooms, duplicated rooms, and unknown room connections.
- In such cases, the program provides appropriate error messages to indicate the issue.

## Input Format

The input file follows a specific structure. Here's an example:



```text
##start
1 23 3
2 16 7
3 16 3
4 16 5
5 9 3
6 1 5
7 4 8
##end
0 9 5
0-4
0-6
1-3
4-3
5-2
3-5
4-2
2-1
7-6
7-2
7-4
6-5
```

- The file begins with the declaration of the start room using `##start` followed by its coordinates.
- The file then lists the rooms and their respective coordinates.
- Any lines starting with `#` are comments and should be ignored during parsing.
- The end room is declared using `##end` followed by its coordinates.
- The connections between rooms are defined using the `name1-name2` format.

## Output Format

The program displays the results on the standard output in the following format:

```text
number_of_ants
the_rooms
the_links

Lx-y Lz-w Lr-o ...
```

- `number_of_ants` indicates the total number of ants in the simulation.
- `the_rooms` lists the rooms in the ant farm, with each room defined by `name coord_x coord_y` format.
- `the_links` describes the connections between rooms, using `name1-name2` format.
- The subsequent lines display the movements of the ants, where `x`, `z`, `r` represent the ant numbers, and `y`, `w`, `o` represent the room names.

## Instructions and Constraints

- Create tunnels and rooms to construct the ant farm.
- Room names should not start with `L` or `#` and should not contain spaces.
- Connect the rooms using tunnels, ensuring that each tunnel connects only two rooms.
- Each room, except for `##start` and `##end`, can contain only one ant at a time.
- Each tunnel can be used only once per turn.
- The objective is to guide the ants through the shortest paths while avoiding congestion and stepping over other ants.
- Display only the ants that moved in each turn and move each ant only once through a tunnel.
- Ignore any unknown commands encountered during parsing.
- The program should handle errors gracefully and avoid unexpected crashes.
- Room coordinates will always be integers.
- The project must be implemented in Go.
- The code should follow good programming practices.

## Usage

To run the Lem-in program, follow these steps:

```shell
$ go build -o lem-in
$ ./lem-in input.txt
```

Please ensure that you have Go installed on your system before running the program.

## Dependencies

The Lem-in project does not require any external dependencies.
