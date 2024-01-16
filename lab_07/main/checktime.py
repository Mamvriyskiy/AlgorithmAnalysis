import string, random
from time import process_time
import matplotlib.pyplot as plt
import numpy as np

def standart_alg(text, pattern):
    text_len = len(text)
    pattern_len = len(pattern)

    comparisons = 0
    for i in range(text_len - pattern_len + 1):
        j = 0
        comparisons += 1

        while j < pattern_len and text[i + j] == pattern[j]:
            j += 1
            comparisons += 1

        if j == pattern_len:
            return comparisons

    return comparisons

def generate_bad_char_table(pattern):
    bad_char_table = dict()
    pattern_len = len(pattern)

    for i in range(pattern_len - 1):
        bad_char_table[pattern[i]] = pattern_len - 1 - i

    return bad_char_table

def generate_good_suffix_table(pattern):
    pattern_len = len(pattern)
    suffix_table = [0] * pattern_len
    last_prefix_position = pattern_len

    for i in range(pattern_len - 1, -1, -1):
        if is_prefix(pattern, i + 1):
            last_prefix_position = i + 1
        suffix_table[pattern_len - 1 - i] = last_prefix_position - i + pattern_len - 1

    for i in range(pattern_len - 1):
        j = get_suffix_length(pattern, i)
        suffix_table[j] = min(suffix_table[j], pattern_len - 1 - i + j)

    return suffix_table

def is_prefix(pattern, p):
    pattern_len = len(pattern)
    j = 0

    for i in range(p, pattern_len):
        if pattern[i] != pattern[j]:
            return False
        j += 1

    return True

def get_suffix_length(pattern, p):
    length = 0
    pattern_len = len(pattern)
    j = pattern_len - 1

    while p >= 0 and pattern[p] == pattern[j]:
        length += 1
        p -= 1
        j -= 1

    return length

def boyer_moore_search(text, pattern):
    text_len = len(text)
    pattern_len = len(pattern)

    bad_char_table = generate_bad_char_table(pattern)
    good_suffix_table = generate_good_suffix_table(pattern)

    i = pattern_len - 1
    result = -1
    comparisons = 0  

    while i < text_len:
        j = pattern_len - 1

        while j >= 0 and text[i] == pattern[j]:
            i -= 1
            j -= 1
            comparisons += 1

        if j == -1:
            comparisons += 1
            result = i + 1
            break

        bad_char_shift = bad_char_table.get(text[i], pattern_len)
        good_suffix_shift = good_suffix_table[pattern_len - 1 - j]

        i += max(bad_char_shift, good_suffix_shift)

    return result

def create_random_text(length):
    #letters = 'абвгдеёжзийклмнопрстуфхцчшщъыьэюя'
    letters = 'я'
    result_str = ''.join(random.choice(letters) for _ in range(length - 4))
    return result_str

TIMES = 1000


# def get_process_time(func, key, dictionary):
#     time_res = 0

#     for _ in range(TIMES):
#         time_start = process_time()
#         func(key, dictionary)
#         time_end = process_time()

#         time_res += time_end - time_start

#     return time_res / TIMES

g1A = []
g2A = []
g1B = []
g2B = []

g1AC = []
g2AC = []
g1BC = []
g2BC = []

i = 512
while i <= 8192:
    print("")
    text = create_random_text(i)

    #######
    suml = 0
    for _ in range(10):
        start = process_time()
        cm = standart_alg(text + "data", "data")
        end = process_time()
        suml += end - start
    g1A.append(suml / 10 * 1e6)
    g1AC.append(cm)
    print("Size", i, "Standart", "Time", suml / 10 * 1e6, "Compare: ", cm)

    #######
    suml = 0
    for _ in range(10):
        start = process_time()
        cm = boyer_moore_search(text + "data", "data")
        end = process_time()
        suml += end - start
    g2A.append(suml / 10 * 1e6)
    g2AC.append(cm)
    print("Size", i, "Mur", "Time", suml / 10 * 1e6, "Compare: ", cm)

    #######
    suml = 0
    for _ in range(10):
        start = process_time()
        cm = standart_alg(text + "asdf", "data")
        end = process_time()
        suml += end - start
    g1B.append(suml / 10 * 1e6)
    g1BC.append(cm)
    print("Size", i, "Standart", "Time", suml / 10 * 1e6, "Compare: ", cm)
    
    #######
    suml = 0
    for _ in range(10):
        start = process_time()
        boyer_moore_search(text + "asdf", "data")
        end = process_time()
        suml += end - start
    g2B.append(suml / 10 * 1e6)
    g2BC.append(cm)
    print("Size", i, "Mur", "Time", suml / 10 * 1e6, "Compare: ", cm)
    i *= 2

