gray on grey
============

Install:
```
go get -u github.com/therealplato/grayongrey/cmd/gg
```

Usage with pipe:
```
echo 'Athens north=Beirut west=Cairo south=Dunkirk
Beirut south=Athens west=Fargo' | gg
```

Usage with file:
```
gg filename
```

Custom alien count:
```
gg -n 10 filename
```

Problem Statement
-----------------

Input a map describing city topography as name and cardinal direction links:
```
Athens north=Beirut west=Cairo south=Dunkirk
Beirut south=Athens west=Fargo
```

Place (command line flag) N aliens randomly on the cities.
Each iteration consists of alien movement, alien fights and destroyed cities.
Assumption: Probability of alien moving along every edge is equal.
Assumption: Probability of alien not moving is equal to probability of any specific move.

After moves, if two aliens are in the same place, they fight and destroy each other and the city.
Assumption: If more than two aliens are in the same place, they all are destroyed along with the city.

Destroyed cities cannot be traveled to.

When a fight occurs, log the city and participants.

Terminate the simulation when all aliens are destroyed, or each alien has moved at least 10,000 times.
Assumption: Staying still counts as a move.

Print out the final state of the world in the same format as input.
