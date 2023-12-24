import pandas as pd
import matplotlib.pyplot as plt
import numpy as np

data = {
    "Размерноcть": [1, 10, 50, 100, 200, 300, 400, 500, 600, 700, 800, 900, 1000],
    "конвейерная обработка": [9, 20, 101, 200, 454, 546, 722, 955, 1235, 1299, 1386, 1549, 1656],
    "полседовательная обработка": [5, 35, 210, 406, 909, 1239, 1671, 1848, 2219, 2879, 3126, 3261, 3659],
}

markers = ['o', 's', '^', 'D', 'x']

# Создание DataFrame из данных
df = pd.DataFrame(data)

# Отрисовка графиков
plt.figure(figsize=(12, 6))

k = 0
# Добавление графиков для каждой колонки
for i, column in enumerate(df.columns[1:]):
    plt.plot(df['Размерноcть'], df[column], label=column, marker = markers[k])

# Настройка легенды и заголовка
plt.legend(loc='upper left')
plt.title('График зависимости времени выполнения от размерности матрицы')
plt.xlabel('Размерность')
plt.ylabel('Время выполнения (мc)')

# Отображение графика
plt.show()

