#include <stdio.h>
#include <stdlib.h>
#include "matrix.h"

double** allocated_matrix(int n) {
	double** matrix = (double**)malloc(n * sizeof(double*));
	
	#pragma omp parallel for
	for (int i = 0; i < n; i++) {
		matrix[i] = (double*)malloc(n * sizeof(double));
	}

	return matrix;
}

void read_matrix(double** matrix, int n) {
	for (int i = 0; i < n; i++) {
		for (int j = 0; j < n; j++) {
			scanf("%lf", &matrix[i][j]);
		}
	}
}

void copyMatrix(int rows, int cols, double **matrixA,double **matrixC) {
	#pragma omp parallel for
    for (int i = 0; i < rows; ++i) {
		#pragma omp parallel for
        for (int j = 0; j < cols; ++j) {
            matrixC[i][j] = matrixA[i][j];
        }
    }
}

void print_matrix(double** matrix, int n) {
	for (int i = 0; i < n; i++) {
		for (int j = 0; j < n; j++) {
			printf("%lf ", matrix[i][j]);
		}
		printf("\n");
	}
}
