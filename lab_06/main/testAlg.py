import matplotlib.pyplot as plt
from time import process_time
from fullComb import full_combinations
from ant import ant_algorithm
from random import randint
import numpy as np

def generate_matrix(size, rand_start, rand_end):
    
    matrix = np.zeros((size, size), dtype = int)

    for i in range(size):
        for j in range(size):

            if (i == j):
                num = 0
            else:
                num = randint(rand_start, rand_end)

            matrix[i][j] = num
            matrix[j][i] = num

    return matrix

def time_test():
    
    size_start = int(input("\n\nВведите начальный размер матрицы: "))
    if size_start <= 1:
        print("Данные введены неверно. Введите от 2 до ...")
        return
    size_end = int(input("Введите конечный размер матрицы: "))
    fuel = 1000000000000

    time_full_combs = []
    time_ant = []

    sizes_arr = [i for i in range(size_start, size_end + 1)]

    count = 0
    print()

    # Time
    for size in sizes_arr:
        count += 1

        matrix = generate_matrix(size, 1, 2)

        # Full combinations
        start = process_time()
        full_combinations(matrix, size, fuel)
        end = process_time()

        time_full_combs.append(end - start)

        # Ant algorythm
        start = process_time()
        ant_algorithm(matrix, size, 0.5, 0.5, 0.5, 250, fuel)
        end = process_time()

        time_ant.append(end - start)

        print("Progress: %3d%s" %((count / len(sizes_arr)) * 100, "%"))

    # Table
    print("\n %s | %s | %s" %("Размер", "Время полного перебора", "Время муравьиного алгоритма"))
    print("-" * (63))

    for i in range(len(sizes_arr)):
        print(" %6d | %22.6f | %27.6f" %(sizes_arr[i], time_full_combs[i], time_ant[i]))

    # Graph
    plt.figure(figsize=(10, 7))
    plt.plot(sizes_arr, time_full_combs, label = "Полный перебор")
    plt.plot(sizes_arr, time_ant, label = "Муравьиный алгоритм")

    plt.legend()
    plt.title("Временные характеристики")
    plt.ylabel("Затраченное время (с)")
    plt.xlabel("Размерность матрицы")
    plt.semilogy()
    plt.show()
    