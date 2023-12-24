import numpy as np

def read_file_matrix(file_name):
    #file_name = "data/m1.txt"
    print(file_name)
    try:
        file = open(file_name, "r")
    except FileNotFoundError:
        print("Файл не найден")
        return 0

    size = len(file.readline().split()) # определение размера матрицы
    file.seek(0)

    matrix = np.zeros((size, size), dtype = int)
    
    i = 0

    for line in file.readlines():
        j = 0

        for num in line.split():
            matrix[i][j] = int(num)
            j += 1

        i += 1

    file.close()

    return matrix

def print_matrix():
    num_file = input("\nВведите имя файла: ")

    matrix = read_file_matrix(num_file)

    print("\n")

    for i in range(len(matrix)):
        for j in range(len(matrix[0])):

            print("%4d" % (matrix[i][j]), end = "")

        print()


def read_matrix():
    num_file = input("\nВыберите файл: ")
    matrix = read_file_matrix(num_file)

    return matrix
