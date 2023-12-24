import numpy as np
from random import random 

MIN_PHEROMONE = 0.01

def calc_q(matrix, size):  
    q = 0
    count = 0

    for i in range(size):
        for j in range(size):
            if (i != j):
                q += matrix[i][j]
                count += 1

    return q / count

def get_pheromones(size):
    return [[1 for _ in range(size)] for _ in range(size)]

def get_visibility(matrix, size):
    return [[(1.0 / matrix[i][j] if (i != j) else 0) for j in range(size)] for i in range(size)]

def get_visited_places(route, ants):
    visited = [list() for _ in range(ants)]

    for ant in range(ants):
        visited[ant].append(route[ant])

    return visited

def calc_length(matrix, route):
    length = 0

    for way_len in range(1, len(route)):
        length += matrix[route[way_len - 1], route[way_len]]

    return length

def update_pheromones(matrix, places, visited, pheromones, q, k_evaporation):
    ants = places

    for i in range(places):
        for j in range(places):
            delta_pheromones = 0

            for ant in range(ants):
                length = calc_length(matrix, visited[ant])
                delta_pheromones += q / length

            pheromones[i][j] *= (1 - k_evaporation)
            pheromones[i][j] += delta_pheromones

            if (pheromones[i][j] < MIN_PHEROMONE):
                pheromones[i][j] = MIN_PHEROMONE

    return pheromones

def find_posibilities(pheromones, visibility, visited, places, ant, alpha, beta):
    pk = [0] * places

    for place in range(places):
        if place not in visited[ant]:
            ant_place = visited[ant][-1]

            pk[place] = pow(pheromones[ant_place][place], alpha) * \
                        pow(visibility[ant_place][place], beta)
        else:
            pk[place] = 0

    sum_pk = sum(pk)

    for place in range(places):
        pk[place] /= sum_pk  

    return pk

def choose_next_place_by_posibility(pk):
    posibility = random()
    choice, chosen_place = 0, 0

    while ((choice < posibility) and (chosen_place < len(pk))):
        choice += pk[chosen_place]
        chosen_place += 1

    return chosen_place

def res_calc_length(matrix, route, max_fuel):
    length = 0
    for way_len in range(1, len(route)):
        length += matrix[route[way_len - 1], route[way_len]]
        if length > max_fuel:
            length -= matrix[route[way_len - 1], route[way_len]]
            return length, way_len, 1

    return length, len(route), 0

def ant_algorithm(matrix, places, alpha, beta, k_evaporation, days, max_fuel):
    q = calc_q(matrix, places)

    best_way = []
    min_dist = float("inf")
    max_visit_cites = 0

    pheromones = get_pheromones(places)
    visibility = get_visibility(matrix, places)

    ants = places

    for _ in range(days): 
        route = np.arange(places)
        visited = get_visited_places(route, ants)

        for ant in range(ants):
            while len(visited[ant]) != places:
                
                pk = find_posibilities(pheromones, visibility, visited, places, ant, alpha, beta)
                chosen_place = choose_next_place_by_posibility(pk)
                visited[ant].append(chosen_place - 1)
        
            cur_dist, cur_visit, flag  = res_calc_length(matrix, visited[ant], max_fuel)

            if (cur_dist < min_dist and flag == 0):
                min_dist = cur_dist
                best_way = visited[ant]
                max_visit_cites = cur_visit
            elif (max_visit_cites < cur_visit and cur_dist < min_dist):
                min_dist = cur_dist
                max_visit_cites = cur_visit
                best_way = visited[ant][ : cur_visit]
    
        pheromones = update_pheromones(matrix, places, visited, pheromones, q, k_evaporation)

    return min_dist, best_way
