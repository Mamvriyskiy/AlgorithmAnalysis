#include <pthread.h>
#include <stdlib.h>
#include <stdio.h>
#include <math.h>
#include "matrix.h"

void multiplyRow(double **matrix, int row, double factor, int n) {
    for (int j = 0; j < n; j++) {
        matrix[row][j] *= factor;
    }
}

void subtractRows(double **matrix, int destRow, int sourceRow, double factor, int n) {
    for (int j = 0; j < n; j++) {
        matrix[destRow][j] -= factor * matrix[sourceRow][j];
    }
}

void swapRows(double **matrix, int row1, int row2, int n) {
    double *temp = matrix[row1];
    matrix[row1] = matrix[row2];
    matrix[row2] = temp;
}

void gauss(double **matrix, double **inverse, int n) {
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

        for (int j = 0; j < n; j++) {
            if (j != i) {
                double factor = matrix[j][i];
                subtractRows(matrix, j, i, factor / pivot, n);
                subtractRows(inverse, j, i, factor / pivot, n);
            }
        }
    }
}

