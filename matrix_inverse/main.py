import numpy as np

def main():
    # Define Matrix (1,1,0) , (1,0,1), (0,1,1)
    M = [[1, 1, 0], [1, 0, 1], [0, 1, 1]]
    # 转置
    M_T = np.transpose(M)
    # Inverse of Matrix
    M_inv = np.linalg.inv(M_T)
    print(M_inv)
    # 矩阵乘法
    print(np.dot(M_T, M_inv))

if __name__ == '__main__':
    main()