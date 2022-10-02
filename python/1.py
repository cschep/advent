import collections
from itertools import islice


def sliding_window(iterable, n):
    # sliding_window('ABCDEFG', 4) -> ABCD BCDE CDEF DEFG
    it = iter(iterable)
    window = collections.deque(islice(it, n), maxlen=n)
    if len(window) == n:
        yield tuple(window)
    for x in it:
        window.append(x)
        yield tuple(window)


with open("1.input") as f:
    lines = f.readlines()
    depths = [int(x) for x in lines]
    prev = depths[0]
    total = 0
    for depth in depths:
        if depth > prev:
            total += 1
        prev = depth
    # part 1
    print(total)

    prev = None
    total = 0
    curr = 0
    for triple in sliding_window(depths, 3):
        curr = sum(triple)
        if prev is None:
            prev = curr
            continue
        if curr > prev:
            total += 1
        prev = curr
    # part 2
    print(total)
