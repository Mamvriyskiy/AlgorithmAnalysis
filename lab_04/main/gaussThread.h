#ifndef _GAUSS_THREAD_H_
#define _GAUSS_THREAD_H_

#define MAX_N 100

typedef struct {
    double **matrix;
    double **inverse;
    int start;
    int end;
    int n;
    int num;
} GaussArgs;

void *subtractRowsThread(void *args);
void subtractRowsEx(double **matrix, double **inverse, int start, int end, int n, int i);
void gauss_thread(double **matrix, double **inverse, int n, int numThreads);
void subtractRowsThreadEx(double **matrix, double **inverse, int n, int numThreads, int i);

#endif
