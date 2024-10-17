#include <stdio.h>
#include <time.h>
#include <stdlib.h>
#include <string.h>

int* insertionSort(int* arr, int len) {
    for(int j = 2; j < len; j++) {
        int key = arr[j];
        int i = j - 1;

        while(i >= 0 && arr[i] > key) {
            arr[i+1] = arr[i];
            i -= 1;
        }

        arr[i+1] = key;
    }

    return arr;
}

int main() {
    int arr[10000];
    for(int i = 0; i < 10000; i++) {
        while(1) {
            int val = rand()%10000;
            int isDuplicated = 0;

            for(int j = 0; j < i; j++) {
                if(val == arr[j]) {
                    isDuplicated = 1;
                    break;
                }
            }

            if(!isDuplicated) {
                arr[i] = val;
                break;
            }
        }
    }

    for(int i = 0; i < 10; i++) {
        int carr[10000];
        memcpy(carr, arr, sizeof(arr));
        clock_t start = clock();
        insertionSort(carr, 10000);
        clock_t end = clock();
        printf("Execution Time: %f ms\n", ((double)(end - start) / CLOCKS_PER_SEC) * 1000);
    }

}