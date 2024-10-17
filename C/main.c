#include <stdlib.h>
#include <time.h>
#include <stdio.h>
#include <string.h>
#include "insertion/insertion.h" 

#define N 10000

int main() {
    int arr[N];
    for(int i = 0; i < N; i++) {
        while(1) {
            int val = rand()%N;
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
        int carr[N];
        memcpy(carr, arr, sizeof(arr));
        clock_t start = clock();
        insertionSort(carr, N);
        clock_t end = clock();
        printf("Execution Time: %f ms\n", ((double)(end - start) / CLOCKS_PER_SEC) * 1000);
    }

}