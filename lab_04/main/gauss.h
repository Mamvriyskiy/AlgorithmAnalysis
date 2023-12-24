#ifndef _GAUSS_H_
#define _GAUSS_H_

void multiplyRow(double **matrix, int row, double factor, int n);
void subtractRows(double **matrix, int destRow, int sourceRow, double factor, int n);
void gauss(double **matrix, double **inverse, int n);
void swapRows(double **matrix, int row1, int row2, int n);

#endif
