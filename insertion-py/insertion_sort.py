def insertionSort(arr: list[int]) -> list[int]:
    for j in range(2, len(arr)):
        key = arr[j]

        i = j - 1
        
        while i >= 0 and arr[i] > key:
            arr[i+1] = arr[i]
            i -= 1
        
        arr[i+1] = key

    return arr

import random as rd
from datetime import datetime
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
