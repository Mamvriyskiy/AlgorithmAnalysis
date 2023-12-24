#ifndef _MATRIX_H_
#define _MATRIX_H_

double** allocated_matrix(int n);
void read_matrix(double** matrix, int n);
void copyMatrix(int rows, int cols, double **matrixA,double **matrixC);
void print_matrix(double** matrix, int n);

#endif
