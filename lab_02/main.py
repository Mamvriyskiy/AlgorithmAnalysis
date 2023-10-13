import numpy as np
import matplotlib.pyplot as plt
from mpl_toolkits.mplot3d import Axes3D

# Создаем данные
x = np.linspace(0, 3, 100)  # Определите диапазон значений x
y = np.linspace(0, 3, 100)  # Определите диапазон значений y
x, y = np.meshgrid(x, y)
z = y**2

# Создаем график
fig = plt.figure()
ax = fig.add_subplot(111, projection='3d')

# Построение поверхности z = y^2
ax.plot_surface(x, y, z, cmap='viridis', alpha=0.8, label='z = y^2')

# Построение линии y = 3 - x
x_line = np.linspace(0, 3, 100)
y_line = 3 - x_line
ax.plot(x_line, y_line, zs=0, zdir='y', label='y = 3 - x', color='blue')

# Построение плоскости z = 4
z_plane = np.full_like(x, 4)
ax.plot_surface(x, y, z_plane, color='red', alpha=0.5, label='z = 4')

# Построение линии x = 0
y_line_x = np.linspace(0, 3, 100)
x_line_x = np.zeros_like(y_line_x)
ax.plot(x_line_x, y_line_x, zs=0, zdir='x', label='x = 0', color='green')

# Добавление меток
ax.set_xlabel('X')
ax.set_ylabel('Y')
ax.set_zlabel('Z')

# Создаем легенду с указанием цветов
# legend = ax.legend(loc='best')
# for handle in legend.legendHandles:
#     handle.set_color(handle.get_linestyle())

# Отображение графика
plt.show()
