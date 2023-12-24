import pandas as pd
import matplotlib.pyplot as plt
import numpy as np
import sys

args = sys.argv[1:]

lenl = int(args[-1])
lenght = []
k = 0
for i in range(lenl):
    lenght.append(int(args[k]))
    k += 1

a = []
for i in range(lenl):
    a.append(int(args[k]))
    k += 1

c = []
for i in range(lenl):
    c.append(int(args[k]))
    k += 1

d = []
for i in range(lenl):
    d.append(int(args[k]))
    k += 1


# Данные
data = {
    'Размерность': lenght,
    'Сандартный' : a,
    'Виноград' : c,
    'ВиноградОпт' : d
}


markers = ['o', 's', '^', 'D', 'x']

# Создание DataFrame из данных
df = pd.DataFrame(data)

# Отрисовка графиков
plt.figure(figsize=(12, 6))

k = 0
# Добавление графиков для каждой колонки
for i, column in enumerate(df.columns[1:]):
    plt.plot(df['Размерность'], df[column], label=column, marker = markers[i])

plt.yscale('log')
# plt.semilogx()

# Настройка легенды и заголовка
plt.legend(loc='upper left')
plt.title('График зависимости времени выполнения от размерности матрицы')
plt.xlabel('Размерность')
plt.ylabel('Время выполнения (мс)')

# Отображение графика
plt.show()
