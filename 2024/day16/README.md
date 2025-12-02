# Day 16

This one was hard.

Part1 I had an implementation pretty quickly that satisfied the example. Unfortunately, it didn't work for the real input. I had to implement a mechanism to hold the state of the best path score to reach each point and short out if any new path came through with a score of greater or equal value.

Part2 I tried multiple ways of storing paths and passing them through the recursion. Nothing worked. Finally I settled on piggybacking on the implementation above. I had to tweak to allow paths of equal value. Then I recorded the best scores as the recursion returned. Finally I iterated the maze and counted where the best score had been recorded.