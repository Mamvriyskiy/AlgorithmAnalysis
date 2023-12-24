#include <stdio.h>
#include <stdlib.h>
#include <pthread.h>
#include "gaussThread.h"
#include "gauss.h"
#include "matrix.h"

void subtractRowsEx(double **matrix, double **inverse, int start, int end, int n, int i) {
    for (int j = start; j < end; j++) {
        if (j != i) {
            double factor = matrix[j][i];
            subtractRows(matrix, j, i, factor / matrix[i][i], n);
            subtractRows(inverse, j, i, factor / matrix[i][i], n);
        }
    }
}


void *subtractRowsThread(void *args) {
    GaussArgs *gaussArgs = (GaussArgs *)args;
    subtractRowsEx(gaussArgs->matrix, gaussArgs->inverse, gaussArgs->start, gaussArgs->end, gaussArgs->n, gaussArgs->num);
    return NULL;
}

void gauss_thread(double **matrix, double **inverse, int n, int numThreads) {
    int err = 0;
    for (int i = 0; i < n; i++) {
        double pivot = matrix[i][i];
        if (pivot == 0.0) {
            int k;
            for (k = i + 1; k < n && matrix[k][i] == 0.0; k++);
            if (k >= n) {
                break;
            } else {
                swapRows(matrix, i, k, n);
                swapRows(inverse, i, k, n);
                pivot = matrix[i][i];
            }
        }

        subtractRowsThreadEx(matrix, inverse, n, numThreads, i);
    }
}

void subtractRowsThreadEx(double **matrix, double **inverse, int n, int numThreads, int i) {
    pthread_t threads[MAX_N];
    GaussArgs args[MAX_N];
    for (int t = 0; t < numThreads; t++) {
        args[t].matrix = matrix;
        args[t].inverse = inverse;
        args[t].n = n;
        args[t].start = t * (n / numThreads);
        args[t].end = (t + 1) * (n / numThreads);
        args[t].num = i;

        if (t == numThreads - 1) {
            args[t].end = n;
        }

        pthread_create(&threads[t], NULL, subtractRowsThread, (void *)&args[t]);
    }

    for (int t = 0; t < numThreads; ++t) {
        pthread_join(threads[t], NULL);
    }
}
