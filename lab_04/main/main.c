#include <stdio.h>
#include <stdlib.h>
#include <pthread.h>
#include <string.h>
#include "gaussThread.h"
#include "gauss.h"
#include "matrix.h"
#include "checktime.h"

#define MAX_N 100

int main() {
    printf("Меню:\n1)Одинчный подсчет;\n2)Замер времени;\n3)Выход.\n");
    int num;
    scanf("%d", &num);
    if (num == 1) {
        int n;
        printf("Введите размерность матрицы N: ");
        scanf("%d", &n);

        const int numThreads = 1;

        double **matrix1 = allocated_matrix(n);
        double **inverse1 = allocated_matrix(n);
        double **inverse2 = allocated_matrix(n);
        double **matrix2 = allocated_matrix(n);

        for (int i = 0; i < n; i++) {
            inverse1[i][i] = 1;
            inverse2[i][i] = 1;
        }
        read_matrix(matrix1, n);
        copyMatrix(n, n, matrix1, matrix2);

        gauss(matrix1, inverse1, n);
        gauss_thread(matrix2, inverse2, n, numThreads);

        printf("Обратная матрица: \n");
        print_matrix(inverse1, n);

        printf("Обратная матрица параллельным алгоритмом: \n");
        print_matrix(inverse2, n);

        // Вывод результата или выполнение других операций

        for (int i = 0; i < n; ++i) {
            free(matrix1[i]);
            free(matrix2[i]);
            free(inverse1[i]);
            free(inverse2[i]);
        }

        free(matrix1);
        free(matrix2);
        free(inverse1);
        free(inverse2);
    } else if (num == 2) {
        check_time();
    }

    return 0;
}
