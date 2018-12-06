import datetime
import time

start = time.time()
obj = {}

for i in range(1000000):
    tm = time.time()
    obj[str(i) + '_' + str(tm)] = tm;

print(time.time() - start)
