# 开启多进程
import multiprocessing

def default(arg):
    print(arg)

if __name__ == "__main__":
    t1 = multiprocessing.Process(target = default, args = (1,))
    t2 = multiprocessing.Process(target = default, args = (2,))
    # t1.daemon = True设置守护进程，当主线程死亡后t1也死亡，守护进程要在start前设置
    t1.start()
    t2.start()
    # 等待子线程结束才往下执行
    t1.join()
    t2.join()

# 开启多线程
import threading

def defaults(arg):
    print(arg)

if __name__ == "__main__":
    t1 = threading.Thread(target = defaults, args=(1,))
    t2 = threading.Thread(target = defaults, args=(2,))
    # t1.setDaemon = True设置守护线程
    t1.start()
    t2.start()
    t1.join()
    t2.join()