print(g1A)
print(g1B)

# Названия алгоритмов
sizes = [512, 1024, 2048, 4096, 8192]

# Названия алгоритмов
# algorithms = ['Standart', 'Mur']

cat_par = [f"{sizes[i]}" for i in range(len(sizes))]

# Using the first size from the sizes array
fig, ax = plt.subplots(figsize=(sizes[0]/100, sizes[0]/100))

width = 0.35
x = np.arange(len(cat_par))
rects1 = ax.bar(x - width/2, g1A, width, label='Стандартный')
rects2 = ax.bar(x + width/2, g2A, width, label='Бойер-Мур')

ax.set_title('Сравнение алгоритмов поиска подстроки в строке. Искомая подстрока в конце строки. Время работы реализации алгоритмов.')
ax.set_xticks(x)
ax.set_xticklabels(cat_par)
ax.legend()

plt.show()



# Названия алгоритмов
sizes = [512, 1024, 2048, 4096, 8192]

# Названия алгоритмов
# algorithms = ['Standart', 'Mur']

cat_par = [f"{sizes[i]}" for i in range(len(sizes))]

# Using the first size from the sizes array
fig, ax = plt.subplots(figsize=(sizes[0]/100, sizes[0]/100))

width = 0.35
x = np.arange(len(cat_par))
rects1 = ax.bar(x - width/2, g1B, width, label='Стандартный')
rects2 = ax.bar(x + width/2, g2B, width, label='Бойер-Мур')

ax.set_title('Сравнение алгоритмов поиска подстроки в строке. Искомой подстроки нет в строке. Время работы реализации алгоритмов.')
ax.set_xticks(x)
ax.set_xticklabels(cat_par)
ax.legend()

plt.show()



# Названия алгоритмов
sizes = [512, 1024, 2048, 4096, 8192]

# Названия алгоритмов
# algorithms = ['Standart', 'Mur']

cat_par = [f"{sizes[i]}" for i in range(len(sizes))]

# Using the first size from the sizes array
fig, ax = plt.subplots(figsize=(sizes[0]/100, sizes[0]/100))

width = 0.35
x = np.arange(len(cat_par))
rects1 = ax.bar(x - width/2, g1AC, width, label='Стандартный')
rects2 = ax.bar(x + width/2, g2AC, width, label='Бойер-Мур')

ax.set_title('Сравнение алгоритмов поиска подстроки в строке. Искомая подстрока в конце строки. Количество сравнений.')
ax.set_xticks(x)
ax.set_xticklabels(cat_par)
ax.legend()

plt.show()



# Названия алгоритмов
sizes = [512, 1024, 2048, 4096, 8192]

# Названия алгоритмов
# algorithms = ['Standart', 'Mur']

cat_par = [f"{sizes[i]}" for i in range(len(sizes))]

# Using the first size from the sizes array
fig, ax = plt.subplots(figsize=(sizes[0]/100, sizes[0]/100))

width = 0.35
x = np.arange(len(cat_par))
rects1 = ax.bar(x - width/2, g1BC, width, label='Стандартный')
rects2 = ax.bar(x + width/2, g2BC, width, label='Бойер-Мур')

ax.set_title('Сравнение алгоритмов поиска подстроки в строке. Искомой подстроки нет в строке. Количество сравнений.')
ax.set_xticks(x)
ax.set_xticklabels(cat_par)
ax.legend()

plt.show()
