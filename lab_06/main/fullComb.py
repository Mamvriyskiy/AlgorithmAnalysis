import numpy as np
import itertools as it

def full_combinations(matrix, size, max_fuel):
    places = np.arange(size)
    places_combs = list()

    for combination in it.permutations(places):
        comb_arr = list(combination)
        places_combs.append(comb_arr)

    min_dist = float("inf")
    max_visit_cites = 0

    for i in range(len(places_combs)):
        flag = 0
        cur_dist = 0
        cur_visit = 0

        for j in range(size - 1):
            start_city = places_combs[i][j]
            end_city = places_combs[i][j + 1]

            cur_dist += matrix[start_city][end_city]
            if cur_dist > max_fuel:
                flag = 1
                cur_dist -= matrix[start_city][end_city]
                break

            cur_visit += 1

        if (cur_dist < min_dist and flag == 0):
            min_dist = cur_dist
            best_way = places_combs[i]
            max_visit_cites = cur_visit
        elif (max_visit_cites < cur_visit and cur_dist < min_dist):
            min_dist = cur_dist
            max_visit_cites = cur_visit
            best_way = places_combs[i][ : cur_visit]

    return min_dist, best_way
