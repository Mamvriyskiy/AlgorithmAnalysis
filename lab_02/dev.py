def is_safe(board, row, col):
    # Проверяем, что нет ферзя в том же столбце выше по доске
    for i in range(row):
        if board[i][col] == 1:
            return False
    
    # Проверяем, что нет ферзя на диагонали слева вверх
    for i, j in zip(range(row, -1, -1), range(col, -1, -1)):
        if board[i][j] == 1:
            return False
    
    # Проверяем, что нет ферзя на диагонали справа вверх
    for i, j in zip(range(row, -1, -1), range(col, 8)):
        if board[i][j] == 1:
            return False
    
    return True

def solve_queens(board, row):
    if row == 25:
        # Все ферзи размещены, добавляем текущую конфигурацию в решения
        solutions.append([row[:] for row in board])
        return
    
    for col in range(25):
        if is_safe(board, row, col):
            board[row][col] = 1
            solve_queens(board, row + 1)
            board[row][col] = 0

# Создаем пустую доску
chessboard = [[0 for _ in range(25)] for _ in range(25)]
solutions = []

# Начинаем решение с первой строки
solve_queens(chessboard, 0)

print(len(solutions))
# Выводим найденные конфигурации
for solution in solutions:
    for row in solution:
        print(' '.join(['Q' if x == 1 else '.' for x in row]))
    print('\n')



