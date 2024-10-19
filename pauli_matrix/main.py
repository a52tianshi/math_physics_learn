import numpy as np

def main():
    sigma_x = np.array([[0, 1], [1, 0]])
    sigma_y = np.array([[0, -1j], [1j, 0]])
    sigma_z = np.array([[1, 0], [0, -1]])
    print(sigma_x)
    print(sigma_y)
    print(sigma_z)
    print("sigma_x * sigma_x\n",np.dot(sigma_x, sigma_x))
    print("sigma_y * sigma_y\n",np.dot(sigma_y, sigma_y))
    print("sigma_z * sigma_z\n",np.dot(sigma_z, sigma_z))
    print("sigma_x * sigma_y\n",np.dot(sigma_x, sigma_y))
    print("sigma_y * sigma_x\n",np.dot(sigma_y, sigma_x))
    print("sigma_y * sigma_z\n",np.dot(sigma_y, sigma_z))
    print("sigma_z * sigma_y\n",np.dot(sigma_z, sigma_y))
    print("sigma_x * sigma_z\n",np.dot(sigma_x, sigma_z))
    print("sigma_z * sigma_x\n",np.dot(sigma_z, sigma_x))

    print("sigma_x * sigma_y - i * sigma_z\n",  np.dot(sigma_x, sigma_y) - 1j * sigma_z)


if __name__ == '__main__':
    main()
