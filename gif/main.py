import matplotlib.pyplot as plt
import matplotlib.animation as animation

def main():
    
    # 创建图形和坐标轴
    fig, ax = plt.subplots()
    ax.set_xlim(-10, 10)
    ax.set_ylim(-10, 10)

    vX = 0
    vY = 2
    positionX = 5
    positionY = 0

    interval = 10 # 更新时间间隔 ms
    totalSeconds = 10

    # 创建初始圆
    circleSun = plt.Circle((0, 0), 1, color="red", label="sun")
    circleEarth = plt.Circle((positionX,positionY), 0.2, color="blue")
    ax.add_patch(circleEarth)
    ax.add_patch(circleSun)

    frameSeq = 0

    # 定义更新圆位置的函数
    def update(frame):
        nonlocal positionX, positionY
        nonlocal vX, vY
        nonlocal frameSeq

        # 设置地球的加速度
        # 加速度  a = v^2 / r = 4 / 5 = 0.8
        squareDistance = positionX * positionX + positionY * positionY
        aX = -positionX / pow(squareDistance, 1.5) * 20
        aY = -positionY / pow(squareDistance, 1.5) * 20

        # 设置阴影的位置
        if (frameSeq % (1000/interval)) == 0:
            circleShadow = plt.Circle((positionX, positionY), 0.1, color="gray")
            ax.add_patch(circleShadow)
            print("shadow", positionX, positionY)
        
        frameSeq += 1

        
        # 设置地球的速度
        vXNew = vX + aX * interval / 1000
        vYNew = vY + aY * interval / 1000

        # 设置圆的中心随帧移动
        positionX = positionX + (vX + vXNew) / 2 * interval / 1000
        positionY = positionY + (vY + vYNew) / 2 * interval / 1000

        vX = vXNew
        vY = vYNew


        circleEarth.center = positionX, positionY
        return circleEarth,

    # 创建动画，interval 控制速度
    frames = int(totalSeconds *  1000 / interval)
    ani = animation.FuncAnimation(fig, update, frames=frames, interval=interval, blit=False)

    # 显示动画
    plt.show()

if __name__ == '__main__':
    main()