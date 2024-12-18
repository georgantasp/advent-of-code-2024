# Day 17

Part 1. This took time just to parse the directions. In the end it was pretty simple to implement.

Part 2. Okay, this certainly takes things to a new level. First thought was to just try brute force it (ofc). I started at A: 0 and iterated. I created some optomization to exit early if the output began to diverge from the program. I left it running for the better part of a day. Nothing. Give up.

Back the next day. It was clear I would need to reverse engineer the program to find a solution. I created this explaination:
```
2,4 -> B = A % 8           :: the first 3 bits : 0-7
1,1 -> B = B ^ 1           :: xor the 1 bit : 1,0,3,2,5,4,7,6
7,5 -> C = A / 2pow(reg B) :: A / 2,1,8,4,32,16,128,64
0,3 -> A = A / 8           :: shift A by 3 bits
1,4 -> B = B ^ 4           :: xor the 2 bit : 1,0,3,2,5,4,7,6 -> 5,4,7,6,1,0,3,2
4,5 -> B = B ^ C           :: ???
5,5 -> out  B % 8          :: the first 3 bits
3,0 -> jump to start
```
First thing that jumps out is that 8 (3 bits) is important. I also see that the starting values for B and C don't matter. Next, I noticed that A has to be sufficiently large to be divided by 8 (shift 3 bits right) at last 15 times to produce enough output. So register A needs to start with a value consisting of 16 3 bit (0-7) "pieces".

Great. I started with `1 << (3*15)` and tried brute forcing again (kidding).

Playing around with different inputs, `2 << (3*15)`, `3 << (3*15)`, `4 << (3*15)`, ect, I saw I could change _just_ the last output. I started manually traversing the remaining "pieces":

```
//r = &registers{
//	A: 5<<(3*15) ^ // 0
//		6<<(3*14) ^ //3
//		0<<(3*13) ^ //5
//		0<<(3*12) ^ //5
//		0<<(3*11) ^ //5
//		0<<(3*10) ^ //4
//		0<<(3*9) ^ //4
//		0<<(3*8) ^ //1
//		0<<(3*7) ^ //3
//		0<<(3*6) ^ //0
//		0<<(3*5) ^ //5
//		0<<(3*4) ^ //7
//		0<<(3*3) ^ //1
//		0<<(3*2) ^ //1
//		0<<(3*1) ^ //4
//		0<<(3*0), //2
//}
```
This exercise ended up being useful because I found that at "piece" 10, no value worked. I would need to "back up" and find the next value for "piece" 11, then iterrate "piece" 10 again.

Finally, I implemented a loop with this logic.