# Mars Rovers

## Problem

A squad of robotic rovers is to be landed by NASA on a plateau on Mars. The rovers must navigate this curiously rectangular plateau so that their onboard cameras can get a complete view of the surrounding terrain to send back to Earth. The plateau is divided up into a grid to simplify navigation. A rover's position and location are represented by a combination of x and y coordinates and a letter describing one of the four cardinal compass points. An example position might be 0, 0, N, which means the rover is in the bottom left corner and facing North. To control a rover, NASA sends a simple string of letters.  
The possible letters are `L`, `R`, and `M`. `L` and `R` make the rover spin 90 degrees left or right without moving from its current spot. 'M' means moving one grid point forward and maintaining the same heading.

Assume that the square directly North from (x, y) is (x, y+1).

### Input

Each rover has two lines of input. The first line gives the rover's position, and the second is a series of instructions telling the rover how to explore the plateau. The first line of input is the upper-right coordinates of it; the lower-left coordinates are assumed to be 0,0. The rest of the input is information about the rovers that have been deployed. The position comprises two integers and a letter separated by spaces, corresponding to the x and y coordinates and the rover's orientation. Each rover will be finished sequentially, which means the second rover won't start moving until the first one has finished moving.

### Output

The output for each rover should be its final coordinates and heading.

#### Test Input

```text
5 5
1 2 N
LMLMLMLMM
3 3 E
MMRMMRMRRM
```

#### Expected Output

```text
1 3 N
5 1 E
```
