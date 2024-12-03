import numpy as np
import matplotlib.pyplot as plt

def main():
    ## 电场计算
    a = 6
    b = 12
    rho = 100
    max = 50
    densePxiel = 100

    # 电场计算 作图

    X = np.linspace(0, max, max * densePxiel, endpoint=False)
    #[0,a] 为 0
    Y = np.zeros(max * densePxiel)
    #[a,b] 为 rho * (x-a) / x^2
    Y[a * densePxiel:b * densePxiel] = rho * (X[a * densePxiel:b * densePxiel] - a) / X[a * densePxiel:b * densePxiel] ** 2
    #[b,max] 为 rho * (b-a) / x^2
    Y[b * densePxiel:] = rho * (b - a) / X[b * densePxiel:] ** 2

    plt.plot(X, Y, color='green', linewidth=1, linestyle="-")

    plt.show()






if __name__ == '__main__':
    main()
