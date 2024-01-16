def standart_alg(text, pattern):
    text_len = len(text)
    pattern_len = len(pattern)

    if pattern_len == 0:
            return "Подстрока пустая"

    if pattern_len > text_len:
        return "Подстрока больше, чем строка"

    i = 0
    result = -1
    while i <= text_len - pattern_len:
        j = 0

        while j < pattern_len and text[i + j] == pattern[j]:
            j += 1

        if j == pattern_len:
            result = i
            break

        i += 1

    return result

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

    if pattern_len == 0:
        return "Подстрока пустая"

    if pattern_len > text_len:
        return "Подстрока больше, чем строка"

    bad_char_table = generate_bad_char_table(pattern)
    good_suffix_table = generate_good_suffix_table(pattern)

    i = pattern_len - 1
    result = -1
    while i < text_len:
        j = pattern_len - 1

        while j >= 0 and text[i] == pattern[j]:
            i -= 1
            j -= 1

        if j == -1:
            result = i + 1
            break

        bad_char_shift = bad_char_table.get(text[i], pattern_len)
        good_suffix_shift = good_suffix_table[pattern_len - 1 - j]

        i += max(bad_char_shift, good_suffix_shift)

    return result

text = input("Введите строку: ")
pattern = input("Введите подстроку: ")

result = standart_alg(text, pattern)
if result != -1:
    print(f"Подстрока найдена на позиции {result}")
else:
    print("Подстрока не найдена")

result = boyer_moore_search(text, pattern)
if result != -1:
    print(f"Подстрока найдена на позиции {result}")
else:
    print("Подстрока не найдена")

