#include <pthread.h>
#include <time.h>
#include <stdio.h>
#include <stdlib.h> 
#include <unistd.h>
#include "gauss.h"
#include "gaussThread.h"
#include "matrix.h"
#include "/usr/local/opt/libomp/include/omp.h"

static long long getThreadCpuTimeNs() {
    struct timespec t;
    if (clock_gettime(CLOCK_MONOTONIC, &t)) {
        perror("clock_gettime");
        return 0;
    }
    return t.tv_sec * 1000LL + t.tv_nsec / 1000000LL;
}

double** allocated_matrix_init(int n) {
    double** matrix = (double**)malloc(n * sizeof(double*));
    int k = 1;
    #pragma omp parallel for
    for (int i = 0; i < n; i++) {
        matrix[i] = (double*)malloc(n * sizeof(double));

        #pragma omp parallel for
        for (int j = 0; j < n; j++) {
            matrix[i][j] = k++;
        }
    }

    matrix[n - 1][n - 1] = 0;
    return matrix;
}

void printTableHeader() {
    printf("+--------+--------+--------+--------+--------+--------+--------+\n");
    printf("|Размер. | 1      | 2      | 4      | 8      | 16     | 32     |\n");
    printf("+--------+--------+--------+--------+--------+--------+--------+\n");
}

// void printTableRow(int size, int threads, long long time1, long long time2, long long time3) {
//     printf("| %5d | %6d | %6lld | %6lld | %6lld |\n", size, threads, time1, time2, time3);
// }

void printTableFooter() {
    printf("+-------+--------+--------+--------+--------+\n");
}

void check_time() {
    printTableHeader();
    for (int i = 10000; i <= 10000; i += 1000)
    {   
        printf("| %d  ", i);
        for (int j = 1; j <= 32; j *= 2) {
            long long timeRes = 0;
            for (int k = 0; k < 10; k++) {
                double **matrix2 = allocated_matrix_init(i);
                double **inverse2 = allocated_matrix(i);
                for (int k = 0; k < i; k++) {
                    inverse2[k][k] = 1;
                }
                long long start = getThreadCpuTimeNs();
                // clock_t start, end;
                // start = clock();
                gauss_thread(matrix2, inverse2, i, j);
                long long finish = getThreadCpuTimeNs();
                timeRes += finish - start;
                //printf("| %4lf ", cpu_time_used);
                for (int l = 0; l < i; l++) {
                    free(matrix2[i]);
                    free(inverse2[i]);
                }
                free(matrix2);
                free(inverse2);
            }
            printf("| %4lld ", timeRes / 10);
        }
        printf("|\n");
    }
}
