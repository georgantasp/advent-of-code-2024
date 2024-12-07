# Day 7

Part1 ~30 min. Right before starting I read a message about stack overflow and recursion. Did that influence me or is recursion just the right way to do it? Anyway, I think I let myself get clever here, the equations get computed left to right, so my recursive function applies the inverse computation right to left.

Part2 ~45 min. The cleverness above didn't help me here. Trying to make the same approach work with the concat operator got really ugly. I made repeated attempts throwing in more and more conditions and even considered if int overflow was happening. Finally, I switched to passing in the runningTest param and going left to right. It ended up way more simple. Why didn't I do this in part 1?
