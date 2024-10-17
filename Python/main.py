import random as rd
from datetime import datetime
from insertion import insertionSort


if __name__ == "__main__":
    arr = [0] * 10000
    for i in range(10000):
        while True:
            val = rd.randint(0, 9999)
            isDuplicate = False

            for j in range(i):
                if val == arr[j]:
                    isDuplicate = True
                    break

            if not isDuplicate:
                arr[i] = val
                break

    #print(arr)
    for _ in range(10):
        carr = list(arr)
        start = datetime.now()
        insertionSort(carr)
        print("Excution Time: ", (datetime.now() - start).microseconds / 1000, "ms")