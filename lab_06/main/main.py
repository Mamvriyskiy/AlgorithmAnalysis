from fullComb import full_combinations
from matrix import *
from ant import ant_algorithm
from testAlg import time_test

MSG = "\n         Меню \n \
1. Полный перебор \n \
2. Муравьиный алгоритм \n \
3. Параметризация \n \
4. Замерить время \n \
0. Выход \n \
\
Выбор: \
    "

def parametrization(type = 1):
    alpha_arr = [num / 10 for num in range(1, 10)]
    k_eva_arr = [num / 10 for num in range(1, 9)]
    days_arr = [1, 3, 5, 10, 50, 100, 300, 500]

    size = 9

    matrix1 = read_file_matrix("data/mat9_dif.csv")

    optimal1 = full_combinations(matrix1, size, 100000000)

    file1 = open("data/parametrization_class1.txt", "w")

    count = 0
    count_all = len(alpha_arr) * len(k_eva_arr)

    print()

    for alpha in alpha_arr:
        beta = 1 - alpha

        for k_eva in k_eva_arr:
            count += 1

            for days in days_arr:
                res1 = ant_algorithm(matrix1, size, alpha, beta, k_eva, days, 100000000)

                if (type == 0):
                    sep = " & "
                    ender = " \\\\"
                elif (type == 1):
                    sep = ", "
                    ender = ""
                else:
                    sep = " | " 
                    ender = ""

                if res1[0] - optimal1[0] == 0:
                    str1 = "%4.1f%s%4.1f%s%4d%s%5d%s%5d%s\n" \
                        % (alpha, sep, k_eva, sep, days, sep, optimal1[0], sep, res1[0] - optimal1[0], ender)
                    file1.write(str1)
            

            print("Progress: %3d%s" %((count / count_all) * 100, "%"))
            
    file1.close()

def parse_full_combinations():
    
    matrix = read_matrix()
    if not(matrix.any()):
        return 
    
    size = len(matrix)

    max_fuel = int(input("Введите максимально количетво топлива: "))
    result = full_combinations(matrix, size, max_fuel)

    print("\n\nМинимально затраченное топливо = ", result[0], 
            "\nПуть: ", result[1])

def read_koefs():
    alpha = float(input("\n\nВведите коэффициент alpha: " ))
    beta = 1 - alpha
    k_evaporation = float(input("Введите коэффициент evaporation: " ))
    days = int(input("Введите кол-во дней: " ))
    
    return alpha, beta, k_evaporation, days

def parse_ant_alg():

    matrix = read_matrix()
    size = len(matrix)

    alpha, beta, k_evaporation, days = read_koefs()

    max_fuel = int(input("Введите максимально количетво топлива: "))
    result = ant_algorithm(matrix, size, alpha, beta, k_evaporation, days, max_fuel)

    print("\n\nМинимально затраченное топливо = ", result[0], 
            "\nПуть: ", result[1])  

def main():
    option = -1

    while (option != 0):
        option = int(input(MSG))

        if (option == 1):
            parse_full_combinations()
        elif (option == 2):
            parse_ant_alg()
        elif (option == 3):
            parametrization(type = 1)
        elif (option == 4):
            time_test()

        else:
            print("\nПовторите ввод\n")


if __name__ == "__main__":
    main()
